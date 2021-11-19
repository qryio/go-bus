package bus

import "go-bus/tree"

type Handler func(p interface{})

type Bus struct {
	tree.Node
}
