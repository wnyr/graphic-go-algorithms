package main

import "fmt"

type Node struct{
	data int
	left *Node
	right *Node
}

var root *Node=nil

func createNewNode(newData int)*Node  {
	var newNode *Node=new(Node)
	newNode.data=newData
	newNode.left=nil
	newNode.right=nil
	return newNode
}

//按顺序遍历二叉树查找
func inOrder(root *Node)  {
	if root==nil{
		return
	}
	inOrder(root.left)//遍历左子树
	fmt.Printf("%d,",root.data)
	inOrder(root.right)//遍历右子树
}
func insert(node *Node,newData int)  {
	if root == nil {
		root = &Node{data: newData, left: nil , right: nil }
		return
	}
	var compareValue = newData - node.data
	//递归左子树，继续查找插入位置
	if compareValue < 0 {
		if node.left == nil {
			node.left = createNewNode(newData)
		} else {
			insert(node.left, newData)
		}
	} else if compareValue > 0 {
		//递归右子树，继续查找插入位置
		if node.right == nil {
			node.right = createNewNode(newData)
		} else {
			insert(node.right, newData)
		}
	}
}

func main() {
	//构造二叉树查找
	insert(root, 60 )
	insert(root, 40 )
	insert(root, 20 )
	insert(root, 10 )
	insert(root, 30 )
	insert(root, 50 )
	insert(root, 80 )
	insert(root, 70 )
	insert(root, 90 )
	fmt.Printf("In-order遍历二叉树查找 \n" )
	inOrder(root)
}