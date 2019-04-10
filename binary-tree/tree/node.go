package tree

// Node a tree node
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// NewTree create a new tree consisting of a single node with the given data
func NewTree(data int) *Node {
	return Insert(nil, data)
}

// Insert create a node with the given data and insert it into the tree rooted at the given node
func Insert(node *Node, data int) *Node {
	if node == nil {
		node = &Node{Data: data}
	} else if data < node.Data {
		node.Left = Insert(node.Left, data)
	} else if data > node.Data {
		node.Right = Insert(node.Right, data)
	}

	return node
}

// Remove remove the node with the given data from the tree rooted at the given node
func Remove(root *Node, data int) *Node {
	if root == nil {
		return nil
	} else if data < root.Data {
		root.Left = Remove(root.Left, data)
	} else if data > root.Data {
		root.Right = Remove(root.Right, data)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		root.Data = minValue(root.Right)
		root.Right = Remove(root.Right, root.Data)
	}

	return root
}

func minValue(root *Node) int {
	minV := root.Data
	for root.Left != nil {
		minV = root.Left.Data
		root = root.Left
	}
	return minV
}

// Exists check if the given data exists in the tree
func Exists(root *Node, data int) bool {
	if root == nil {
		return false
	} else if data < root.Data {
		return Exists(root.Left, data)
	} else if data > root.Data {
		return Exists(root.Right, data)
	}

	return true
}
