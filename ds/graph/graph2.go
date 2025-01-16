package graph

type graph2 struct {
	adjcencyList map[int][]int
}

func Newgraph2() *graph2 {
	return &graph2{
		adjcencyList: make(map[int][]int),
	}
}

func (g *graph2) addEdge(u, v int, directed bool) {
	if _, exists := g.adjcencyList[u]; !exists {
		g.adjcencyList[u] = []int{v}
	} else {
		g.adjcencyList[u] = append(g.adjcencyList[u], v)
	}
	if !directed {
		if _, exists := g.adjcencyList[v]; !exists {
			g.adjcencyList[v] = []int{u}
		} else {
			g.adjcencyList[v] = append(g.adjcencyList[v], u)
		}
	}
}

func dfsTraversal(g *graph2, src int) [] int {
	var result []int
	if g == nil {
		return result
	}
	result = make([]int, 0)
	visited := make(map[int]struct{})
	// handle disconnected component
	// for i:=0; i < v; i++ {
	// 	if _, exists := visited[i]; !exists {
	// 		dfs(i, g.adjcencyList, visited, &result)
	// 	}
	// }
	dfs(src, g.adjcencyList, visited, &result)
	return result
}

func dfs(u int, adj map[int][]int, visited map[int]struct{}, result *[]int) {
	if _, visit := visited[u]; visit {
		return
	}
	visited[u] = struct{}{}
	*result = append(*result, u)
	for _, nbr := range adj[u] {
		if _, visit := visited[nbr]; !visit {
			dfs(nbr, adj, visited, result)
		}
	}
}

