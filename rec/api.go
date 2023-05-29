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

func (node *Node) isTail() bool {
	if node.next == nil {
		return true
	} else {
		return false
	}
}

type List struct {
	node *Node
}

// func (list *List) delete(newNode *Node) {
// 	if list.node == nil {
// 		return
// 	}

// 	pointerToFind := &newNode
// 	fmt.Println(pointerToFind)

// 	prevNode := list.node
// 	nextNode := prevNode.next
// 	for {
// 		if prevNode != nil {
// 			fmt.Println(crrNode)
// 			if pointerToFind == &crrNode {
// 				fmt.Println(crrNode)
// 				if crrNode.next != nil {
// 					prevNode.next = crrNode.next
// 				} else {
// 					prevNode.next = nil
// 				}
// 				fmt.Println(crrNode)
// 				break
// 			} else {
// 				prevNode = crrNode
// 				crrNode = crrNode.next
// 			}
// 			break
// 		} else {
// 			fmt.Println(crrNode)
// 		}

// 	}

// }

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

	var node3 *Node = &Node{3, nil}

	list.insert(&Node{1, nil})
	list.insert(&Node{2, nil})
	list.insert(node3)
	list.insert(&Node{4, nil})
	list.insert(&Node{5, nil})

	// list.delete(node3)

	list.print()
}
