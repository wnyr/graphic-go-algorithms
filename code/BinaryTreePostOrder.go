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
//Post-order 遍历二叉树查找
func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)   // 左子树的递归遍历
	postOrder(root.right) // 右子树的递归遍历
	fmt.Printf("%d, " , root.data)
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
	//构造二叉树
	insert(root, 60 )
	insert(root, 40 )
	insert(root, 20 )
	insert(root, 10 )
	insert(root, 30 )
	insert(root, 50 )
	insert(root, 80 )
	insert(root, 70 )
	insert(root, 90 )
	fmt.Printf("Post-order 遍历二叉树查找 \n" )
	postOrder(root)
}
