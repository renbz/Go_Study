package main

import "fmt"

type TreeNode struct {
	No    int
	Name  string
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历 根左右
func PreOrder(node *TreeNode) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历 左根右
func InfixOrder(node *TreeNode) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

// 后序遍历 左右根
func NextOrder(node *TreeNode) {
	if node != nil {
		NextOrder(node.Left)
		NextOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	//构建一个二叉树
	root := &TreeNode{
		No:   1,
		Name: "宋江",
	}

	left1 := &TreeNode{
		No:   2,
		Name: "吴用",
	}

	node10 := &TreeNode{
		No:   10,
		Name: "tom",
	}

	node12 := &TreeNode{
		No:   12,
		Name: "jack",
	}
	left1.Left = node10
	left1.Right = node12

	right1 := &TreeNode{
		No:   3,
		Name: "卢俊义",
	}

	root.Left = left1
	root.Right = right1

	right2 := &TreeNode{
		No:   4,
		Name: "林冲",
	}
	right1.Right = right2

	PreOrder(root)
	//InfixOrder(root) //
	//PostOrder(root) //后序遍历
}
