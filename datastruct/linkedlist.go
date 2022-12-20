package datastructures

// LLNode represents a single linked list node
type LLNode struct {
	Value int
	Next  *LLNode
	Prev  *LLNode
}

type LinkedList struct {
	Length int
	Head   *LLNode
	Tail   *LLNode
}

func (l *LinkedList) InsertAtHead(value int) {
	n := &LLNode{Value: value, Prev: nil}
	switch l.Head {
	case nil:
		l.Head = n
		l.Tail = n
	default:
		l.Head.Prev, n.Next = n, l.Head
		l.Head = n
	}
	l.Length += 1
	return
}

func (l *LinkedList) InsertAtTail(value int) {
	n := &LLNode{Value: value}
	switch l.Tail {
	case nil:
		l.Head = n
		l.Tail = n
	default:
		n.Prev, l.Tail.Next = l.Tail, n
		l.Tail = n
	}
	l.Length += 1
	return
}

func (l *LinkedList) InsertAt(value int, idx int) *LLNode {
	newNode, node := &LLNode{Value: value}, l.GetAt(idx)
	newNode.Prev, newNode.Next, node.Next = node, node.Next, newNode
	return newNode
}

func (l *LinkedList) Insert(value ...int) {
	for _, v := range value {
		l.InsertAtTail(v)
	}
}

func (l *LinkedList) Get(value int) *LLNode {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}

func (l *LinkedList) GetAt(idx int) *LLNode {
	if idx > l.Length-1 {
		return nil
	}
	currentNode, pos := l.Head, 0
	for currentNode != nil && pos < idx {
		currentNode, pos = currentNode.Next, pos+1
	}
	return currentNode
}

func (l *LinkedList) Delete(value int) bool {
	node := l.Get(value)
	if node == nil {
		return false
	}
	node.Prev.Next = node.Next
	node.Prev, node.Next = nil, nil
	return true
}

func (l *LinkedList) DeleteAt(index int) bool {
	ll := l.GetAt(index)
	if ll == nil {
		return false
	}
	prev, next := ll.Prev, ll.Next
	prev.Next, ll = next, nil
	return true
}

func (l *LinkedList) Traverse(callback func(node *LLNode)) {
	currentNode := l.Head
	for currentNode != nil {
		callback(currentNode)
		currentNode = currentNode.Next
	}
}

func (l *LinkedList) ToArray() []int {
	var nums []int
	l.Traverse(func(n *LLNode) {
		nums = append(nums, n.Value)
	})
	return nums
}

func (l *LinkedList) Pop() int {
	switch l.Length {
	case 0:
		return 0
	case 1:
		val := l.Head.Value
		l.Head, l.Tail, l.Length = nil, nil, 0
		return val
	default:
		n := l.Head
		l.Head, l.Head.Next.Prev, l.Length = l.Head.Next, nil, l.Length-1
		return n.Value
	}
}

func (l *LinkedList) Peak() int {
	switch l.Length {
	case 0:
		return 0
	case 1:
		val := l.Head.Value
		l.Head, l.Tail, l.Length = nil, nil, 0
		return val
	default:
		n := l.Tail
		l.Tail, l.Tail.Prev.Next, l.Length = l.Tail.Prev, nil, l.Length-1
		return n.Value
	}
}

func (l *LinkedList) Print() {
	l.Traverse(func(n *LLNode) {
		print(n.Value)
	})
}

func (l *LinkedList) ReversePrint() {
	current := l.Tail

	for current != nil {
		print(current.Value)
		current = current.Prev
	}
}

func NewLinkedList(value ...int) *LinkedList {
	ll := &LinkedList{}
	ll.Insert(value...)
	return ll
}
