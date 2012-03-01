package main

import "container/list"
//import "fmt"

var (
    marked []bool
    onstack []bool
    edgeTo []int
    hasCycle bool
    cycle *list.List
    )

// Find if there is any cycle in given graph. If so, return it, o/w 
// the graph is DAG and return nil
func DAG(g *digraph) *list.List {
    marked = make([]bool, g.V)
    onstack = make([]bool, g.V)
    edgeTo = make([]int, g.V)
    hasCycle = false
    for i := 0; i < g.V; i++ {
            edgeTo[i] = -1
    }
    for v := 0; v < g.V; v++ {
            if c := dagdfs(g, v); c != nil {
                    return c
            }
    }
    return nil
}

func dagdfs(g *digraph, s int) *list.List {
    onstack[s] = true
    marked[s] = true
    for _, v := range g.adj[s] {
            if hasCycle {return cycle}
            if !marked[v] {
                    edgeTo[v] = s
                    dagdfs(g, v)
            } else if onstack[v] {
                    cycle = list.New() // Should use stack. Too lazy here:)
                    cycle = cycle.Init()
                    for x := v; x != s && x != -1; x = edgeTo[x] {
                            cycle.PushBack(x)
                    }
                    cycle.PushFront(v)
                    cycle.PushFront(s)
                    hasCycle = true
            }
    }
    onstack[s] = false
    return nil
}
