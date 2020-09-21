package main

import (
	"fmt"
)

type Node struct {
	data string
	prev *Node
	next *Node
}

var head *Node=new(Node)
var tail *Node=new(Node)

func initial()  {
	head.data = "San Francisco"
	head.prev = nil
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland" , prev: head, next: nil }
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley" , prev: nodeOakland, next: nil }
	nodeOakland.next = nodeBerkeley

	tail.data = "Fremont"
	tail.prev = nodeBerkeley
	tail.next = nil
	nodeBerkeley.next = tail
}

func removeNode(removePostion int) {
	var p =head
	var i=0
	//将节点移动到要删除的节点的前一节点位置
	for{
		if p.next==nil || i>=removePostion-1{
			break
		}
		p=p.next
		i++
	}
	var temp=p.next//保存要删除的节点
	p.next=p.next.next//p.next指向要删除节点的下一节点
	p.next.prev=p
	temp.next=nil
	temp.prev=nil
}

func output(node *Node){
	var p=node
	var end *Node=nil
	for{
		if p == nil{
			break
		}
		fmt.Printf("%s->",p.data)
		end=p
		p=p.next
	}
	fmt.Printf("End\n")
	p=end
	for{
		if p == nil {
			break
		}
		fmt.Printf("%s->",p.data)
		p=p.prev
	}
	fmt.Printf("Start\n\n")
}

func main() {
	initial()

	fmt.Printf("删除index=2的节点 : \n" )
	removeNode(2 )

	output(head)
}