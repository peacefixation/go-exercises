package main

import (
	"exercises/binary-tree/tree"
	"fmt"
)

func main() {
	data := []int{5, 3, 7, 16, 6, 2, 9}

	t := tree.NewTree(data[0])
	for _, d := range data[1:] {
		t = tree.Insert(t, d)
	}

	tree.PreOrder(t)

	t = tree.Remove(t, 5)

	tree.PreOrder(t)

	fmt.Printf("Exists (20): %v\n", tree.Exists(t, 20))
	fmt.Printf("Exists (3): %v\n", tree.Exists(t, 3))
	fmt.Printf("Exists (16): %v\n", tree.Exists(t, 16))
	fmt.Printf("Exists (4): %v\n", tree.Exists(t, 4))
}
