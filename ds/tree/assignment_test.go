package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrderTraversal(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7,8}
	tree := BuildTree(nums)
	assert.Equal(t, nums, LevelOrderTravesal(tree))
}

func TestMaxEachLevel(t *testing.T) {
	nums := []int{1,2,3,4,5,6,7,8}
	tree := BuildTree(nums)
	assert.Equal(t,[]int{1,3,7,8}, MaxAtEachLevel(tree))
}