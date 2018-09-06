package tree

import "fmt"

func (node *Node) Traverse() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}
//中序遍历
//f是了个函数
func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}
	// nil带可以调函数，不需要分别判断。
	node.Left.TraverseFunc(f)
	f(node)
	//node.Print()
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
