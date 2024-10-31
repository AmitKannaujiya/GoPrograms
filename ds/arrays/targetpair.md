# Question : [Two Sum - Pair with Given Sum - (GFG : )](https://www.geeksforgeeks.org/problems/key-pair5616/1)

Given an array arr of positive integers and another number target. Determine whether two elements exist in arr whose sum is exactly target or not. Return a boolean value true if two elements exist in arr else return false.

### Example 1

```
Input: arr[] = [1, 4, 45, 6, 10, 8], target =16
Output: true
Explanation: arr[3] + arr[4] = 6 + 10 = 16

```

### Example 2

```
Input: arr[] = [1, 2, 4, 3, 6], target = 11
Output: false
Explanation: None of the pair makes a sum of 11
```

### Constraints

-    1 ≤ arr.size ≤ 10^5
-    1 ≤ arr[i] ≤ 10^5

<!-- **Note**: This recursive solution lead to a time limit exceeded (TLE) error on large inputs because of the exponential time complexity. In the future, we will solve this problem using dynamic programming to optimize the solution and pass all test cases. -->

## Solution

```GO
package main
import "fmt"
    // nums array can be positive and unsorted
    func TargetPairExists(nums []int, target int) bool {
        // create a map to store sum is found or not
        targetMap := make(map[int]struct{})
        // iterate to all nos if target is found then return true
        for i:=0; i< len(nums); i++ {
            remaining := target - nums[i]
            if _, exists := targetMap[remaining]; exists {
                return true
            }
            // store no in map and iterate further
            targetMap[nums[i]] = struct{}{}
        }
        // return false as target is not matched with any no
        return false
    }

// Function to initialize the recursion with starting indices of 0
func main() {
    nums := []int{1, 4, 45, 6, 10, 8}
    target := 16
    fmt.Println(TargetPairExists(nums, target))
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```