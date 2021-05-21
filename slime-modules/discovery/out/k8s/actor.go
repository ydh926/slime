package k8s

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	istio_types "slime.io/slime/slime-framework/apis/networking/v1alpha3"
	"slime.io/slime/slime-framework/util"
	"slime.io/slime/slime-modules/discovery/model"
	"time"
)

type Actor struct {
	mailbox   model.MailBox
	eventChan chan struct{}
	stop      chan struct{}
	client    client.Client
}

func New(client client.Client) *Actor {
	return &Actor{
		mailbox:   model.NewVersionedMailBox(),
		eventChan: make(chan struct{}),
		client:    client,
		stop:      make(chan struct{}),
	}
}

func eventToObject(event *model.Event) runtime.Object {
	spec, err := util.ProtoToMap(event.Message)
	if err == nil {
		switch event.TypeUrl {
		case util.TypeUrl_ServiceEntry:
			return &istio_types.ServiceEntry{
				ObjectMeta: metav1.ObjectMeta{
					Name:      event.Name,
					Namespace: event.Namespace,
				},
				Spec: spec,
			}
		case util.TypeUrl_Sidecar:
			return &istio_types.Sidecar{
				ObjectMeta: metav1.ObjectMeta{
					Name:      event.Name,
					Namespace: event.Namespace,
				},
				Spec: spec,
			}
		}
	} else {
		//TODO LOG
	}
	return nil

}

func (a *Actor) Start() {
	for {
		select {
		case <-a.eventChan:
			for {
				nextEvent := a.mailbox.Next()
				if nextEvent != nil {
					obj := eventToObject(nextEvent)
					if obj != nil {
						// TODO error branch
						_ = a.client.Create(context.TODO(), obj)
					}
				} else {
					// 等待1s，进行二次检查
					time.Sleep(1000)
					if a.mailbox.IsEmpty() {
						break
					}
				}
			}

		case <-a.stop:
			return
		}
	}
}

func (a *Actor) Stop() {
	a.stop <- struct{}{}
}

// 判断事件是否过期，如果过期则返回false，如果未过期则将该事件加入Mailbox中
func (a *Actor) Send(event *model.Event) bool {
	if a.mailbox.IsOutDate(event) {
		return false
	}
	if a.mailbox.IsEmpty() {
		a.mailbox.Put(event)
		a.eventChan <- struct{}{}
	} else {
		a.mailbox.Put(event)
	}
	return true
}
