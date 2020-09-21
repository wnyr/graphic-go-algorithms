package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node=new(Node)
var tail *Node=new(Node)

func initial()  {
	head.data="San Francisco"
	head.next=nil

	var nodeOakland *Node = &Node{data:"Oakland",next: nil}
	head.next=nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley" , next: nil }
	nodeOakland.next = nodeBerkeley

	tail.data="Fremont"
	tail.next=nil
	nodeBerkeley.next = tail
}

func removeNode(removePosition int){
	var p=head
	var i=0
	//将节点移动到要删除节点的前一位置
	for{
		if p.next==nil || i>=removePosition-1{
			break
		}
		p=p.next
		i++
	}
	var temp=p.next//保存要删除的节点
	p.next=p.next.next//上个节点的next指向要删除节点的下一个节点
	temp.next=nil
}

func output(node *Node)  {
	var p=node
	for{
		if p == nil {
			break
		}
		fmt.Printf("%s->",p.data)
		p=p.next
	}
	fmt.Printf("End\n\n")
}



func main() {
	initial()
	fmt.Printf("删除index=2处的新节点Berkeley \n")
	removeNode(2)
	output(head)
}