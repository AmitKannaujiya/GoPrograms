# Question : [Find middle of stack - (Assignment : )](https://leetcode.com/problems/maximum-average-subarray-i/description/)

find the middle elment of stack

### Example 1

```
Input: stak = [1,12,-5,-6,50,3]
Output: -6
Explanation: len = 6, mid = 3 , mid elm = -6

```

### Constraints

-    n == nums.length
-    1 <= k <= n <= 10^5
-    -10^4 <= nums[i] <= 10^4


## Solution

```GO
package main
import "fmt"

// Sliding window approach to find maximum
func findMiddleOfStack(stack []int, mid int) int {
	// base case
	if mid == 0 {
		return stack[0]	
	}
	// 1 case
	size := len(stack)
	top := stack[size-1]
	stack = stack[:size-1]
	m := findMiddleOfStack(stack, mid-1)
	stack = append(stack, top)
	return m
}

// Function to initialize the recursion with starting indices of 0
func main() {
    stack := []int{}
    stack = append(stack, 1)
    stack = append(stack, 5)
    stack = append(stack, 6)
    stack = append(stack, 2)
    stack = append(stack, 8)
    mid := len(stack)/2
    fmt.Println(findMiddle(stack, mid)) //  print : 6
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```