package list

type Deque interface {
	Enqueue(int)
	Dequeue() int
}

type node struct {
	value int
	next  *node
}

type List struct {
	length     int
	head, tail *node
}

func (l *List) Enqueue(value int) {
	n := &node{
		value: value,
	}
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		last := l.tail
		last.next = n
		l.tail = n
	}
	l.length++
}

func (l *List) Dequeue() int {
	value := l.head.value
	l.head = l.head.next
	l.length--
	return value
}
