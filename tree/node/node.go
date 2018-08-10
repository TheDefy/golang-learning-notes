package node

import "fmt"

type TreeNode struct {

	Value int

	Left, Right *TreeNode
}

func (node TreeNode) Print() {
	fmt.Printf("%d ", node.Value)
}
