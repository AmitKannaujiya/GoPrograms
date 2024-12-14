package context

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContextWithTimeOut(t *testing.T) {
	ContextWithTimeOut()
}

func TestContextWithValue(t *testing.T) {
	//assert.True(t, ContextWithValueExample("amit"))
	assert.False(t, ContextWithValueExample("amit123"))
}

func TestContextWithCancell(t *testing.T) {
	ContextWithCancel()
}

func TestContextWithCancel(t *testing.T) {
	ContextWithCancel()
}