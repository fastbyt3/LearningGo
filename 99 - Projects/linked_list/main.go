package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
}

func newNode(val int) *Node {
	n := &Node{}
	n.next = nil
	n.prev = nil
	n.value = val
	return n
}

func (ll *LinkedList) insertNode(val int) {
	fmt.Println("[*] Inserting node with value=", val)
	n := newNode(val)
	temp := ll.head

	if ll.head == nil {
		ll.head = n
		return
	}

	for ; temp.next != nil; temp = temp.next {

	}

	temp.next = n
	n.prev = temp
}

func (ll *LinkedList) deleteNode(val int) {
	if ll.head == nil {
		fmt.Println("[X] No nodes in Linked List...")
		return
	}

	for curr := ll.head; curr != nil; curr = curr.next {
		if curr.value == val {
			prevNode := curr.prev
			nextNode := curr.next

			prevNode.next = nextNode
			nextNode.prev = prevNode
			return
		}
	}

	fmt.Printf("[!] Node with value=%d doesnt exist..\n", val)
}

func (ll *LinkedList) reverse() {
	var prev, next *Node
	next = nil
	prev = nil

	for curr := ll.head; curr != nil; curr = next {
		next = curr.next
		curr.next = prev
		curr.prev = next
		prev = curr
	}

	ll.head = prev

	ll.print()
}

func (ll *LinkedList) print() {
	if ll.head == nil {
		fmt.Println("[X] No Linked list exists!")
		return
	}

	fmt.Println()
	fmt.Println("Contents of Linked list: ")
	for curr := ll.head; curr != nil; curr = curr.next {
		fmt.Printf("%d -> ", curr.value)
	}
	fmt.Println("Nil")
}

func main() {
	ll := &LinkedList{}
	ll.insertNode(10)
	ll.insertNode(20)
	ll.insertNode(30)
	ll.insertNode(40)
	ll.insertNode(50)
	// ll.deleteNode(20)
	ll.print()
	ll.reverse()

}
