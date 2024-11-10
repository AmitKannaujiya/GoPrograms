# Question : [Root to leaf paths sum - (GFG )](https://www.geeksforgeeks.org/problems/root-to-leaf-paths-sum/1)

Given a binary tree, where every node value is a number. Find the sum of all the numbers that are formed from root to leaf paths. The formation of the numbers would be like 10*parent + current (see the examples for more clarification).

### Example 1

```
![alt text](https://media.geeksforgeeks.org/img-practice/prod/addEditProblem/700454/Web/Other/blobid0_1730705508.png)
Output: 13997
Explanation : There are 4 leaves, resulting in leaf path of 632, 6357, 6354, 654 sums to 13997.

```

### Example 2

```
![alt text](https://media.geeksforgeeks.org/img-practice/prod/addEditProblem/700454/Web/Other/blobid1_1730705776.png)
Output: 2630
Explanation: There are 3 leaves, resulting in leaf path of 1240, 1260, 130 sums to 2630.

```

### Example 3

```
Input:    
           1
          /
         2                    
Output: 12
Explanation: There is 1 leaf, resulting in leaf path of 12.

```

### Constraints

-   1 ≤ number of nodes ≤ 31
-   1 ≤ node->data ≤ 100

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