# Question : [Missing Number - (Leetcode : 268 )](https://leetcode.com/problems/missing-number/description/)

Given an array nums containing **n distinct** numbers in the range **[0, n]**, return the only number in the range that is missing from the array.

### Example 1

```
Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the missing number in the range since it does not appear in nums.

```

### Example 2

```
Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in the range since it does not appear in nums.
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

## Solution 1

```GO
package main
import "fmt"
// here expected sum of all array will be minus from actual sum
func missingNumberMethod1(nums []int) int {
	sum, n := 0, len(nums)
	// calculate sum
	for _, num := range nums {
		sum += num
	}
	// now missing no will be expected total sum - current sum
	// expected sum for [0,n] will be 
	// expected sum = (n * (n + 1))/2
	return ((n * (n + 1)) / 2) - sum
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{9,6,4,2,3,5,7,0,1}
    fmt.Println(missingNumberMethod1(nums)) //  print : 3
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```
## Solution 2

```GO
package main
import "fmt"
// store occurance of no in array
func missingNumberMethod2(nums []int) int {
	countArray := make([]int, len(nums))
	// count occurance of each no
	for _, num := range nums {
		countArray[num]++
	}
	// check which index does not have any elements
	for index, value := range countArray {
		if value == 0 {
			return index
		}
	}
	// for safe return
    return 0
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{9,6,4,2,3,5,7,0,1}
    fmt.Println(missingNumberMethod2(nums)) //  print : 3
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```

## Solution 3

```GO
package main
import "fmt"
// store occurance of no in array
// cyclic sort : execution of below array
// 2, 3, 1, 5, 0 swap 2 with 1
// 1, 3, 2, 5, 0 swap 1 with 3
// 3, 1, 2, 5, 0 swap 3 with 5
// 5, 1, 2, 3, 0 swap 5 with 0
// 0, 1, 2, 3, 5 now every element is at correct position
// increament i from 0 to 5 
func missingNumberMethod3(nums []int) int {
	i := 0
	for i < len(nums) {
		// here range of nums is [0, n] , so correct position i if array was sorted will be equal to the nums[i]
		correctPosition := nums[i]
		if nums[i] != nums[correctPosition] {  // swap num with correct postion no if it is not matching
			nums[i], nums[correctPosition] = nums[correctPosition], nums[i]
		} else {
			i++   // num is already at correct postion , check next element , so incrementing
		}
	}
	// missing number will the one who's index will not match with the number
	for i:=0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	// if n is missing 
	return len(nums)
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{9,6,4,2,3,5,7,0,1}
    fmt.Println(missingNumberMethod3(nums)) //  print : 3
}

Time Complexity = O(n)
Space Complexity = O(1)
=> where n is the lengths of the array, respectively.
```