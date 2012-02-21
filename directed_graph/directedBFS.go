package main

import "container/list"

// Find vertices in g that are reachable from s 
func DirectedBFS(g *digraph, s int) []bool{
    marked := make([]bool, g.V)
    bfs(g, s, marked)
    return marked
}

func bfs(g *digraph, s int, marked []bool) {
    myqueue := list.New() // Use list to simulate a simple queue
    marked[s] = true
    myqueue.PushBack(s)

    for myqueue.Len() > 0 {
        e := myqueue.Front()
        v := e.Value.(int)
        myqueue.Remove(e)
        //fmt.Println("Now we bfs vertex ", v)
        for _, v1 := range g.adj[v] {
            if !marked[v1] { // Not visited
                marked[v1] = true
                myqueue.PushBack(v1)
            }
        }
    }
}
