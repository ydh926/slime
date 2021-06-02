package zookeeper

import (
	model2 "istio.io/istio-mcp/pkg/model"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/samuel/go-zookeeper/zk"
	"istio.io/api/networking/v1alpha3"
	"slime.io/slime/slime-modules/discovery/api/v1alpha1"
	"slime.io/slime/slime-modules/discovery/model"
	meshsource "slime.io/slime/slime-modules/discovery/source"
)

//var Scope = log.RegisterScope("zookeeper", "zookeeper debugging", 0)

type ServiceEntryWithMeta struct {
	ServiceEntry *v1alpha3.ServiceEntry
	Meta         model2.ConfigMeta
}

// source is a simplified client interface for listening/getting Kubernetes resources in an unstructured way.
type source struct {
	cache                 cmap.ConcurrentMap
	Con                   *zk.Conn
	dubboRegisterRootNode string
	gatewayModel          bool
	outer                 model.Actor
	MappingNamespace      string
	stop                  chan bool
}

func (s *source) GetMappingNamespace() string {
	return s.MappingNamespace
}

func (s *source) SetMappingNamespace(namespace string) {
	s.MappingNamespace = namespace
}

func New(zookeeper *v1alpha1.Zookeeper, outer model.Actor, mappingNamespace string) (meshsource.Source, error) {
	if con, _, err := zk.Connect(zookeeper.Address, time.Duration(zookeeper.Timeout)); err != nil {
		return nil, err
	} else {
		return &source{
			cache:                 cmap.New(),
			Con:                   con,
			dubboRegisterRootNode: zookeeper.Root,
			// TODO
			gatewayModel:     false,
			outer:            outer,
			MappingNamespace: mappingNamespace,
		}, nil
	}
}

func (s *source) Start() {
	s.Watch(s.dubboRegisterRootNode, s.ServiceNodeUpdate, nil)
}

func (s *source) Stop() {
	close(s.stop)
}

const (
	ConsumerNode = "consumers"
	ProviderNode = "providers"
)

func (s *source) Watch(path string, updateFunc func([]string, string), deleteFuc func(string)) {
	children, _, err := s.Con.Children(path)
	if updateFunc != nil || deleteFuc != nil {
		s.doWatch(path, updateFunc, deleteFuc)
	}
	if err != nil {
		//log.Errorf(err.Error())
		return
	}
	for _, child := range children {
		if child == ProviderNode {
			go s.doWatchInstance(path, s.EndpointUpdate, s.ServiceNodeDelete)
		} else {
			if path != "/" {
				s.Watch(path+"/"+child, nil, nil)
			} else {
				s.Watch(path+child, nil, nil)
			}
		}
	}
}

func (s *source) doWatch(path string, updateFunction func([]string, string), deleteFunction func(string)) {
retry:
	children, _, event, err := s.Con.ChildrenW(path)
	if err != nil {
		//log.Errorf("Watch zookeeper  path %s failed , %s", path, err.Error())
	}
	//init update
	updateFunction(children, path)

	for {
		select {
		case watchEvent := <-event:
			switch watchEvent.Type {
			case zk.EventNodeDeleted:
				if deleteFunction != nil {
					deleteFunction(path)
				}
				return
			case zk.EventNodeChildrenChanged:
				if children, _, err := s.Con.Children(path); err == nil {
					updateFunction(children, path)
				}
			default:
				goto retry
			}
		case <-s.stop:
			return
		}
	}
}

func (s *source) doWatchInstance(path string, updateFunction func([]string, []string, string), deleteFunction func(string)) {
retry:
	providersPath := path + "/" + ProviderNode
	consumersPath := path + "/" + ConsumerNode

	providerChildren, _, providerEvent, err := s.Con.ChildrenW(providersPath)
	if err != nil {
		//log.Errorf("Watch zookeeper  path %s failed , %s", providersPath, err.Error())
	}
	consumerChildren, _, consumerEvent, err := s.Con.ChildrenW(consumersPath)
	if err != nil {
		//log.Errorf("Watch zookeeper  path %s failed , %s", consumersPath, err.Error())
	}
	//init update
	updateFunction(providerChildren, consumerChildren, providersPath)

	for {
		select {
		case providerWatchEvent := <-providerEvent:
			switch providerWatchEvent.Type {
			case zk.EventNodeDeleted:
				if deleteFunction != nil {
					deleteFunction(providersPath)
				}
				return
			case zk.EventNodeChildrenChanged:
				if children, _, err := s.Con.Children(providersPath); err == nil {
					if consumers, _, err := s.Con.Children(consumersPath); err == nil {
						updateFunction(children, consumers, providersPath)
					}
				}
			default:
				goto retry
			}
		case consumerWatchEvent := <-consumerEvent:
			if consumerWatchEvent.Type == zk.EventNodeChildrenChanged {
				if children, _, err := s.Con.Children(providersPath); err == nil {
					if consumers, _, err := s.Con.Children(consumersPath); err == nil {
						updateFunction(children, consumers, providersPath)
					}
				}

			}
		case <-s.stop:
			return
		}
	}
}

func BuildEvent(eventType model.EventType, item proto.Message, meta model2.ConfigMeta) *model.Event {
	return &model.Event{
		EventType: eventType,
		Config: model2.Config{
			Spec: item,
			ConfigMeta: meta,
		},
		Version:   time.Now().Unix(),
	}
}

