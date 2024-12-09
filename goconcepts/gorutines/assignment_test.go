package gorutines

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMinMaxOutput(t *testing.T) {
	data := [][]int{
		{9, 6, 19, 4, 90, 67, 0, 2}, {15, 17, 18, 4, 1, 78}, {89, 1, 7, 18, 67, 189, 89, 4, 56}, {89, 19, 17, 189, 56, 901, 90},
	}
	result := executeTheJob(data)
	fmt.Println(result)
	assert.Equal(t, []int{0, 90}, result[0])
}
