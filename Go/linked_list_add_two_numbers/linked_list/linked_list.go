package linked_list

type Node struct {
	Val  any
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil}
}

func (ll *LinkedList) Append(val any) {
	node := &Node{Val: val}
	if ll.Head == nil {
		ll.Head = node
		return
	}

	for current := ll.Head; current != nil; current = current.Next {
		if current.Next == nil {
			current.Next = node
			return
		}
	}
}

func (ll *LinkedList) Remove(val any) bool {
	if ll.Head == nil {
		return false
	}

	prev := ll.Head
	if prev.Val == val {
		ll.Head = prev.Next
		return true
	}

	for current := ll.Head.Next; current != nil; current = current.Next {
		if current.Val == val {
			prev.Next = current.Next
			return true
		}
		prev = current
	}

	return false
}

func (ll *LinkedList) Find(val any) *Node {
	for current := ll.Head; current != nil; current = current.Next {
		if current.Val == val {
			return current
		}
	}

	return nil
}

func (ll *LinkedList) PushFront(val any) {
	node := &Node{Val: val, Next: ll.Head}
	ll.Head = node
}

func (ll *LinkedList) PushBack(val any) {
	node := &Node{Val: val}
	if ll.Head == nil {
		ll.Head = node
		return
	}

	for current := ll.Head; current != nil; current = current.Next {
		if current.Next == nil {
			current.Next = node
			return
		}
	}
}

func (ll *LinkedList) Len() int {
	count := 0
	for current := ll.Head; current != nil; current = current.Next {
		count++
	}
	return count
}

func (ll *LinkedList) Iterate(f func(any)) {
	for current := ll.Head; current != nil; current = current.Next {
		f(current.Val)
	}
}
