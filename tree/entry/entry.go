package main

import "Demo01/tree"

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

	root = tree.Node{}
	root.left = &treeNode{}
	//root.left = &treeNode{4,nil,nil}
	root.right = &treeNode{5,nil,nil}
	root.right.left = new (treeNode)
	root.left.right = createNode(2)
	root.right.left.setValue(4)

	root.traverse()

}