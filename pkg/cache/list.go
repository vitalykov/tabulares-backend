package cache

type ListNode[T any] struct {
	val  T
	prev *ListNode[T]
	next *ListNode[T]
}

type List[T any] struct {
	dummy *ListNode[T]
	size  int
}

func NewList[T any]() *List[T] {
	node := ListNode[T]{}
	node.prev = &node
	node.next = &node
	list := List[T]{}
	list.dummy = &node
	return &list
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Front() *ListNode[T] {
	return l.dummy.next
}

func (l *List[T]) Back() *ListNode[T] {
	return l.dummy.prev
}

func (l *List[T]) PushBack(val T) {
	node := ListNode[T]{
		val:  val,
		prev: l.Back(),
		next: l.dummy,
	}
	l.Back().next = &node
	l.dummy.prev = &node
	l.size++
}

func (l *List[T]) PushFront(val T) {
	node := ListNode[T]{
		val:  val,
		prev: l.dummy,
		next: l.Front(),
	}
	l.Front().prev = &node
	l.dummy.next = &node
	l.size++
}

func (l *List[T]) PopBack() {
	preTail := l.Back().prev
	preTail.next = l.dummy
	l.dummy.prev = preTail
	l.size = max(0, l.size-1)
}

func (l *List[T]) PopFront() {
	postHead := l.Front().next
	postHead.prev = l.dummy
	l.dummy.next = postHead
	l.size = max(0, l.size-1)
}

func (l *List[T]) MoveFront(node *ListNode[T]) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	front := l.Front()
	l.dummy.next = node
	node.prev = l.dummy
	node.next = front
	front.prev = node
}

func (l *List[T]) MoveBack(node *ListNode[T]) {
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	back := l.Back()
	l.dummy.prev = node
	node.next = l.dummy
	node.prev = back
	back.next = node
}

func (l *List[T]) Delete(node *ListNode[T]) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}
