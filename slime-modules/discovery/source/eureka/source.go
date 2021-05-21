package eureka

import (
	"reflect"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1"
	"time"

	"github.com/gogo/protobuf/proto"
	networking "istio.io/api/networking/v1alpha3"
	"slime.io/slime/slime-framework/util"
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

//var Scope = log.RegisterScope("eureka", "eureka debugging", 0)

func New(eureka v1alpha1.Eureka, outer model.Actor, mappingNamespace string) (meshsource.Source, error) {

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
			meta := model.Meta{
				Name:      service,
				Namespace: s.mappingNamespaces,
				TypeUrl:   util.TypeUrl_ServiceEntry,
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
			meta := model.Meta{
				Name:      service,
				Namespace: s.mappingNamespaces,
				TypeUrl:   util.TypeUrl_ServiceEntry,
			}
			event := buildEvent(model.Add, newEntry, meta)
			s.outer.Send(event)
		} else {
			if !reflect.DeepEqual(oldEntry, newEntry) {
				meta := model.Meta{
					Name:      service,
					Namespace: s.mappingNamespaces,
					TypeUrl:   util.TypeUrl_ServiceEntry,
				}
				// UPDATE
				s.cache[service] = newEntry
				event := buildEvent(model.Update, newEntry, meta)
				s.outer.Send(event)
			}
		}
	}
}

func buildEvent(kind model.EventType, item proto.Message, meta model.Meta) *model.Event {
	return &model.Event{
		EventType: kind,
		Version:   time.Now().Unix(),
		Message:   item,
		Meta:      meta,
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
	s.stop <- struct{}{}
}
