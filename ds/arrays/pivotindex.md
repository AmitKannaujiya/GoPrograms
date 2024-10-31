# Question : [Find Pivot Index - (Leetcode : 724 )](https://leetcode.com/problems/find-pivot-index/description/)

Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.

### Example 1

```
Input: nums = [1,7,3,6,5,6]
Output: 3
Explanation:
The pivot index is 3.
Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
Right sum = nums[4] + nums[5] = 5 + 6 = 11

```

### Example 2

```
Input: nums = [1,2,3]
Output: -1
Explanation:
There is no index that satisfies the conditions in the problem statement.
```
### Example 3

```
Input: nums = [2,1,-1]
Output: 0
Explanation:
The pivot index is 0.
Left sum = 0 (no elements to the left of index 0)
Right sum = nums[1] + nums[2] = 1 + -1 = 0
```

### Example 4

```
Input: nums = [2,-2,-1]
Output: 2
Explanation:
The pivot index is 2.
Right sum = 0 (no elements to the right of index 2)
Left sum = nums[0] + nums[1] = 2 + -2 = 0
```
### Example 5

```
Input: nums = [-1]
Output: -1
Explanation:
no left sum exists
```
### Example 6

```
Input: nums = [-1, 1]
Output: -1
Explanation:
no left sum exists
```

### Example 7

```
Input: nums = [-1, 2]
Output: -1
Explanation:
no left sum exists
```

### Constraints

-    1 <= nums.length <= 10^4
-    -1000 <= nums[i] <= 1000

**Note**: This question is the same as 1991: https://leetcode.com/problems/find-the-middle-index-in-array/

## Solution

```GO
package main
import "fmt"
    // nums array can be positive and unsorted
    func pivotIndex(nums []int) int {
        // take the length
        n := len(nums)
        sum := 0
        // calculate total sum
        for i := 0; i < n; i++ {
            sum += nums[i]
        }
        leftSum := 0
        // pivot index will index where leftSum and remaining sum will be equal skiping current num
        for i := 0; i < n; i++ {
            // remainingSum = sum - leftSum
            // skip current element = -nums[i]
            if sum-leftSum-nums[i] == leftSum {
                return i
            }
            leftSum += nums[i]
        }
        // if no such sum exists
        return -1
    }

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1,7,3,6,5,6}
    fmt.Println(pivotIndex(nums)) //  print : 3
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```