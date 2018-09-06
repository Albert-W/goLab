package tree

import "fmt"

type Node struct {
	//public
	Value       int
	Left, Right *Node
}
//接收者，也是指传递。本质是普通函数
func (node Node) Print() {
	fmt.Print(node.Value, " ")
}
//接收者，只有是指针，才能设置成功。
func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	//nil 拿不到value;
	node.Value = value
}
//工厂函数，返加局部变量的地址，也可以给外部用。
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
