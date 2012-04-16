package main

import "container/list"

var (
	count int
	id    []int
	m     map[int]bool
)

func SCC(g *Digraph) int {
	count = 0
	marked = make(map[int]bool)
	for _, v := range g.Vertex() {
		marked[v] = false
	}
	id = make([]int, g.VNum())
	order := reversePost(g.Reverse())
	for order.Len() > 0 {
		s := order.Remove(order.Front()).(int)
		if !marked[s] {
			m = make(map[int]bool)
			println("Size of the scc: ", dfsscc(g, s))
			count++
		}
	}

	return count
}

/* Are v and w strongly connected? */
func StronglyConnected(v, w int) bool {
	return id[v] == id[w]
}

func dfsscc(g *Digraph, s int) int {
	marked[s] = true
	id[s] = count
	m[s] = true
	for _, v := range g.Adj(s) {
		if !marked[v.vertex] {
			m[v.vertex] = true
			dfsscc(g, v.vertex)
		}
	}
	return len(m)
}

func reversePost(g *Digraph) *list.List {
	l := list.New()
	l = l.Init()
	marked := make(map[int]bool)
	for _, v := range g.Vertex() {
		marked[v] = false
	}
	for _, v := range g.Vertex() {
		if !marked[v] {
			dfsrp(g, v, marked, l)
		}
	}
	return l
}

func dfsrp(g *Digraph, s int, marked map[int]bool, l *list.List) {
	marked[s] = true
	for _, v := range g.Adj(s) {
		if !marked[v.vertex] {
			dfsrp(g, v.vertex, marked, l)
		}
	}
	l.PushFront(s) // Reverse post
}
