package main

import (
	"fmt"
)

type DataStructOps interface {
	insert()
	update()
	delete()
	searchForAnItem()
	sort()
	print()
}

type Node struct {
	value int
	next  *Node
}

type List struct {
	node *Node
}

func (list *List) insert(newNode *Node) {

	if list.node == nil {
		list.node = newNode
		return
	}

	crrNode := list.node
	for {
		if crrNode.next == nil {
			crrNode.next = newNode
			break
		} else {
			crrNode = crrNode.next
		}
	}
}

func (list *List) print() {
	result := "["
	if list.node == nil {
		result += "]"
		fmt.Println(result)
		return
	}
	crrNode := list.node
	for {
		result += fmt.Sprintf("%v", crrNode.value)
		if crrNode.next != nil {
			crrNode = crrNode.next
		} else {
			break
		}
	}
	result += "]"
	fmt.Println(result)
}

func Demo() {
	var list List
	list.insert(&Node{1, nil})
	list.insert(&Node{2, nil})
	list.insert(&Node{3, nil})
	list.insert(&Node{4, nil})
	list.insert(&Node{5, nil})

	list.print()
}
