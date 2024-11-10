# Question : [Construct binary tree from inorder & preorder - (Leetcode - 105 )](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/)

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.

### Example 1

```
![alt text](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

```

### Example 2

```
Input: preorder = [-1], inorder = [-1]
Output: [-1]

```

### Constraints

-   1 <= preorder.length <= 3000
-	inorder.length == preorder.length
-	-3000 <= preorder[i], inorder[i] <= 3000
-	preorder and inorder consist of unique values.
-	Each value of inorder also appears in preorder.
-	preorder is guaranteed to be the preorder traversal of the tree.
-	inorder is guaranteed to be the inorder traversal of the tree.

## Solution

```GO
package main
import "fmt"

func TreePathSum(root *Tree) int {
	sum := 0
	num := 0
	TreePathSumRec(root, num , &sum)
	return sum
}

func TreePathSumRec(root *Tree, num int, sum *int) int {
	// base case
	if root == nil {
		return *sum
	}
	num = 10 * num + root.Data
	if root.Left == nil && root.Right == nil {
		*sum = *sum + num
		return *sum
	}
	// 1 case
	return TreePathSumRec(root.Left, num, sum) + TreePathSumRec(root.Right, num, sum)
	// recursion
}
// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{6, 3, 2, 5,7,4,5,4}
	tree := BuildTree(nums)
    fmt.Println(TreePathSum(tree)) //  print : 13997
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```