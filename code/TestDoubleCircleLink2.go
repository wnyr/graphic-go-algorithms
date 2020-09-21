package main

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}
var head *Node = new (Node)
var tail *Node = new (Node)


func initial(){
	head.data = "A"
	head.prev = nil
	head.next = nil

	var nodeB *Node = &Node{data: "B" , prev: head, next: nil }
	head.next = nodeB

	var nodeC *Node = &Node{data: "C" , prev: nodeB, next: nil }
	nodeB.next = nodeC

	tail.data = "D"
	tail.prev = nodeC
	tail.next = head
	nodeC.next = tail
	head.prev = tail
}

func insert(insertPosition int, data string) {
	var p=head
	var i=0
	for{
		if p.next == nil || i >= insertPosition-1 {
			break
		}
		p = p.next
		i++
	}

	var newNode *Node=new(Node)
	newNode.data=data
	newNode.next=p.next
	p.next=newNode
	newNode.prev=p
	newNode.next.prev=newNode
}

func output() {
	var p = head
	for {
		fmt.Printf("%s -> " , p.data)
		p = p.next
		if p == head {
			break
		}
	}
	fmt.Printf("%s " , p.data)
	fmt.Printf("End\n" )
	p = tail
	for {
		fmt.Printf("%s -> " , p.data)
		p = p.prev
		if p == tail {
			break
		}
	}
	fmt.Printf("%s " , p.data)
	fmt.Printf("Start\n\n" )
}
func main() {
	initial()
	fmt.Printf("在索引2处插入一个新节点E : \n" )
	insert(2 , "E" )
	output()
}