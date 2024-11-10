package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrderTraversal(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tree := BuildTree(nums)
	assert.Equal(t, nums, LevelOrderTravesal(tree))
}

func TestMaxEachLevel(t *testing.T) {
	nums := []int{1, 2, 2, 4, 5, 6, 7, 8}
	tree := BuildTree(nums)
	assert.Equal(t, []int{1, 3, 7, 8}, MaxAtEachLevel(tree))
}

func TestIsSymetricTree(t *testing.T) {
	//nums := []int{1, 2, 3, 4, 2, 4, 3}
	//tree := BuildTree(nums)
	root := &Tree{
		Data: 1,
	}
	root.Left = &Tree{
		Data: 2,
	}
	root.Right = &Tree{
		Data: 2,
	}
	root.Left.Left = &Tree{
		Data: 3,
	}
	root.Left.Right = &Tree{
		Data: 4,
	}
	root.Right.Left = &Tree{
		Data: 4,
	}
	root.Right.Right = &Tree{
		Data: 3,
	}
	assert.Equal(t, true, IsSymetricTree(root))
}

func TestTreePathSum(t *testing.T) {
	// nums := []int{6, 3, 2, 5, 7, 4, 5, 4}
	// tree := BuildTree(nums)
	preorder := []int{6,3,2,5,7,4,5,4}
	inorder :=[]int{2,3,7,5,4,6,5,4}
	tree := ConstructTreeFromInorderPreorder(preorder, inorder)
	assert.Equal(t, 13997, TreePathSum(tree))
}

func TestMain(t *testing.T) {
	preorder := []int{3,9,20,15,7}
	inorder :=[]int{9,3,15,20,7}
	tree := ConstructTreeFromInorderPreorder(preorder, inorder)
	assert.Equal(t, preorder, PreorderTraversal(tree))
	assert.Equal(t, inorder, InorderTraversal(tree))
}
