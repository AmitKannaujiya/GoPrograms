# Question : [Move all negative elements to the left - (GFG :  )]()

Given an unsorted array arr[ ] having both negative and positive integers. The task is to place all negative elements at the left of the array the order of positive elements and negative elements can change.

**Note:** Don't return any array, just in-place on the array.

### Example 1

```
Input : arr[] = [1, -1, 3, 2, -7, -5, 11, 6 ]
Output : [-1, -5, -7, 11, 6, 1, 3, 2]
Explanation: By doing operations we separated the integers without changing the order.

```

### Example 2

```
Input : arr[] = [-5, 7, -3, -4, 9, 10, -1, 11]
Output : [-5, -3, -4, -1, 7, 9, 10, 11]
```

### Constraints

-   1 ≤ arr.size ≤ 10^6
-   -10^9 ≤ arr[i] ≤ 10^9

## Solution : Dutch national Flag algo(2 pointer)

```GO
package main
import "fmt"

// two pointer approach
func moveNegativeAtLeft(nums []int) {
    low, high := 0, len(nums) - 1
    for low < high {
        // low is negative 
        if nums[low] < 0 {
            low++
            // high is positive
        } else if nums[high] > 0 {
            high--
        } else {
            // both are at wrong place, swap them
            nums[low], nums[high] = nums[high], nums[low]
        }
    }
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1,-1,3,2,-7,-5,11,6}
    fmt.Println(moveNegativeAtLeft(nums)) //  print : [-1, -5, -7, 11, 6, 1, 3, 2]
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```