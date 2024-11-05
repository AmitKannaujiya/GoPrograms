# Question : [Possible Sums made by coins - (CodeSignal : )](https://github.com/kah-ve/codesignal/tree/master/CodingTasks/Hash%20Tables/possibleSums)

You have a collection of coins, and you know the values of the coins and the quantity of each type of coin in it. You want to know how many distinct sums you can make from non-empty groupings of these coins.



**Note**:- The position you return should be according to *1-based* indexing. 

### Example 1

```
Input: coins = [10, 50, 100] and quantity = [1, 2, 1]
Output: 9
Explanation: Here are all the possible sums:

	50 = 50;
	10 + 50 = 60;
	50 + 100 = 150;
	10 + 50 + 100 = 160;
	50 + 50 = 100;
	10 + 50 + 50 = 110;
	50 + 50 + 100 = 200;
	10 + 50 + 50 + 100 = 210;
	10 = 10;
	100 = 100;
	10 + 100 = 110.

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


func possibleSums(coins []int, quants []int) int {
	sumMap := make(map[int]struct{})
	possibleSumRec(coins, quants, 0, 0, sumMap)
	delete(sumMap, 0)
	return len(sumMap)
}

func possibleSumRec(coins []int, quants []int, index , sum int, sumMap map[int]struct{}) {
	// base case

	if index == len(coins) {
		sumMap[sum] = struct{}{}
		return
	}

	// 1 case
	for i:=0; i <= quants[index]; i++ {
		newSum := sum + coins[index] * i
		possibleSumRec(coins, quants, index + 1, newSum, sumMap)
	}

	// recursion
}

// Function to initialize the recursion with starting indices of 0
func main() {
    coins := []int{10, 50, 100}
	quantity := []int{1, 2, 1}
    fmt.Println(possibleSumRec(coins, quantity)) //  print : 9
}

Time Complexity = O(n)
Space Complexity = O(n)
=> where n is the lengths of the array, respectively.
```
