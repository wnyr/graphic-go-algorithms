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

func insert(insertPosition int,data string)  {
	var p=head
	var i=0
	//将节点移动到插入位置
	for{
		//原示例代码i>insertPosition-1没有=号，Walnut插入在了3号位置
		//San Francisco->Oakland->Berkeley->Walnut->Fremont->End
		if p.next==nil || i>=insertPosition-1{
			break
		}
		p=p.next
		i++
	}
	var newNode *Node=new(Node)
	newNode.data=data
	newNode.next=p.next//newNode.next指向下一个点
	p.next=newNode//当前的下一点是newNode
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
	fmt.Printf("在index = 2处插入一个新节点Walnut \n")
	insert(2,"Walnut")
	output(head)
}