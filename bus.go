// Copyright 2022 Thiago Souza <tcostasouza@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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


