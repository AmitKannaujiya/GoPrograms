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

func TestGraphDFSTraversal1(t *testing.T) {
	graph := Newgraph2()
	graph.addEdge(1, 2, false)
	// 0--> 1--> 2 --> 3
	graph.addEdge(2, 3, false)
	graph.addEdge(3, 0, false)
	// 0--> 3--> 4
	graph.addEdge(1, 4, false)
	graph.addEdge(4, 3, false)
	assert.Equal(t, []int{1,2,3,0,4},  dfsTraversal(graph, 1))
}