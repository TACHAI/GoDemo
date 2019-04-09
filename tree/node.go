package tree

import "fmt"

type Node struct {
	value int
	left,right *Node
}

func createNode(value int) *Node  {
	return &Node{value:value}
}

func (node Node) Print(){
	fmt.Println(node.value)
}

func (node *Node)SetValue(value int)  {
	node.value=value
}

// 前序遍历
func (node *Node)Tree() {
	if node==nil{
		return
	}
	node.Print()
	node.left.Tree()
	node.right.Tree()
}

// 中序遍历
func (node *Node)Traverse(){
	if node==nil{
		return
	}
	node.left.Tree()
	node.Print()
	node.right.Tree()
}

//  后序遍历
func (node *Node)Traverse1()  {
	if node==nil{
		return
	}
	node.left.Tree()
	node.right.Tree()
	node.Print()

}


func main()  {
	var root Node

	root = Node{value:3}
	root.left = &Node{}
	//root.left = &Node{4,nil,nil}
	root.right = &Node{5,nil,nil}
	root.right.left = new (Node)
	root.left.right = createNode(2)
	root.right.left.SetValue(4)

	root.Traverse()

	root.Tree()

	root.Traverse1()


	//nodes := []Node {
	//	{value: 3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)




	//root.print()
	//root.setValue(100)
	//
	//pRoot :=&root
	//pRoot.print()
	//pRoot.setValue(200)
	//pRoot.print()



}
