package xds

import (
	"sync"
	"time"

	"istio.io/istio-mcp/pkg/config/schema/resource"
	"istio.io/istio-mcp/pkg/mcp/xds"
	mcpxds "istio.io/istio-mcp/pkg/mcp/xds"
	model2 "istio.io/istio-mcp/pkg/model"
	"slime.io/slime/slime-modules/discovery/model"
)

type configStore struct {
	lock    sync.RWMutex
	version string
	cache   map[resource.GroupVersionKind]map[string][]model2.Config
}

func (c *configStore) Version() string {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.version
}

func (c *configStore) Configs(gvk resource.GroupVersionKind, namespace string) []model2.Config {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.cache[gvk] == nil {
		return nil
	}
	return c.cache[gvk][namespace]
}

func (c *configStore) List(gvk resource.GroupVersionKind, namespace string) ([]model2.Config, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	if c.cache[gvk] == nil {
		return nil, nil
	}
	return c.cache[gvk][namespace], nil
}

func (c *configStore) Update(event *model.Event) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if c.cache[event.GroupVersionKind] == nil {
		c.cache[event.GroupVersionKind] = make(map[string][]model2.Config)
	}
	if c.cache[event.GroupVersionKind][event.Namespace] == nil {
		c.cache[event.GroupVersionKind][event.Namespace] = make([]model2.Config, 0)
	}
	c.cache[event.GroupVersionKind][event.Namespace] = append(c.cache[event.GroupVersionKind][event.Namespace], event.Config)
}

func (c *configStore) Snapshot() model2.ConfigSnapshot {
	return c
}

func (c *configStore) VersionSnapshot(version string) model2.ConfigSnapshot {
	if c.version != version {
		return nil
	}
	return c
}

type Actor struct {
	mailbox   model.MailBox
	eventChan chan struct{}
	stop      chan struct{}
	server    *xds.Server
	cache     *configStore
	lock      sync.RWMutex
}

func New(url string) *Actor {
	parsed, _ := mcpxds.ParseServerOptions(url)
	a := &Actor{
		mailbox:   model.NewVersionedMailBox(),
		eventChan: make(chan struct{}),
		stop:      make(chan struct{}),
		server:    xds.NewServer(parsed),
		cache: &configStore{
			cache: make(map[resource.GroupVersionKind]map[string][]model2.Config),
		},
	}
	a.server.SetConfigStore(a.cache)
	return a
}

func (a *Actor) Start() {
	a.server.Start(nil)
	for {
		select {
		case <-a.eventChan:
			for {
				nextEvent := a.mailbox.Next()
				if nextEvent != nil {
					a.cache.Update(nextEvent)
				} else {
					// 等待1s，进行二次检查
					time.Sleep(1000)
					if a.mailbox.IsEmpty() {
						break
					}
				}
			}
		case _, ok := <-a.stop:
			if !ok {
				return
			}
		}
	}
}

func (a *Actor) Stop() {
	close(a.stop)
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
