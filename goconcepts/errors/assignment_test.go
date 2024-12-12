package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyError(t *testing.T) {
	assert.Error(t, validateZipCode(100))
	assert.Error(t, validateZipCode(10000))

	fmt.Println(validateZipCode(-100))

	fmt.Println(validateZipCode(10000))
}