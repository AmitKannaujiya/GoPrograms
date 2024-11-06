package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPossibleSumRec(t *testing.T) {
	assert.Equal(t, 9, possibleSums([]int{10, 50, 100}, []int{1, 2, 1}))
}
