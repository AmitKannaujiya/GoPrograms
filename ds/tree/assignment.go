package tree

import (
	"math"
	s "go-program/ds/stack"
)

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

func IsSymetricTree(root *Tree) bool {
	if root == nil || root.Left == nil && root.Right == nil {
		return true
	}
	return IsSymetricTreeRec(root.Left, root.Left.Right)
}

func IsSymetricTreeRec(left, right *Tree) bool {
	// both are empty tree
	if left == nil && right == nil {
		return true
	}
	// either one is empty
	if left == nil || right == nil {
		return false
	}
	if left.Data != right.Data {
		return false
	}
	return IsSymetricTreeRec(left.Left, right.Right) && IsSymetricTreeRec(left.Right, right.Left)
}

func TreePathSum(root *Tree) int {
	sum := 0
	num := 0
	TreePathSumRec(root, num, &sum)
	return sum
}

func TreePathSumRec(root *Tree, num int, sum *int) int {
	// base case
	if root == nil {
		return *sum
	}
	num = 10*num + root.Data
	if root.Left == nil && root.Right == nil {
		*sum = *sum + num
		return *sum
	}
	// 1 case
	return TreePathSumRec(root.Left, num, sum) + TreePathSumRec(root.Right, num, sum)
	// recursion
}

func ConstructTreeFromInorderPreorder(preorder []int, inorder []int) *Tree {
	preS := 0
	inS := 0
	inE := len(inorder)
	inIndexMap := make(map[int]int, inE)
	for i := inS; i < inE; i++ {
		inIndexMap[inorder[i]] = i
	}
	return constructBinarayTreeFromInorderPreorderRec(preorder, inorder, &preS, inS, inE, inIndexMap)
}

func constructBinarayTreeFromInorderPreorderRec(preorder []int, inorder []int, preS *int, inS, inE int, inIndexMap map[int]int) *Tree {
	// baseCase
	if *preS >= len(preorder) {
		return nil
	}

	if inS > inE {
		return nil
	}
	//
	data := preorder[*preS]
	node := &Tree{
		Data: preorder[*preS],
	}
	in := inIndexMap[data]
	*preS = *preS + 1
	node.Left = constructBinarayTreeFromInorderPreorderRec(preorder, inorder, preS, inS, in-1, inIndexMap)
	node.Right = constructBinarayTreeFromInorderPreorderRec(preorder, inorder, preS, in+1, inE, inIndexMap)
	return node
}

func PreorderTraversal(root *Tree) []int {
	result := []int{}
	PreorderTraversalRec(root, &result)
	return result
}

func PreorderTraversalRec(root *Tree, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Data)
	PreorderTraversalRec(root.Left, result)
	PreorderTraversalRec(root.Right, result)
}

func InorderTraversal(root *Tree) []int {
	result := []int{}
	InorderTraversalRec(root, &result)
	return result
}

func InorderTraversalRec(root *Tree, result *[]int) {
	if root == nil {
		return
	}
	InorderTraversalRec(root.Left, result)
	*result = append(*result, root.Data)
	InorderTraversalRec(root.Right, result)
}

//Consider a special family of Engineers and Doctors. This family has the following rules:

//Everybody has two children. The first child of an Engineer is an Engineer and the second child is a Doctor.
// The first child of a Doctor is a Doctor and the second child is an Engineer. All generations of Doctors and Engineers start with an Engineer.
// We can represent the situation using this diagram:
// 0 : E , 1 : D
func FindProfession(level, pos int) rune {
	p := FindProfessionRec(level, pos)
	if p {
		return 'E'
	} else {
		return 'D'
	}
}
func FindProfessionRec(level, pos int) bool {

	if level == 1 {
		return true // enginer
	}
	if pos%2 == 1 {
		return FindProfessionRec(level-1, (pos+1)/2)
	} else {
		return !FindProfessionRec(level-1, (pos+1)/2)
	}
}

func FindKthSmallestInBST(root *Tree, k int) int {
	stack := s.CreateStack[*Tree]()
	for {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		k--
		root, _ = stack.Pop()
		if k == 0 {
			return root.Data
		}
		root = root.Right
	}
}
