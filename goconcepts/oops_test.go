package goconcepts

import (
	"go-program/goconcepts/inheritence"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInheritence(t *testing.T) {
	assert.Equal(t, "red", inheritence.TestStructInheritence())
}

func TestInheritence2(t *testing.T) {
	assert.Equal(t, "yellow", inheritence.TestInterfaceInheritence())
}

func TestInheritence3(t *testing.T) {
	assert.Equal(t, "green", inheritence.TestInterfaceInheritence2())
}

func TestInheritence4(t *testing.T) {
	assert.Equal(t, "blue", inheritence.TestInterfaceInheritence3())
}

func TestInheritence5(t *testing.T) {
	assert.Equal(t, "lion", inheritence.TestHeirarchicalInheritence())
}
