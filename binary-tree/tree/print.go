package tree

import "fmt"

// InOrder print the tree in order
func InOrder(node *Node) {
	if node != nil {
		InOrder(node.Left)
		fmt.Println(node.Data)
		InOrder(node.Right)
	}
}

// PostOrder print the tree in post order
func PostOrder(node *Node) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Println(node.Data)
	}
}

// PreOrder print the tree in pre order
func PreOrder(node *Node) {
	if node != nil {
		fmt.Println(node.Data)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}
