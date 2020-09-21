package main
import "fmt"
type Node struct {
	data   int
	left   *Node
	right *Node
}
var root *Node = nil
func createNewNode(newData int ) *Node {
	var newNode *Node = new (Node)
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
	return searchMinValue(node.left) //递归地从左子树中找到最小值
}
func searchMaxValue(node *Node) *Node { //最大值
	if node == nil || node.data == 0 {
		return nil
	}
	if node.right == nil {
		return node
	}
	return searchMaxValue(node.right) //递归地从右子树中找到最小值
}
func insert(node *Node, newData int ) {
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
	//构建二叉树
	insert(root, 60 )
	insert(root, 40 )
	insert(root, 20 )
	insert(root, 10 )
	insert(root, 30 )
	insert(root, 50 )
	insert(root, 80 )
	insert(root, 70 )
	insert(root, 90 )
	fmt.Printf("\n最小值 \n" )
	var minNode = searchMinValue(root)
	fmt.Printf("%d" , minNode.data)
	fmt.Printf("\n最大值 \n" )
	var maxNode = searchMaxValue(root)
	fmt.Printf("%d" , maxNode.data)
}