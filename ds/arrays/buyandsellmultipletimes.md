# Question : [Stock Buy and Sell – Multiple Transaction Allowed - (Leetcode : 287)](https://leetcode.com/problems/find-the-duplicate-number/description/)

The cost of stock on each day is given in an array **price[]**. Each day you may decide to either buy or sell the stock at price[i], you can even buy and sell the stock on the same day. Find the **maximum profit** that you can get.

**Note:** A stock can only be sold if it has been bought previously and multiple stocks cannot be held on any given day.

### Example 1

```
Input: prices[] = [100, 180, 260, 310, 40, 535, 695]
Output: 865
Explanation: Buy the stock on day 0 and sell it on day 3 => 310 – 100 = 210. Buy the stock on day 4 and sell it on day 6 => 695 – 40 = 655. Maximum Profit = 210 + 655 = 865.

```

### Example 2

```
Input: prices[] = [4, 2, 2, 2, 4]
Output: 2
Explanation: Buy the stock on day 3 and sell it on day 4 => 4 – 2 = 2. Maximum Profit = 2.

```

### Constraints

-   1 <= prices.size() <= 105
-   0 <= prices[i] <= 104

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