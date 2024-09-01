package main

import (
	"fmt"
	"math/rand"
)

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

func (n *treeNode) Insert(val int) {
	if val < n.data && n.left == nil {
		n.left = NewTreeNode(val)
		return
	} else if val > n.data && n.right == nil {
		n.right = NewTreeNode(val)
		return
	}

	if val < n.data {
		n.left.Insert(val)
	} else if val > n.data {
		n.right.Insert(val)
	}
}

func (n *treeNode) Print() {
	if n.left != nil {
		n.left.Print()
	}

	fmt.Printf("%d ", n.data)

	if n.right != nil {
		n.right.Print()
	}

}

func (n *treeNode) Search(val int) *treeNode {
	if n == nil {
		return nil
	}

	if val < n.data {
		return n.left.Search(val)
	} else if val > n.data {
		return n.right.Search(val)
	}

	return n
}

func (n *treeNode) minValNode() *treeNode {
	if n == nil {
		return nil
	}

	current := n
	for current.left != nil {
		current = current.left
	}

	return current
}

func (n *treeNode) Delete(val int) *treeNode {
	if n == nil {
		return nil
	}

	if val < n.data {
		n.left = n.left.Delete(val)
	} else if val > n.data {
		n.right = n.right.Delete(val)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}

		minRight := n.right.minValNode()
		n.data = minRight.data
		n.right = n.right.Delete(minRight.data)
	}

	return n
}

func NewTreeNode(data int) *treeNode {
	return &treeNode{data: data}
}

func main() {
	root := NewTreeNode(42)

	fmt.Print("Insert: ")
	for i := 0; i < 10; i++ {
		num := rand.Intn(100)
		fmt.Printf("%d ", num)
		root.Insert(num)
	}
	fmt.Println()
	root.Print()
	root.Delete(42)
	fmt.Println()
	root.Print()
}
