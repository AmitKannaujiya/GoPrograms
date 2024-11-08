package tree

import "math"

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

func BuildTree(nums []int) *Tree {
	return buildTreeRec(nums, 0)
}

func buildTreeRec(nums []int, index int) *Tree {
	if index >= len(nums) {
		return nil
	}
	root := &Tree{Data: nums[index]}
	root.Left = buildTreeRec(nums, 2*index+1)
	root.Right = buildTreeRec(nums, 2*index+2)
	return root
}

func LevelOrderTravesal(root *Tree) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	if root.Left == nil && root.Right == nil {
		result = append(result, root.Data)
		return result
	}

	queue := []*Tree{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			result = append(result, node.Data)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

func MaxAtEachLevel(root *Tree) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	if root.Left == nil && root.Right == nil {
		result = append(result, root.Data)
		return result
	}
	queue := []*Tree{root}
	for len(queue) > 0 {
		size := len(queue)
		maxi := math.MinInt
		for i := 0; i < size; i++ {
			node := queue[0]
			maxi = max(maxi, node.Data)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, maxi)
	}
	return result
}

