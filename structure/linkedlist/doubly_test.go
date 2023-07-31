package linkedlist

import (
	"reflect"
	"testing"
)

func TestDoubly(t *testing.T) {

	t.Run("Test insert item at end", func(t *testing.T) {

		list := NewDoubly[int]()

		list.Push(1)
		list.Push(2)
		list.Push(3)

		expected := []int{1, 2, 3}
		received := []int{}

		cur := list.Head.Next
		received = append(received, cur.Data)

		for cur.Next != list.Head {
			cur = cur.Next
			received = append(received, cur.Data)
		}

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("receive: %v, expected: %v", received, expected)
		}

	})

	t.Run("Test insert item at the beginning", func(t *testing.T) {

		list := NewDoubly[int]()

		list.Prepend(1)
		list.Prepend(2)
		list.Prepend(3)

		expected := []int{3, 2, 1}
		received := []int{}

		cur := list.Head.Next
		received = append(received, cur.Data)

		for cur.Next != list.Head {
			cur = cur.Next
			received = append(received, cur.Data)
		}

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("receive: %v, expected: %v", received, expected)
		}

	})
}
