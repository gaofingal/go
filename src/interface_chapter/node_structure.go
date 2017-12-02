package main

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func (n *Node) setData(data interface{}) {
	n.data = data
}
func (n *Node) setLe(le *Node) {
	n.le = le
}
func (n *Node) setRi(ri *Node) {
	n.ri = ri
}

func NewNode(left *Node, right *Node) *Node {
	return &Node{left, nil, right}
}

func main() {
	root := NewNode(nil, nil)
	root.setData("root node")
	ri := NewNode(nil, nil)
	ri.setData("right node")
	le := NewNode(nil, nil)
	le.setData("left node")
	root.setLe(le)
	root.setRi(ri)
	fmt.Printf("%t\n", root)
}
