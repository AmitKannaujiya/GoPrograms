# Question : [Find duplicate Number - (Leetcode : 287)](https://leetcode.com/problems/find-the-duplicate-number/description/)

Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

**Note:** You must solve the problem without modifying the array nums and using only constant extra space.

### Example 1

```
Input: nums = [1,3,4,2,2]
Output: 2

```

### Example 2

```
Input: nums = [3,1,3,4,2]
Output: 3

```
### Example 2

```
Input: nums = [3,3,3,3,3]
Output: 3

```

### Constraints

-  1 <= n <= 10^^5
-  nums.length == n + 1
-  1 <= nums[i] <= n
-  All the integers in nums appear only once except for precisely one integer which appears two or more times.

### Follow up:

-  How can we prove that at least one duplicate number must exist in nums?
-  Can you solve the problem in linear runtime complexity?

## Solution : Using Map 

```GO
package main
import "fmt"


func findDuplicate(nums []int) int {
	freqMap := make(map[int]int, len(nums))
	// count freq of each no
	for _, num := range freqMap {
		freqMap[num]++
		if freqMap[num] > 1 {
			return num
		}
	}
	// return for safe exit
	return -1
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1,3,4,2,2}
    fmt.Println(findDuplicate(nums)) //  print : 2
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```

## Solution : Using cyclic sort 

```GO
package main
import "fmt"


func findDuplicate(nums []int) int {
	i := 0
	n := len(nums)
	for i < n {
		correctIndex := nums[i] - 1 // as nos are 1 to n
		if nums[i] != nums[correctIndex] {
			nums[i], nums[correctIndex] = nums[correctIndex], nums[i]
		} else {
			i++ // move to next index
		}
	}
	return nums[n-1] //  as duplicate no will move to end 
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1,3,4,2,2}
    fmt.Println(findDuplicate(nums)) //  print : 2
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```

## Solution : Using Fast & Slow pointer 

```GO
package main
import "fmt"


// fast and slow pointers
func findDuplicate(nums []int) int {
    if len(nums) > 1 {
        slow := nums[0]
        fast := nums[nums[0]]
        // they will meet on the same number
        // They must meet the same item when slow==fast. In fact, they meet in a circle,
        for fast != slow {
            slow = nums[slow]
            fast = nums[nums[fast]]
        }
        // after meeting the duplicate num will be at the start of the cycle
        //  the duplicate number must be the entry point of the circle when visiting the array from nums[0]. Next we just need to find the entry point. 
        fast = 0
        for slow != fast {
            slow = nums[slow]
            fast = nums[fast]
        }
        return slow
    }
    return -1
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1,3,4,2,2}
    fmt.Println(findDuplicate(nums)) //  print : 2
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```