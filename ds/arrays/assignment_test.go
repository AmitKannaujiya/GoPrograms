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