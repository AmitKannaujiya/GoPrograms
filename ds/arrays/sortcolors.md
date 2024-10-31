# Question : [Sort Colors - (Leetcode : 75 )](https://leetcode.com/problems/sort-colors/description/)

Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function

### Example 1

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

```

### Example 2

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

### Constraints

-   n == nums.length
-   1 <= n <= 300
-   nums[i] is either 0, 1, or 2.


## Solution : Counting Method

```GO
package main
import "fmt"

// Count zeros, ones, twos
func sortColors(nums []int) {
	zeros, ones, twos := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			zeros++
		case 1:
			ones++
		case 2:
			twos++
		}
	}
    //insert zeros
    i := 0
    for zeros > 0 {
        nums[i] = 0
        i++
        zeros--
    }
    //insert ones
    for ones > 0 {
        nums[i] = 1
        i++
        ones--
    }
    //insert twos
    for twos > 0 {
        nums[i] = 2
        i++
        twos--
    }
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{2,0,2,1,1,0}
    fmt.Println(sortColors(nums)) //  print : [0,0,1,1,2,2]
}

Time Complexity = O(n + n) = O(2*n) = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```

## Solution : 3 pointer approach

```GO
package main
import "fmt"

// 3 pointers 
func sortColors(nums []int) {
	start, mid, end := 0, 0, len(nums)-1
    // <= is needed when mid is at 1 by swaping 2 and loops end with some pending operations
    // try for [2, 0, 1] case
	for mid <= end {
        // if mid is at 0 , swap it with start as start points to 0
		if nums[mid] == 0 {
			nums[mid], nums[start] = nums[start], nums[mid]
			mid++
			start++
            // if mid is at 1, it is at correct place 
		} else if nums[mid] == 1 {
			mid++
		} else {
            // mid is at 2, swap it with end and just decrease end 
			nums[mid], nums[end] = nums[end], nums[mid]
			end-- // mid is not incremented as it is similar to binary sort and end is at right side 
		}
	}
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{2,0,2,1,1,0}
    fmt.Println(sortColors(nums)) //  print : [0,0,1,1,2,2]
}

Time Complexity = O(n + n) = O(2*n) = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```