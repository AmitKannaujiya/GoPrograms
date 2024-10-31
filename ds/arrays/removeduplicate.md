# Question : [Remove Duplicates from sorted array - (Leetcode : 26 )](https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/)

Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
Return k

### Example 1

```
Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

```

### Example 2

```
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).
```


### Constraints

-         1 <= nums.length <= 3 * 104
-         -100 <= nums[i] <= 100
-         nums is sorted in **non-decreasing order**.

## Solution

```GO
package main
import "fmt"
// remove duplicated 
func removeDuplicates(nums []int) int {
    // correct index, should start from zero
    correctIndex := 0
    // loop will start from 1
    for i :=1; i < len(nums); i++ {
        // if no is not duplicate
        if nums[i] != nums[correctIndex] {
            // then place it at 
            nums[correctIndex + 1] = nums[i]
            correctIndex++
        }
    }
    return correctIndex + 1
}

// main function
func main() {
    nums := []int{0,0,1,1,1,2,2,3,3,4}
    fmt.Println(removeDuplicates(nums)) //  print : 3
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```