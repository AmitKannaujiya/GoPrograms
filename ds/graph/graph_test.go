package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphConnection(t *testing.T) {
	graph := Newgraph()

	graph.connect(3, 4)
	graph.connect(4, 5)
	assert.True(t, graph.isConnected(3, 4))
	assert.True(t, graph.isConnected(3, 5))
	assert.True(t, graph.isConnected(3, 3))

	assert.False(t, graph.isConnected(3, 8))
}