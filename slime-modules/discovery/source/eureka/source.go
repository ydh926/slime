package eureka

import (
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	networking "istio.io/api/networking/v1alpha3"
	model2 "istio.io/istio-mcp/pkg/model"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1"
	"slime.io/slime/slime-modules/discovery/model"
	meshsource "slime.io/slime/slime-modules/discovery/source"
)

type source struct {
	cache             map[string]*networking.ServiceEntry
	client            Client
	outer             model.Actor
	stop              chan struct{}
	refreshPeriod     time.Duration
	started           bool
	gatewayModel      bool
	mappingNamespaces string
}

func (s *source) GetMappingNamespace() string {
	return s.mappingNamespaces
}

func (s *source) SetMappingNamespace(namespace string) {
	s.mappingNamespaces = namespace
}

//var Scope = log.RegisterScope("eureka", "eureka debugging", 0)

func New(eureka *v1alpha1.Eureka, outer model.Actor, mappingNamespace string) (meshsource.Source, error) {

	client := NewClient(eureka.Address)
	if client == nil {
		return nil, Error{
			msg: "Init eureka client failed",
		}
	}
	return &source{
		cache:         make(map[string]*networking.ServiceEntry),
		client:        client,
		refreshPeriod: time.Duration(eureka.RefreshPeriod),
		outer:         outer,
		started:       false,
		gatewayModel:  false,
		mappingNamespaces: mappingNamespace,
	}, nil
}

func (s *source) refresh() {
	if s.started {
		return
	}
	defer func() {
		s.started = false
	}()
	s.started = true

	apps, err := s.client.Applications()
	if err != nil {
		//log.Errorf("get eureka app failed: " + err.Error())
		return
	}
	newServiceEntryMap, err := ConvertServiceEntryMap(apps, s.gatewayModel)
	if err != nil {
		//log.Errorf("convert eureka servceentry map failed: " + err.Error())
		return
	}

	for service, oldEntry := range s.cache {
		if _, ok := newServiceEntryMap[service]; !ok {
			meta := model2.ConfigMeta{
				Name:             service,
				Namespace:        s.mappingNamespaces,
				GroupVersionKind: model.GVKServiceEntry,
				ResourceVersion: time.Now().String(),
			}
			// DELETE
			delete(s.cache, service)
			event := buildEvent(model.Delete, oldEntry, meta)
			s.outer.Send(event)
		}
	}

	for service, newEntry := range newServiceEntryMap {
		if oldEntry, ok := s.cache[service]; !ok {
			// ADD
			s.cache[service] = newEntry
			meta := model2.ConfigMeta{
				Name:             service,
				Namespace:        s.mappingNamespaces,
				GroupVersionKind: model.GVKServiceEntry,
			}
			event := buildEvent(model.Add, newEntry, meta)
			s.outer.Send(event)
		} else {
			if !reflect.DeepEqual(oldEntry, newEntry) {
				meta := model2.ConfigMeta{
					Name:             service,
					Namespace:        s.mappingNamespaces,
					GroupVersionKind: model.GVKServiceEntry,
				}
				// UPDATE
				s.cache[service] = newEntry
				event := buildEvent(model.Update, newEntry, meta)
				s.outer.Send(event)
			}
		}
	}
}

func buildEvent(kind model.EventType, item proto.Message, meta model2.ConfigMeta) *model.Event {
	return &model.Event{
		EventType: kind,
		Version:   time.Now().Unix(),
		Config: model2.Config{
			Spec:       item,
			ConfigMeta: meta,
		},
	}
}

func (s *source) Start() {
	go func() {

		ticker := time.NewTicker(s.refreshPeriod)
		defer ticker.Stop()
		for {
			select {
			case <-s.stop:
				return
			case <-ticker.C:
				s.refresh()
			}
		}
	}()
}

func (s *source) Stop() {
	close(s.stop)
}
