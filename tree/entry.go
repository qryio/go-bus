package tree

type Entry struct {
	key		uintptr
	owner   *Node
	handler interface{}
}

func (e *Entry) Remove() {
	if e.owner == nil {
		return
	}
	delete(e.owner.entries, e.key)
	e.owner.prune()
	e.key = 0
	e.owner = nil
	e.handler = nil
}

