package bus

import (
	"github.com/tsouza/go-bus/tree"
	"sync"
)

type Handler func(t []string, s *Subscription, p interface{})

type Bus struct {
	root  *tree.Node
	mutex sync.RWMutex
}

type Subscription struct {
	bus	  	*Bus
	handler	Handler
	entry 	*tree.Entry
}

func New() *Bus {
	return &Bus{ root: tree.NewRoot() }
}

func (b *Bus) Subscribe(t []string, h Handler) *Subscription {
	s := &Subscription{ bus: b, handler: h }
	b.mutex.Lock()
	e := b.root.Add(t, s)
	s.handler = h
	s.entry = e
	b.mutex.Unlock()
	return s
}

func (b *Bus) Publish(t []string, p interface{}) {
	b.mutex.RLock()
	b.root.Accept(t, func(d interface{}) {
		s := d.(*Subscription)
		s.handler(t, s, p)
	})
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


