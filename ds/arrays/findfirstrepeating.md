# Question : [First Repeating Element - (GFG : )](https://www.geeksforgeeks.org/problems/first-repeating-element4018/1)

Given an array arr[], find the first repeating element. The element should occur more than once and the index of its first occurrence should be the smallest.

**Note**:- The position you return should be according to *1-based* indexing. 

### Example 1

```
Input: arr[] = [1, 5, 3, 4, 3, 5, 6]
Output: 2
Explanation: 5 appears twice and its first appearance is at index 2 which is less than 3 whose first the occurring index is 3.

```

### Example 2

```
Input: arr[] = [1, 2, 3, 4]
Output: -1
Explanation: All elements appear only once so answer is -1.

```
### Time Complexity
*Expected Time Complexity*: O(n)
*Expected Auxilliary Space*: O(n)

### Constraints

-  1 ≤ arr.size() ≤ 10^6
-  0 <= arr[i]<= 10^6

## Solution : Using Map 

```GO
package main
import "fmt"


func firstRepeated(nums []int) int {
	freqMap := make(map[int]int, len(nums))
	for _, num := range nums {
		freqMap[num]++
	}
	for i, num := range nums {
		if freqMap[num] > 1 {
			return i + 1 // 1 based index 
		}
	}
	return -1
}

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1, 5, 3, 4, 3, 5, 6}
    fmt.Println(firstRepeated(nums)) //  print : 2
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```
