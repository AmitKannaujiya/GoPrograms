package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPossibleSumRec(t *testing.T) {
	assert.Equal(t, 9, possibleSums([]int{10, 50, 100}, []int{1, 2, 1}))
	assert.Equal(t, 9, possibleSumIterative([]int{10, 50, 100}, []int{1, 2, 1}))
}

func TestStockBuyAndSellMultipleTimesApproach1(t *testing.T) {
	assert.Equal(t, 865, StockBuyAndSellMultipleTimesApproach1([]int{100, 180, 260, 310, 40, 535, 695}))
	assert.Equal(t, 2, StockBuyAndSellMultipleTimesApproach1([]int{4, 2, 2, 2, 4}))
}

func TestStockBuyAndSellMultipleTimesApproach2(t *testing.T) {
	assert.Equal(t, 865, StockBuyAndSellMultipleTimesApproach2([]int{100, 180, 260, 310, 40, 535, 695}))
	assert.Equal(t, 2, StockBuyAndSellMultipleTimesApproach2([]int{4, 2, 2, 2, 4}))
}

func TestGetMaxMinDiffApproachOne(t *testing.T) {
	//assert.Equal(t, 5, GetMaxMinDiffApproachOne([]int{1, 5, 8, 10}, 2))
	assert.Equal(t, 11, GetMaxMinDiffApproachOne([]int{3, 9, 12, 16, 20}, 3))

}

func TestGetMaxMinDiffApproachTwo(t *testing.T) {
	assert.Equal(t, 5, GetMaxMinDiffApproachRecursive([]int{1, 5, 8, 10}, 2))
	assert.Equal(t, 11, GetMaxMinDiffApproachRecursive([]int{3, 9, 12, 16, 20}, 3))

}

func TestLongestSubstringWithKDistinctCharacters(t *testing.T) {
	assert.Equal(t, 5, longestSubstringWithKDistinctCharacters("aabcbbc", 2))
	assert.Equal(t, 7, longestSubstringWithKDistinctCharacters("aabcbbccddcdc", 2))
}

func TestCheckIfArmStrongNo(t *testing.T) {
	assert.Equal(t, true, CheckIfArmStrongNo(153))
	assert.Equal(t, false, CheckIfArmStrongNo(152))
	assert.Equal(t, false, CheckIfArmStrongNo(111))
}