func (s *source) ServiceNodeUpdate(services []string, path string) {
	if !verifyPath(path, false) {
		return
	}
	if s.shouldUpdateService(services) {
		for _, child := range services {
			s.Watch(path+"/"+child, nil, nil)
		}
	}
}

func (s *source) ServiceNodeDelete(path string) {
	ss := strings.Split(path, "/")
	service := ss[len(ss)-2]
	if seMap, ok := s.cache.Get(service); ok {
		if ses, ok := seMap.(*cmap.ConcurrentMap); ok {
			for serviceKey, value := range ses.Items() {
				if se, ok := value.(*ServiceEntryWithMeta); ok {
					event := BuildEvent(model.Add, se.ServiceEntry, se.Meta)
					s.outer.Send(event)
					ses.Pop(serviceKey)
				}
			}
		}
		s.cache.Pop(service)
	}
}

func (s *source) EndpointUpdate(provider, consumer []string, path string) {
	if !verifyPath(path, true) {
		return
	}
	ss := strings.Split(path, "/")
	service := ss[len(ss)-2]
	if s.cache == nil {
		s.cache = cmap.New()
	}
	if _, ok := s.cache.Get(service); !ok {
		s.cache.Set(service, cmap.New())
	}
	if seMap, err := convertServiceEntry(provider, consumer, service, s.gatewayModel); err == nil {
		for serviceKey, a := range seMap {
			news := &ServiceEntryWithMeta{
				ServiceEntry: a,
				Meta: model2.ConfigMeta{
					GroupVersionKind: model.GVKServiceEntry,
					Name:      service,
					Namespace: s.MappingNamespace,
					ResourceVersion: time.Now().String(),
				},
			}
			if seCache, ok := s.cache.Get(service); ok {
				ses, castok := seCache.(cmap.ConcurrentMap)
				if castok {
					if value, exist := ses.Get(serviceKey); !exist {
						ses.Set(serviceKey, news)
						event := BuildEvent(model.Add, news.ServiceEntry, news.Meta)
						//log.Infof("add zk se, hosts: %s, ep size: %d ", news.ServiceEntry.Hosts[0], len(news.ServiceEntry.Endpoints))
						s.outer.Send(event)
					} else {
						if v, ok := value.(*ServiceEntryWithMeta); ok {
							oriEndpoints := v.ServiceEntry.Endpoints
							newEndpoints := a.Endpoints
							sortEndpoint(v.ServiceEntry.Endpoints)
							sortEndpoint(a.Endpoints)
							//oriInboundEndpoints := v.ServiceEntry.InboundEndPoints
							//newInboundEndpoints := a.InboundEndPoints
							//sortInboundEndpoint(oriInboundEndpoints)
							//sortInboundEndpoint(newInboundEndpoints)
							if reflect.DeepEqual(oriEndpoints, newEndpoints) { //&& reflect.DeepEqual(oriInboundEndpoints, newInboundEndpoints) {
								return
							} else {
								seMap, _ := s.cache.Get(service)
								if ses, castok := seMap.(*cmap.ConcurrentMap); castok {
									ses.Set(serviceKey, news)
									event := BuildEvent(model.Update, news.ServiceEntry, news.Meta)
									//log.Infof("update zk se, hosts: %s, ep size: %d ", news.ServiceEntry.Hosts[0], len(news.ServiceEntry.Endpoints))
									s.outer.Send(event)
								}
							}
						}
					}
				}
			}
		}
	} else {
		//log.Infof("cause:%s, delete service:%s", err.Error(), service)
		if seCache, ok := s.cache.Get(service); ok {
			if ses, castok := seCache.(cmap.ConcurrentMap); castok {
				for _, value := range ses.Items() {
					if v, ok := value.(*ServiceEntryWithMeta); ok {
						event := BuildEvent(model.Delete, v.ServiceEntry, v.Meta)
						s.cache.Remove(service)
						//log.Infof("delete zk se, hosts: %s, ep size: %d ", v.ServiceEntry.Hosts[0], len(v.ServiceEntry.Endpoints))
						s.outer.Send(event)
					}
				}
			}
		}
	}
}

func sortEndpoint(endpoints []*v1alpha3.WorkloadEntry) {
	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].Address < endpoints[j].Address
	})

}

/*func sortInboundEndpoint(endpoints []*v1alpha3.ServiceEntry_InboundEndPoint) {
	if endpoints == nil || len(endpoints) == 0 {
		return
	}
	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].Address < endpoints[j].Address
	})

}*/

func (s *source) shouldUpdateService(children []string) bool {
	if len(children) != len(s.cache) {
		return true
	} else {
		for _, child := range children {
			if _, ok := s.cache.Get(child); !ok {
				return false
			}
		}
	}
	return true
}

func verifyPath(path string, isInstance bool) bool {
	ss := strings.Split(path, "/")
	if len(ss) < 2 {
		//log.Errorf("Invalid watch path")
		return false
	}
	if isInstance {
		if ss[len(ss)-1] != ConsumerNode && ss[len(ss)-1] != ProviderNode {
			//log.Errorf("Invalid watch path")
			return false
		}
	}
	return true
}
