package arrays

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
	// for safe return
	return 0
}

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

// Sliding window approach to find maximum
func findMaxAverage(nums []int, k int) float64 {
    max := -12121122211.656 // minFloat not available in lib so used some dummy valye
    start, sum := 0, 0
    for end:=0; end < len(nums); end++ {
        // add element in window
        sum += nums[end]
        if end - start + 1 == k {
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
