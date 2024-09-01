package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func randNumArr(maxSize, maxVal int) []int {
	size := rand.Intn(maxSize-1) + 1
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(maxVal + 1)
	}

	return arr
}

func isValidBST(n *treeNode, min *int, max *int) bool {
	if n == nil {
		return true
	}

	if (min != nil && n.data <= *min) || (max != nil && n.data >= *max) {
		return false
	}

	return isValidBST(n.left, min, &n.data) && isValidBST(n.right, &n.data, max)
}

func TestTreeNodeSearch(t *testing.T) {
	for i := 0; i < 1000; i++ {
		testNums := randNumArr(100, 100)
		root := NewTreeNode(testNums[0])
		for _, n := range testNums[1:] {
			root.Insert(n)
		}
		searchVal := testNums[rand.Intn(len(testNums))]
		resNode := root.Search(searchVal)

		if searchVal != resNode.data {
			t.Errorf("want: %d, got: %d", searchVal, resNode.data)
		}
	}
}

func TestTreeNodeSearchNoMatch(t *testing.T) {
	testNums := []int{1, 5, 66, 0, 87, 73, 23, 43, 23}
	root := NewTreeNode(testNums[0])
	for _, n := range testNums[1:] {
		root.Insert(n)
	}
	searchVal := 99
	want := (*treeNode)(nil)
	got := root.Search(searchVal)
	if want != got {
		t.Errorf("want: %v, got: %v", want, got)
	}
}

func TestTreeNodeInsert(t *testing.T) {
	for i := 0; i < 1000; i++ {
		testNums := randNumArr(100, 100)
		root := NewTreeNode(testNums[0])
		for _, n := range testNums[1:] {
			root.Insert(n)
		}

		// Check if the tree is a valid binary search tree
		if !isValidBST(root, nil, nil) {
			t.Errorf("Insert failed to maintain the binary search tree property")
		}
	}
}

func TestTreeNodeMinValNode(t *testing.T) {
	type fields struct {
		data  int
		left  *treeNode
		right *treeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   *treeNode
	}{
		{
			name: "Single node",
			fields: fields{
				data:  1,
				left:  nil,
				right: nil,
			},
			want: &treeNode{data: 1},
		},
		{
			name: "Two nodes",
			fields: fields{
				data:  1,
				left:  &treeNode{data: 2},
				right: nil,
			},
			want: &treeNode{data: 2},
		},
		{
			name: "Three nodes",
			fields: fields{
				data:  1,
				left:  &treeNode{data: 2},
				right: &treeNode{data: 3},
			},
			want: &treeNode{data: 2},
		},
		{
			name: "Four nodes",
			fields: fields{
				data:  1,
				left:  &treeNode{data: 2, left: &treeNode{data: 4}, right: nil},
				right: &treeNode{data: 3},
			},
			want: &treeNode{data: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &treeNode{
				data:  tt.fields.data,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.minValNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeNode.minValNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeNodeDelete(t *testing.T) {
	root := NewTreeNode(10)
	values := []int{5, 15, 3, 7, 12, 18}
	for _, value := range values {
		root.Insert(value)
	}

	root = root.Delete(10)

	if node := root.Search(10); node != nil {
		t.Errorf("Expected value 10 to be deleted, but it was found")
	}
}
