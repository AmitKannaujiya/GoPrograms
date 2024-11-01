# Question : [Missing And Repeating - (GFG : )](https://www.geeksforgeeks.org/problems/find-missing-and-repeating2512/1)

Given an unsorted array arr of positive integers. One number a from the set [1, 2,....,n] is missing and one number b occurs twice in the array. Find numbers a and b.

Note: The test cases are generated such that there always exists one missing and one repeating number within the range [1,n].

### Example 1

```
Input: arr[] = [2, 2]
Output: [2, 1]
Explanation: Repeating number is 2 and smallest positive missing number is 1.

```

### Example 2

```
Input: arr[] = [1, 3, 3] 
Output: [3, 2]
Explanation: Repeating number is 3 and smallest positive missing number is 2

```
### Example 3

```
Input: arr[] = [4, 3, 6, 2, 1, 1]
Output: [1, 5]
Explanation: Repeating number is 1 and the missing number is 5.

```

### Example 3

```
Input: arr[] = [6, 5, 8, 7, 1, 4, 1, 3, 2]
Output: [1, 9]
Explanation: Repeating number is 1 and the missing number is 9.

```
### Constraints

-  2 ≤ arr.size() ≤ 10^6
-  1 ≤ arr[i] ≤ arr.size()

## Solution : Using Map 

```GO
package main
import "fmt"


func findMissingAndRepeatingElement(nums []int) []int {
	ans := make([]int , 2)

	for i:=0; i < len(nums); i++ {
		index := ABS(nums[i]) // abs is used because we had marked indexed as negative
		// mark the element as visited , convert it to negative
		if nums[index - 1] > 0 {
			nums[index - 1] *= -1
		} else {
			// if it is already negative means it is seen previously 
			// this will be repeated element
			ans = append(ans, index)
		}
	}
	for i:=0; i < len(nums); i++ {
		if nums[i] > 0 {
			// If some element is positive means 
			// that index is missing
			ans = append(ans, i + 1)
		}
	}
	return ans
}

func ABS(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
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
