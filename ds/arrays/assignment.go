package arrays

import (
	"go-program/ds/utill"
	"math"
	_ "math"
)

// nums array can be positive and unsorted
func TargetPairExists(nums []int, target int) bool {
	// create a map to store sum is found or not
	targetMap := make(map[int]struct{})
	// iterate to all nos if target is found then return true
	for i := 0; i < len(nums); i++ {
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

// find pivot where leftSum is eqaul to right sum excluding current no
func pivotIndex(nums []int) int {
	// take the length
	n := len(nums)
	sum := 0
	// calculate total sum
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	leftSum := 0
	// pivot index will index where leftSum and remaining sum will be equal skiping current num
	for i := 0; i < n; i++ {
		// remainingSum = sum - leftSum
		// skip current element = -nums[i]
		if sum-leftSum-nums[i] == leftSum {
			return i
		}
		leftSum += nums[i]
	}
	// if no such sum exists
	return -1
}

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
		if nums[i] != nums[correctPosition] { // swap num with correct postion no if it is not matching
			nums[i], nums[correctPosition] = nums[correctPosition], nums[i]
		} else {
			i++ // num is already at correct postion , check next element , so incrementing
		}
	}
	// missing number will the one who's index will not match with the number
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	// for safe return
	return 0
}

// remove duplicated
func removeDuplicates(nums []int) int {
	// correct index, should start from zero
	correctIndex := 0
	// loop will start from 1
	for i := 1; i < len(nums); i++ {
		// if no is not duplicate
		if nums[i] != nums[correctIndex] {
			// then place it at
			nums[correctIndex+1] = nums[i]
			correctIndex++
		}
	}
	return correctIndex + 1
}

// Sliding window approach to find maximum
func findMaxAverage(nums []int, k int) float64 {
	max := -12121122211.656 // minFloat not available in lib so used some dummy valye
	start, sum := 0, 0
	for end := 0; end < len(nums); end++ {
		// add element in window
		sum += nums[end]
		if end-start+1 == k {
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

// copy positive first, then copy negative
func segregateElements(nums []int) {
	copy := make([]int, len(nums))
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			copy[p] = nums[i]
			p++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			copy[p] = nums[i]
			p++
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = copy[i]
	}
}

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

// cyclic sort : execution of below array
// 2, 3, 1, 5, 0 swap 2 with 1
// 1, 3, 2, 5, 0 swap 1 with 3
// 3, 1, 2, 5, 0 swap 3 with 5
// 5, 1, 2, 3, 0 swap 5 with 0
// 0, 1, 2, 3, 5 now every element is at correct position
// increament i from 0 till 5 as all elements are at correct position
func findDuplicateCyclicSort(nums []int) int {
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

func findMissingAndRepeatingElement(nums []int) []int {
	ans := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		index := utill.ABS(nums[i]) // math.abs is used because
		// mark the element as visited , convert it to negative
		if nums[index-1] > 0 {
			nums[index-1] *= -1
		} else {
			// if it is already negative means it is seen previously
			// this will be repeated element
			ans = append(ans, index)
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			// If some element is positive means
			// that index is missing
			ans = append(ans, i+1)
		}
	}
	return ans
}

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

func possibleSums(coins []int, quants []int) int {
	sumMap := make(map[int]struct{})
	possibleSumRec(coins, quants, 0, 0, sumMap)
	delete(sumMap, 0)
	return len(sumMap)
}

func possibleSumRec(coins []int, quants []int, index, sum int, sumMap map[int]struct{}) {
	// base case

	if index == len(coins) {
		sumMap[sum] = struct{}{}
		return
	}

	// 1 case
	for i := 0; i <= quants[index]; i++ {
		newSum := sum + coins[index]*i
		possibleSumRec(coins, quants, index+1, newSum, sumMap)
	}

	// recursion
}

func possibleSumIterative(coins []int, quants []int) int {
	sumMap := make(map[int]struct{})
	sumMap[0] = struct{}{}
	for i, coin := range coins {
		count := quants[i]
		currentSUmMap := make(map[int]struct{})
		for s := range sumMap{
			for j:=1; j <= count; j++ {
				cs := s + coin * j
				currentSUmMap[cs] = struct{}{}
			}
		}

		for s := range currentSUmMap {
			sumMap[s] = struct{}{}
		}
	}
	delete(sumMap, 0)
	return len(sumMap)
}

func StockBuyAndSellMultipleTimesApproach1(nums []int) int {
	i, n, lMin, lMax, maxProfit := 0, len(nums), nums[0], nums[0], 0

	for i < n - 1 {
		for i < n - 1 && nums[i] >= nums[i + 1] {  // price is dropping
			i++
		}
		lMin = nums[i] // last one is l min
		for i < n - 1 && nums[i] <= nums[i + 1] {  // price is increasing
			i++
		}
		lMax = nums[i] // last one is l Max
		maxProfit += lMax - lMin
	}
	return maxProfit
}

func StockBuyAndSellMultipleTimesApproach2(prices []int) int {
	maxProfit := 0
	for i:=0; i< len(prices) - 1; i++ {
		if prices[i] < prices[i + 1] {
			maxProfit += prices[i+1] - prices[i]
		}
	}
	return maxProfit
}

func GetMaxMinDiffApproachOne(nums []int, k int) int {
	min, max := math.MaxInt, math.MinInt
	for i:= 0; i < len(nums); i++ {
		if nums[i] - k >= 0 {
			nums[i] -= k
		} else {
			nums[i] += k
		}
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}
	return max - min
}

func GetMaxMinDiffApproachRecursive(nums []int, k int) int {
	min, max := math.MaxInt, math.MinInt
	return GetMaxMinDiffRecursive(nums, k, 0, min, max)
}

func GetMaxMinDiffRecursive(nums []int, k, index, mina, maxa int) int {
	// base case
	if index >= len(nums) {
		return maxa - mina
	}

	// 1 case
	minAns := math.MaxInt
	maxAns := math.MinInt
	if nums[index]- k >=0 {
		nums[index] -= k
		if mina > nums[index] {
			mina = nums[index]
		}
		minAns = GetMaxMinDiffRecursive(nums, k , index + 1, mina, maxa)
	} else {
		nums[index] -= k
		if mina > nums[index] {
			mina = nums[index]
		}
		minAns = GetMaxMinDiffRecursive(nums, k , index + 1, mina, maxa)
		nums[index] += k
		if maxa < nums[index] {
			maxa = nums[index]
		}
		minAns = GetMaxMinDiffRecursive(nums, k , index + 1, mina, maxa)
	}
	return min(maxAns - minAns)


	// recursive
}

func longestSubstringWithKDistinctCharacters(str string, k int) int {
	countMap := make(map[byte]int)

	start, maxLen := 0, 0

	for window:=0; window < len(str); window++ {
		countMap[str[window]]++
		for len(countMap) > k {
			countMap[str[start]]--
			if countMap[str[start]] == 0 {
				delete(countMap, str[start])
			}
			start++
		}
		if len(countMap) == k {
			maxLen = max(maxLen, window - start + 1)
		}
	}
	return maxLen
}