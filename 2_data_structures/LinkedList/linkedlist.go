package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	tail *node
	size int
}

func (ll *linkedList) push(data int) {
	ll.size += 1
	if ll.head == nil {
		ll.head = &node{data: data}
		ll.tail = ll.head
		return
	}
	temp := node{data: data, next: ll.head}
	ll.head = &temp
}

func (ll *linkedList) pushMany(vals ...int) {
	for _, val := range vals {
		ll.push(val)
	}
}

func (ll *linkedList) pop() (val int) {
	if ll.head == nil {
		panic("Empty LinkedList")
	}
	val = ll.head.data
	ll.head = ll.head.next
	ll.size -= 1
	return
}

func (ll *linkedList) isEmpty() bool {
	return ll.head == nil
}

func (ll *linkedList) print() {
	for temp := ll.head; temp != nil; temp = temp.next {
		fmt.Printf("%d->", temp.data)
	}
	fmt.Print("nil\n")
}

func main() {
	lil := linkedList{}
	lil.pushMany(10, 20, 30, 40)
	lil.print()
	lil.pop()
	lil.print()
}
