package arrays

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