package generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxOfTypes(t *testing.T) {
	assert.Equal(t, 9, getMaxOfType([]int{8,4,6,0,1,9}))
	assert.Equal(t, "ram", getMaxOfType([]string{"ab", "ba", "dc", "aaa", "bbb", "cc", "ram", "ddd"}))
}

func TestFilter(t *testing.T) {
	evenNos := filter([]int{1,2,3,4,5,6,7,8,9,0}, func(num int) bool {
		return num % 2 == 0
	})
	fmt.Println(evenNos)
	oddNos := filter([]int{1,2,3,4,5,6,7,8,9,0}, func(num int) bool {
		return num % 2 != 0
	})
	fmt.Println(oddNos)
	oddLenStr := filter([]string{"q", "dfg", "fg", "ijji"}, func(str string) bool {
		return len(str) != 2
	})
	fmt.Println(oddLenStr)
	evenLenStr := filter([]string{"q", "dfg", "fg", "ijji"}, func(str string) bool {
		return len(str) == 2
	})
	fmt.Println(evenLenStr)
}

func TestCOntains(t *testing.T) {
	assert.True(t, contains([]string{"Hello", "World", "Kids", "Love", "papa"}, "papa"))
	assert.True(t, contains([]string{"Hello", "World", "Kids", "Love", "papa"}, "Kids"))
	assert.False(t, contains([]string{"Hello", "World", "Kids", "Love", "papa"}, "kids"))
	assert.False(t, contains([]int{4,5,89,12,34,67}, 32))
	assert.True(t, contains([]int{4,5,89,12,34,67}, 34))

}