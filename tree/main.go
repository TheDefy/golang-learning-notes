package main

import "thedefy/LearningNotes/tree/node"

func main() {
	var node01 node.TreeNode = node.TreeNode{Value: 3}
	node01.Print()
	node02 := node.TreeNode{}
	t := node.TreeNode{3, nil, nil}
	node02.Left = &t
	node02.Print()
	node02.Left.Print()
}
