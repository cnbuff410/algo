package main

import "container/list"

// Find vertices in g that are reachable from s 
func DirectedBFS(g *Digraph, s int) map[int]bool {
	marked := make(map[int]bool)
        for _, v := range g.Vertex() {
                marked[v] = false
        }
	bfs(g, s, marked)
	return marked
}

func bfs(g *Digraph, s int, marked map[int]bool) {
	myqueue := list.New() // Use list to simulate a simple queue
	marked[s] = true
	myqueue.PushBack(s)

	for myqueue.Len() > 0 {
		e := myqueue.Front()
		v := e.Value.(int)
		myqueue.Remove(e)
		//fmt.Println("Now we bfs vertex ", v)
		for _, v1 := range g.adj[v] {
			if !marked[v1.vertex] { // Not visited
				marked[v1.vertex] = true
				myqueue.PushBack(v1.vertex)
			}
		}
	}
}
