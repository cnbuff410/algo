package main

import "container/list"

var (
	marked   map[int]bool
	onstack  map[int]bool
	edgeTo   []int
	hasCycle bool
	cycle    *list.List
)

// Find if there is any cycle in given graph. If so, return it, o/w 
// the graph is DAG and return nil
func DAG(g *Digraph) *list.List {
	marked = make(map[int]bool)
        for _, v := range g.Vertex() {
                marked[v] = false
        }
	onstack = make(map[int]bool)
        for _, v := range g.Vertex() {
                onstack[v] = false
        }
	edgeTo = make([]int, g.VNum()*2)
	hasCycle = false
        for _, i := range g.Vertex() {
		edgeTo[i] = -1
	}
        for _, v := range g.Vertex() {
		if c := dagdfs(g, v); c != nil {
			return c
		}
	}
	return nil
}

func dagdfs(g *Digraph, s int) *list.List {
	onstack[s] = true
	marked[s] = true
	for _, vw := range g.adj[s] {
		if hasCycle {
			return cycle
		}
                v := vw.vertex
		if !marked[v] {
			edgeTo[v] = s
			dagdfs(g, v)
		} else if onstack[v] {
			cycle = list.New() // Should use stack. Too lazy here:)
			cycle = cycle.Init()
			for x := s; x != v && x != -1; x = edgeTo[x] {
				cycle.PushFront(x)
			}
                        cycle.PushBack(v)
			cycle.PushFront(v)
			hasCycle = true
		}
	}
	onstack[s] = false
	return nil
}
