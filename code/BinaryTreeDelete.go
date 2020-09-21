package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}
var root *Node = nil
func createNewNode(newData int) *Node {
	var newNode *Node = new(Node)
	newNode.data = newData
	newNode.left = nil
	newNode.right = nil
	return newNode
}
func searchMinValue(node *Node) *Node { //最小值
	if node == nil || node.data == 0 {
		return nil
	}
	if node.left == nil {
		return node
	}
	return searchMinValue(node.left) //递归地寻找最小值 从左子树开始

}
// 按顺序遍历二叉搜索树
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left) // 遍历左子树
	fmt.Printf("%d, ", root.data)
	inOrder(root.right) // 遍历右子树
}
func removeNode(node *Node, newData int) *Node {
	if node == nil {
		return node
	}
	var compareValue = newData - node.data
	if compareValue > 0 {
		node.right = removeNode(node.right, newData)
	} else if compareValue < 0 {
		node.left = removeNode(node.left, newData)
	} else if node.left != nil && node.right != nil {
		//找到右子树的最小节点来替换当前节点 节点
		node.data = searchMinValue(node.right).data
		node.right = removeNode(node.right, node.data)
	} else {
		if node.left != nil {
			node = node.left
		} else {
			node = node.right
		}
	}
	return node
}
func insert(node *Node, newData int) {
	if root == nil {
		root = &Node{data: newData, left: nil, right: nil}
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
	} else if compareValue > 0 {//递归右子树
		if node.right == nil {
			node.right = createNewNode(newData)
		} else {
			insert(node.right, newData)
		}
	}
}
func main() {
	//构建二叉树
	insert(root, 60)
	insert(root, 40)
	insert(root, 20)
	insert(root, 10)
	insert(root, 30)
	insert(root, 50)
	insert(root, 80)
	insert(root, 70)
	insert(root, 90)
	fmt.Printf("\n删除节点: 10 \n")
	removeNode(root, 10)
	fmt.Printf("\nIn-order 遍历二叉树 \n")
	inOrder(root)
	fmt.Printf("\n--------------------------------------------\n")
	fmt.Printf("\n删除节点: 20 \n")
	removeNode(root, 20)
	fmt.Printf("\nIn-order 遍历二叉树 \n")
	inOrder(root)
	fmt.Printf("\n--------------------------------------------\n")
	fmt.Printf("\n删除节点: 40 \n")
	removeNode(root, 40)
	fmt.Printf("\nIn-order 遍历二叉树 \n")
	inOrder(root)
}
