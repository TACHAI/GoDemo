package tree

import "fmt"

type Node struct {
	Value int
	Left,Right *Node
}

func CreateNode(Value int) *Node  {
	return &Node{Value:Value}
}

func (node Node) Print(){
	fmt.Println(node.Value)
}

func (node *Node)SetValue(Value int)  {
	node.Value=Value
}

// 前序遍历
func (node *Node)Tree() {
	if node==nil{
		return
	}
	node.Print()
	node.Left.Tree()
	node.Right.Tree()
}

//// 中序遍历
//func (node *Node)Traverse(){
//	if node==nil{
//		return
//	}
//	node.Left.Tree()
//	node.Print()
//	node.Right.Tree()
//}

//  后序遍历
func (node *Node)Traverse1()  {
	if node==nil{
		return
	}
	node.Left.Tree()
	node.Right.Tree()
	node.Print()

}


func main()  {
	var root Node

	root = Node{Value:3}
	root.Left = &Node{}
	//root.Left = &Node{4,nil,nil}
	root.Right = &Node{5,nil,nil}
	root.Right.Left = new (Node)
	root.Left.Right = CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()

	root.Tree()

	root.Traverse1()


	//nodes := []Node {
	//	{Value: 3},
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
