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

func removeNode(removePosition int ) {
	var p = head
	var i = 0
	for {
		if p.next == nil || i >= removePosition-1 {
			break
		}
		p = p.next
		i++
	}
	var temp = p.next     // 保存想要删除的节点
	p.next = p.next.next // p.next指向要删除节点的下一节点
	p.next.prev = p
	temp.next = nil // 设置要删除节点.next为null
	temp.prev = nil // 设置要删除节点.prev为null
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
	fmt.Printf("删除 index = 2 的节点: \n" )
	removeNode(2)
	output()
}