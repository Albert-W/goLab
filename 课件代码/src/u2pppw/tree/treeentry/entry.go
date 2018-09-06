package main
//一个目录一个包
import (
	"fmt"
	"u2pppw/tree"
	//"imooc.com/ccmouse/learngo/tree"
)
//为了拓展，包含一个指针
type myTreeNode struct {
	node *tree.Node
}
//后续遍历
func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left} //包装一些
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	//tree是包名，Node 是public的结构休；
	var root tree.Node  // 0, nil ,nil;

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node) // new 内建函数，指针也用.
	root.Left.Right = tree.CreateNode(2) //工厂函数，没有构造函数。
	root.Right.Left.SetValue(4)

	fmt.Print("In-order traversal: ")
	root.Traverse()

	//通过包装 组合的方式后序遍历
	fmt.Print("My own post-order traversal: ")
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
	//函数式编程，算结点个数；
	nodeCount := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCount++
	})
	fmt.Println("Node count:", nodeCount)

	//找到最大节点
	//拿到一个channel
	c := root.TraverseWithChannel()
	maxNodeValue := 0
	for node := range c {
		if node.Value > maxNodeValue {
			maxNodeValue = node.Value
		}
	}
	fmt.Println("Max node value:", maxNodeValue)
}
