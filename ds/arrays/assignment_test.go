package arrays

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPossibleSumRec(t *testing.T) {
	assert.Equal(t, 9, possibleSums([]int{10, 50, 100}, []int{1,2,1}))
}