# Question : [Find middle of stack - (Assignment : )](https://leetcode.com/problems/maximum-average-subarray-i/description/)

find the middle elment of stack

### Example 1

```
Input: stack = [1, 12, 30, 40, 45]
Output: true
Explanation: all element inserted in increasing order

```

### Example 1

```
Input: stack = [1, 12, 30, 40, 35]
Output: false
Explanation: 35 is less than 40

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

func checkSortedRec(stack *[]int, element int) bool {
	old := *stack
	size := len(old)
	// base case
	if size == 0  || size == 1{
		return true
	}
	// 1 case 
	top := old[size - 1]
	if top > element {
		return false
	}
	recAns := checkSortedRec(stack, top)
	// rec
	return recAns
}

func checkSortedIterative(stack []int) bool {
	size := len(stack)
	if size <= 1{
		return true
	}
	top := stack[size-1]
	for size > 0 {
		size--
		next := stack[size -1]
		if top < next {
			return false
		}
		top = next
	}
	return true
}

// Function to initialize the recursion with starting indices of 0
func main() {
    stack := []int{}
    stack = append(stack, 1)
    stack = append(stack, 2)
    stack = append(stack, 3)
    stack = append(stack, 4)
    stack = append(stack, 5)
    fmt.Println(checkSortedIterative(stack)) //  print : true
	fmt.Println(checkSortedRec(&stack, math.MaxInt)) //  print : true
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```