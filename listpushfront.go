package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}
