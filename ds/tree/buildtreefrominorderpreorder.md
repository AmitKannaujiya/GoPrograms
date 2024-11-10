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

func ConstructTreeFromInorderPreorder(preorder []int, inorder []int) *Tree {
	preS := 0
	inS := 0
	inE := len(inorder)
	inIndexMap := make(map[int]int, inE)
	for i:= inS; i < inE; i++ {
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
		Data :preorder[*preS],
	}
	in := inIndexMap[data]
	*preS = *preS + 1
	node.Left = constructBinarayTreeFromInorderPreorderRec(preorder, inorder, preS, inS, in -1 , inIndexMap)
	node.Right = constructBinarayTreeFromInorderPreorderRec(preorder, inorder, preS, in + 1, inE, inIndexMap)
	return node
}
// Function to initialize the recursion with starting indices of 0
func main() {
    preorder := []int{3,9,20,15,7}
	inorder :=[]int{9,3,15,20,7}
	tree := ConstructTreeFromInorderPreorder(preorder, inorder)
    fmt.Println(preorder, PreorderTraversal(tree)) //  print : 3,9,20,15,7
	fmt.Println(PreorderTraversal(tree)) //  print : 3,9,20,15,7
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```