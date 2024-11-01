# Question : [Move all negative elements to end - (GFG :  )](https://www.geeksforgeeks.org/problems/move-all-negative-elements-to-end1813/1)

Given an unsorted array arr[ ] having both negative and positive integers. The task is to place all negative elements at the end of the array without changing the order of positive elements and negative elements.

**Note:** Don't return any array, just in-place on the array.

### Example 1

```
Input : arr[] = [1, -1, 3, 2, -7, -5, 11, 6 ]
Output : [1, 3, 2, 11, 6, -1, -7, -5]
Explanation: By doing operations we separated the integers without changing the order.

```

### Example 2

```
Input : arr[] = [-5, 7, -3, -4, 9, 10, -1, 11]
Output : [7, 9, 10, 11, -5, -3, -4, -1]
```

### Constraints

-   1 ≤ arr.size ≤ 10^6
-   -10^9 ≤ arr[i] ≤ 10^9

## Solution : Counting Method

```GO
package main
import "fmt"

// copy positive first, then copy negative
func segregateElements(nums []int) {
	copy := make([]int, len(nums))
    p := 0
	for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            copy[p] = nums[i]
            p++
        }
	}
    for i := 0; i < len(nums); i++ {
        if nums[i] < 0 {
            copy[p] = nums[i]
            p++
        }
	}
    for i := 0; i < len(nums); i++ {
        nums[i] = copy[i]
    }
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1,-1,3,2,-7,-5,11,6}
    fmt.Println(segregateElements(nums)) //  print : [1,3,2,11,6,-1,-7,-5]
}

Time Complexity = O(n + n + n) = O(3*n) = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```