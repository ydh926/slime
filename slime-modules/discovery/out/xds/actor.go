package xds

import (
	"istio.io/istio-mcp/pkg/mcp/xds"
	"slime.io/slime/slime-modules/discovery/model"
)

type Actor struct {
	mailbox   model.MailBox
	eventChan chan struct{}
	stop      chan struct{}
	server    *xds.Server
}

func (a *Actor) Start() {
	a.server.Generators
}

func (a *Actor) Stop() {
	panic("implement me")
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



