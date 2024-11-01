# Question : [Reverse Stack - (Assignment : )](https://leetcode.com/problems/maximum-average-subarray-i/description/)

find the middle elment of stack

### Example 1

```
Input: stack = [1, 12, 30, 40, 45]
Output: [45, 40, 30, 12, 1]
Explanation: all elements should be reverted

```

### Example 1

```
Input: stack = [1]
Output: [1]

```

### Constraints

-    n == nums.length
-    1 <= k <= n <= 10^5
-    -10^4 <= nums[i] <= 10^4


## Solution

```GO
package main
import (
	"fmt"
	"math"
)

func insertBottom(stack *[]int, element int) {
	// empty stack insert
	if len(*stack) == 0 {
		*stack = append(*stack, element)
		return
	}
	old := *stack
	top := old[len(old) - 1] // pull top
	*stack = old[:len(old) - 1]
	insertBottom(stack, element) // insert element recursively
	*stack = append(*stack, top) // now put top back in stack
}

func reverseStack(stack *[]int) {
	//base case
	if len(*stack) == 0 {
		return
	}
	old := *stack
	//1 case
	top := old[len(old) - 1] // pull top element
	*stack = old[:len(old) - 1]

	reverseStack(stack) // recursive call for rest of element

	insertBottom(stack, top) // insert element at the end

	// rec
}

// Function to initialize the recursion with starting indices of 0
func main() {
    stack := []int{}
    stack = append(stack, 1)
    stack = append(stack, 2)
    stack = append(stack, 3)
    stack = append(stack, 4)
    stack = append(stack, 5)
    fmt.Println(reverseStack(stack)) //  print : [5, 4, 3, 2, 1]
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```