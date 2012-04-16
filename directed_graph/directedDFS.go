package main

// Find vertices in g that are reachable from s 
func DirectedDFS(g *Digraph, s int) map[int]bool {
	marked := make(map[int]bool)
	for _, v := range g.Vertex() {
		marked[v] = false
	}
	dfs(g, s, marked)
	return marked
}

func dfs(g *Digraph, s int, marked map[int]bool) {
	// pre.enqueue(v)  // Pre order
	marked[s] = true
	for _, v := range g.Adj(s) {
		if !marked[v.vertex] {
			dfs(g, v.vertex, marked)
		}
	}
	// post.enqueue(v) // Post order
	// reversePost.push(v) // For topological sort
}
