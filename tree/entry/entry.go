package main

import (
	"Demo01/tree"
	"fmt"
)

type MyTreeNode struct {
	node *tree.Node
}


func (myNode *MyTreeNode) postOrder(){
	if myNode == nil || myNode.node ==nil{
		return
	}
}

func main()  {
	var root tree.Node

	root = tree.Node{Value:3}
	root.Left = &tree.Node{}
	//root.left = &treeNode{4,nil,nil}
	root.Right = &tree.Node{5,nil,nil}
	root.Right.Left = new (tree.Node)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	nodeCount :=0
	root.TraverseFunc(func (node *tree.Node) {
		nodeCount++
	})

	fmt.Println("Node count:",nodeCount)
}