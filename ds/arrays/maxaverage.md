# Question : [Max Avrage SubArray I - (Leetcode : 643 )](https://leetcode.com/problems/maximum-average-subarray-i/description/)

You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

### Example 1

```
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

```

### Example 2

```
Input: nums = [5], k = 1
Output: 5.00000
```
### Example 3

```
Input: nums = [-1], k = 1
Output: -1.00000
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
func findMaxAverage(nums []int, k int) float64 {
    max := -12121122211.656 // minFloat not available in lib so used some dummy valye
    start, sum := 0, 0
    for end:=0; end < len(nums); end++ {
        // add element in window
        sum += nums[end]
        if end - start + 1 == k {
            // calculate average when window is reached
            avrage := float64(sum) / float64(k)
            // update max
            if max < avrage {
                max = avrage
            }
            // remove element from window
            sum -= nums[start]
            start++
        }
    }
    return max
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1,12,-5,-6,50,3}
    k := 4
    fmt.Println(findMaxAverage(nums, k)) //  print : 12.75000
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```