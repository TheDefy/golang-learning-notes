package main

import "fmt"

type treeNode struct {
	value int

	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Printf("%d ", node.value)
}

func main() {

	var node01 treeNode = treeNode{value: 3}
	node01.print()
	node := treeNode{}
	t := treeNode{3, nil, nil}
	node.left = &t
	node.print()
	node.left.print()

}
