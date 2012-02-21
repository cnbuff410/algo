package main

// Find vertices in g that are reachable from s 
func DirectedDFS(g *digraph, s int) []bool{
    marked := make([]bool, g.V)
    dfs(g, s, marked)
    return marked
}

func dfs(g *digraph, s int, marked []bool) {
    //fmt.Println("Now we dfs vertex ", s)
    marked[s] = true
    for _, v := range g.adj[s] {
        if !marked[v] { dfs(g, v, marked) }
    }
}