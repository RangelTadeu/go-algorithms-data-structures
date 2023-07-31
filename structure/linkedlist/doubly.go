package linkedlist

import "fmt"

type Doubly[T any] struct {
	Head *Node[T]
}

func (ll *Doubly[T]) Init() *Doubly[T] {
	ll.Head = &Node[T]{}
	ll.Head.Prev = ll.Head
	ll.Head.Next = ll.Head

	return ll
}

func NewDoubly[T any]() *Doubly[T] {
	return new(Doubly[T]).Init()
}

func (ll *Doubly[T]) insert(n, at *Node[T]) *Node[T] {
	n.Prev = at
	n.Next = at.Next
	n.Prev.Next = n
	n.Next.Prev = n

	return n
}

func (ll *Doubly[T]) InsertValue(val T, at *Node[T]) *Node[T] {
	return ll.insert(NewNode(val), at)
}

func (ll *Doubly[T]) Push(val T) {
	ll.InsertValue(val, ll.Head.Prev)
}

func (ll *Doubly[T]) Prepend(val T) {
	ll.InsertValue(val, ll.Head)
}

func (ll *Doubly[T]) Print() {
	for cur := ll.Head.Next; cur != ll.Head; cur = cur.Next {
		fmt.Print(cur.Data, ", ")
	}
}
