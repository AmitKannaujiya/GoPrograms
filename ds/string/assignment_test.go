package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClubAnagramWord(t *testing.T) {
	words := []string{"eat", "tea", "aaron", "iceman", "cinema", "amit"}
	ans := ClubAnagramWord(words)
	fmt.Println(ans)
	assert.Equal(t, 4, len(ans))
}