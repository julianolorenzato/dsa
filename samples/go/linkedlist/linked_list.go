package linkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	value any
	next  *Node
}

func NewNode(value any) *Node {
	return &Node{value: value, next: nil}
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (ll *LinkedList) PrintList() {
	curr := ll.head

	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func (ll *LinkedList) PushStart(value any) {
	node := NewNode(value)

	if ll.head == nil {
		ll.head = node
	} else {
		temp := ll.head
		ll.head = node
		node.next = temp
	}
}

func (ll *LinkedList) PushEnd(value any) {
	node := NewNode(value)

	if ll.head == nil {
		ll.head = node
	} else {
		curr := ll.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = node
	}
}

func (ll *LinkedList) PopStart() (error, any) {
	if ll.head == nil {
		return errors.New("list is empty"), nil
	} else {
		popped := ll.head
		ll.head = ll.head.next
		popped.next = nil
		return nil, popped
	}
}

func (ll *LinkedList) PopEnd() (error, any) {
	if ll.head == nil {
		return errors.New("list is empty"), nil
	} else {
		curr := ll.head
		for curr.next.next != nil {
			curr = curr.next
		}
		curr.next = nil
		return nil, curr
	}
}

func LinkedListDemo() {
	linkedList := NewLinkedList()

	linkedList.PushStart(2)
	linkedList.PushStart("ABAB")
	linkedList.PushStart(true)
	linkedList.PushEnd(56.8)
	linkedList.PushEnd(false)
	linkedList.PushEnd(map[string]bool{"sim": false})
	linkedList.PrintList()
	linkedList.PopStart()
	linkedList.PopStart()
	linkedList.PrintList()
}
