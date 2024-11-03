# Question : [Next Greater Element II - (Leetcode 503)](https://leetcode.com/problems/next-greater-element-ii/description/)

Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]), return the next greater number for every element in nums.

The next greater number of a number x is the first greater number to its traversing-order next in the array, which means you could search circularly to find its next greater number. If it doesn't exist, return -1 for this number.


### Example 1

```
Input: nums = [1,2,1]
Output: [2,-1,2]
Explanation: The first 1's next greater number is 2; 
The number 2 can't find next greater number. 
The second 1's next greater number needs to search circularly, which is also 2.

```

### Example 2

```
Input: nums = [1,2,3,4,3]
Output: [2,3,4,-1,4]

```

**Expected Time Complexity**: O(N)
**Expected Auxiliary Space**: O(N)

### Constraints

-   1<= nums.length <=10^4
-   -10^9 <= nums[i] <= 10^9

<!-- **Follow up:** Can you solve it in *O(n)* time and *O(1)* space? -->

## Solution

```GO
package main
import "fmt"

func nextGreaterElements(nums []int) []int {
    nextGreater := make([]int, len(nums))
    stack := CreateStack[int]()
    for i:= len(nums) - 1; i >= 0; i-- {
        stack.Push(i)
    }
    for i:= len(nums) - 1; i >= 0; i-- {
        nge := -1
        for !stack.IsEmpty() {
			t, _ :=stack.Top()
			if nums[t] <= nums[i] {
				stack.Pop()
			} else {
				break
			}
        }
        if !stack.IsEmpty() {
			t, _ :=stack.Top()
            nge = nums[t]
        }
        stack.Push(i)
        nextGreater[i] = nge
    }
    return nextGreater
}

func main() {
    fmt.Println(nextGreaterElements([]int{1,2,1})) //  print : 2,-1,2
    fmt.Println(nextGreaterElements([]int{1,2,3,4,3})) //  print : 2,3,4,-1,4
    fmt.Println(nextGreaterElements([]int{4, 3, 2, 1})) //  print : -1, 4, 4, 4
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```