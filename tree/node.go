package tree

import "reflect"

type Visitor func(h interface{})

type Node struct {
	key		 string
	parent	 *Node
	entries  map[uintptr]*Entry
	wildcard *Node
	mapped   map[string]*Node
}

func NewRoot() *Node {
	return makeNode("", nil)
}

func (n *Node) Add(p []string, h interface{}) *Entry {
	if len(p) == 0 {
		e := &Entry{owner: n, handler: h}
		e.key = reflect.ValueOf(e).Pointer()
		n.entries[e.key] = e
		return e
	}
	f, r := p[0], p[1:]
	var c *Node
	if f == wildcardKey {
		if n.wildcard == nil {
			c = makeNode(wildcardKey, n)
			n.wildcard = c
		}
		c = n.wildcard
	} else {
		if _, exists := n.mapped[f]; !exists {
			c = makeNode(f, n)
			n.mapped[f] = c
		}
		c = n.mapped[f]
	}
	return c.Add(r, h)
}

func (n *Node) Accept(p []string, v Visitor) {
	if len(p) == 0 {
		for _, e := range n.entries {
			v(e.handler)
		}
		return
	}
	f, r := p[0], p[1:]
	if n.wildcard != nil {
		n.wildcard.Accept(r, v)
	}
	if m, exists := n.mapped[f]; exists {
		m.Accept(r, v)
	}
}

func (n *Node) prune() {
	if n.isEmpty() && !n.isRoot() {
		p := n.parent
		if p.wildcard == n {
			p.wildcard = nil
		} else {
			delete(p.mapped, n.key)
		}
		p.prune()
	}
}

func (n *Node) isEmpty() bool {
	return len(n.entries) == 0 && len(n.mapped) == 0 && n.wildcard == nil
}

func (n *Node) isRoot() bool {
	return n.parent == nil
}

func makeNode(k string, p *Node) *Node {
	return &Node{
		key: k,
		parent: p,
		entries: map[uintptr]*Entry{},
		mapped: map[string]*Node{},
	}
}

const wildcardKey = "*"