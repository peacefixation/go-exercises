package tree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	value := 1
	tree := NewTree(value)
	if tree.Data != value {
		t.Errorf("Expected %d found %v", value, tree.Data)
	}
}

func TestInsertTwice(t *testing.T) {
	value1, value2 := 1, 2
	tree := NewTree(value1)
	tree = Insert(tree, value2)

	if tree.Data != value1 {
		t.Errorf("Expected %d found %v", value1, tree.Data)
	}

	if tree.Right == nil {
		t.Fatalf("Expected not nil found %v", tree.Right)
	}

	if tree.Right.Data != value2 {
		t.Errorf("Expected %d found %v", value2, tree.Right.Data)
	}
}

func TestInsertThrice(t *testing.T) {
	value1, value2, value3 := 4, 2, 7
	tree := NewTree(value1)
	tree = Insert(tree, value2)
	tree = Insert(tree, value3)

	if tree.Data != value1 {
		t.Errorf("Expected %d found %v", value1, tree.Data)
	}

	if tree.Left == nil {
		t.Errorf("Expected not nil found %v", tree.Left)
	}

	if tree.Left != nil && tree.Left.Data != value2 {
		t.Errorf("Expected %d found %v", value2, tree.Left.Data)
	}

	if tree.Right == nil {
		t.Errorf("Expected not nil found %v", tree.Right)
	}

	if tree.Right != nil && tree.Right.Data != value3 {
		t.Errorf("Expected %d found %v", value2, tree.Right.Data)
	}
}
