# Question : [Generate next greater element - (GFG)](https://www.geeksforgeeks.org/problems/expression-contains-redundant-bracket-or-not/1)

Given a array of nums generate the next greater elements for elements

### Example 1

```
Input: array = [6, 5, 8, 3, 4, 2, 1, 4, 7]
Output: [8, 8, -1, 4, 7, 4, 4, 7, -1]

```

### Example 2

```
Input: array = [1, 2, 3, 4]
Output: [2, 3, 4, -1]

```
### Example 3

```
Input: array = [4, 3, 2, 1]
Output: [-1, -1, -1, -1]

```

**Expected Time Complexity**: O(N)
**Expected Auxiliary Space**: O(N)

### Constraints

-   1<=|str|<=10^4

<!-- **Follow up:** Can you solve it in *O(n)* time and *O(1)* space? -->

## Solution

```GO
package main
import "fmt"

func generateNextGreaterElementSecond(nums []int) []int{
	stack := []int{}
	result := make([]int, len(nums)) 
	for i:=len(nums)-1; i >=0; i-- {
		nge := -1
		size := len(stack)-1
		for size >= 0 && nums[stack[size]] <= nums[i] {
			stack = stack[:size]
			size--
		}
		if size >= 0 {
			nge = nums[stack[size]]
		}
		stack = append(stack, i)
		
		result[i] = nge
	}
	return result
}

func main() {
    fmt.Println(generateNextGreaterElementSecond([]int{1, 2, 3, 4})) //  print : 2, 3, 4, -1
    fmt.Println(generateNextGreaterElementSecond([]int{6, 5, 8, 3, 4, 2, 1, 4, 7})) //  print : 8, 8, -1, 4, 7, 4, 4, 7,-1
    fmt.Println(generateNextGreaterElementSecond([]int{4, 3, 2, 1})) //  print : -1, -1, -1, -1
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```