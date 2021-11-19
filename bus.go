package bus

import (
	"bus/tree"
	"sync"
)

type Handler func(p interface{})

type Bus struct {
	root  tree.Node
	mutex sync.RWMutex
}

type Subscription struct {
	bus	  *Bus
	entry *tree.Entry
}

func New() *Bus {
	return &Bus{}
}

func (b *Bus) Subscribe(t []string, h Handler) *Subscription {
	b.mutex.Lock()
	e := b.root.Add(t, h)
	b.mutex.Unlock()
	return &Subscription{b, e}
}

func (b *Bus) Publish(t []string, p interface{}) {
	b.mutex.RLock()
	b.root.Accept(t, func(h interface{}) { h.(Handler)(p) })
	b.mutex.RUnlock()
}

func (s *Subscription) Terminate() {
	if s.entry != nil {
		s.bus.mutex.Lock()
		s.entry.Remove()
		s.entry = nil
		s.bus.mutex.Unlock()
	}
}


