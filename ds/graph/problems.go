package graph

// import (
// 	g "github.com/supreame/graph"
// 	//q "github.com/supreame/queue"
// )

// func bfsTravesal(graph g.Graph) []int {
// 	visited := make(map[int]struct{})

// 	queue := q.CreateQueue[int]()
// 	queue.Enqueue(0)
// 	visited[0] = struct{}{}
// 	var result []int
// 	for !queue.IsEmpty() {
// 		node := queue.Dequeue()
// 		result = append(result, node)
// 		for _, nbr := range graph.Adj[node] {
// 			if _, exists := visited[nbr]; !exists {
// 				queue.Enqueue(nbr)
// 				visited[nbr] = struct{}{}
// 			}
// 		}
// 	}
// 	return result
// }
// endor lab : interview question
// Prepare a graph Data Structure which follows : 
// checkIfNodes are connected : isconnected 
//1. It should follow Reflexive : a is connected to a
//2. It should follow symetrics : a is connected to b, then b should be connected to a
//3. It should follow transative : a is connected to b, and b is connected to c, then a should be connected to c 
// add edges b/w nodes 

type graph struct {
	adjacencyList map[int]map[int]struct{}
}

func Newgraph() *graph {
	return &graph{
		adjacencyList : make(map[int]map[int]struct{}),
	}
}

func (g *graph) connect(a, b int) {
	if g.adjacencyList[a] == nil {
		g.adjacencyList[a] = make(map[int]struct{})
	}
	if g.adjacencyList[b] == nil {
		g.adjacencyList[b] = map[int]struct{}{}
	}
	g.adjacencyList[a][b] = struct{}{}
	g.adjacencyList[b][a] = struct{}{}
	g.adjacencyList[a][a] = struct{}{}
	g.adjacencyList[b][b] = struct{}{}
}

func (g * graph) isConnected(a, b int) bool {
	// // symetric
	// _, exists := g.adjacencyList[a][b];
	// _, exists2 := g.adjacencyList[b][a];
	// if  !exists ||  !exists2 {
	// 	return false
	// }
	// // reflexive
	// _, exists = g.adjacencyList[a][a];
	// _, exists2 = g.adjacencyList[b][b];
	// if !exists || !exists2 {
	// 	return false
	// }

	// transative : do dfs and check
	visited := make(map[int]struct{})
	return g.dfs(a, b, visited)
}

func (g *graph) dfs(src, dest int, visited map[int]struct{}) bool {
	visited[src] = struct{}{}
	if src == dest {
		return true
	}
	for nbr := range g.adjacencyList[src] {
		if _, exists:=  visited[nbr]; !exists {
			if g.dfs(nbr, dest, visited) {
				return true
			}
		}
	}
	return false
}