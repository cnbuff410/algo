package main

import (
	"strconv"
)

/* 
 * A data structure to save vertext and weight in the tuple.
 * It will be used in adjencet list.
 */
type VW struct {
	vertex int
	weight int
}

type Graph struct {
	vnum int          // Number of vertices
	enum int          // Number of edges
	adj  map[int][]*VW // Adjacent lists
	V    []int        // Vertices
}

/* create a V-vertex graph with no edges */
func NewGraph() *Graph {
	a := make(map[int][]*VW)
	v := make([]int, 0)
	return &Graph{V: v, adj: a, enum: 0, vnum: 0}
}

/* Number of vertices */
func (g *Graph) VNum() int {
	return g.vnum
}

/* Number of edges */
func (g *Graph) ENum() int {
	return g.enum
}

/* Degree of graph */
func (g *Graph) Degree(v int) int {
	return len(g.adj[v])
}

/* Return the weight of given two vertex */
func (g *Graph) Weight(u, v int) int {
        if g.IsConnected(u ,v) {
                for _, n := range g.adj[u] {
                        if n.vertex == v {
                                return n.weight
                        }
                }
        }
        return 0
}

/* Return the whole vertices of the graph */
func (g *Graph) Vertex() []int {
        vs := make([]int, g.vnum)
        for i := 0; i < g.vnum; i++ {
                vs[i] = g.V[i]
        }
        return vs
}

/* Add vertex v */
func (g *Graph) AddVertex(v int) {
	if !g.Contains(v) {
		g.V = append(g.V, v)
		g.vnum++
	}
}

/* Add edge v-w */
func (g *Graph) AddEdge(v, w, weight int) {
	if g.Contains(v) && g.Contains(w) {
		if !g.IsConnected(v, w) {
			g.adj[v] = append(g.adj[v], &VW{w, weight})
			g.adj[w] = append(g.adj[w], &VW{v, weight})
			(g.enum)++
		}
	}
}

/* Delete edge v-w */
func (g *Graph) DelEdge(v int, w int) {
	if g.IsConnected(v, w) {
		var index int
		var n *VW
		for index, n = range g.adj[v] {
			if n.vertex == w {
				break
			}
		}
		println("find i ", index)
		g.adj[v] = append(g.adj[v][:index], g.adj[v][index+1:]...)
		for index, n = range g.adj[w] {
			if n.vertex == v {
				break
			}
		}
		g.adj[w] = append(g.adj[w][:index], g.adj[w][index+1:]...)

		(g.enum)--
	}
}

func (g *Graph) DelVertex(v int) {
	var index, v1 int
	var n, n1 *VW
	// Remove v from the adjacent list of its connected vertex
	for _, n = range g.adj[v] {
		s := g.adj[n.vertex]
		for index, n1 = range s {
			if n1.vertex == v {
				break
			}
		}
		g.adj[n.vertex] = append(s[:index], s[index+1:]...)
	}

	// Remove v from Vertex list
	for index, v1 = range g.V {
		if v == v1 {
			break
		}
	}
	g.V = append(g.V[:index], g.V[index+1:]...)
	g.vnum--
}

/* Vertices adjacent to v */
func (g *Graph) Adj(v int) []*VW {
	return g.adj[v]
}

/* Check if two vertex are connected to each other */
func (g *Graph) IsConnected(s int, v int) bool {
	for _, n := range g.adj[s] {
		if n.vertex == v {
			return true
		}
	}
	return false
}

/* Check if given vertex is in graph already */
func (g *Graph) Contains(s int) bool {
	for _, n := range g.V {
		if n == s {
			return true
		}
	}
	return false
}

/* Merge u to v, where v is in the graph and u is removed */
func (g *Graph) Merge(u int, v int) {
	// Port all u's connection to v
	for _, n := range g.adj[u] {
		if n.vertex == v {
			continue
		}
		if !g.IsConnected(n.vertex, v) {
			g.AddEdge(n.vertex, v, n.weight)
		} else {
			for _, n1 := range g.adj[v] {
				if n1.vertex == n.vertex {
					n1.weight += n.weight
				}
			}
		}
	}

	// Remove all information regarding u
	g.DelVertex(u)
}

func (g *Graph) ToString() string {
	nv := g.VNum()
	s := "Current graph: \n"
	s += strconv.Itoa(nv) + " vertices, " + strconv.Itoa(g.ENum()) + " edges\n"
	for i := 0; i < nv; i++ {
		v := g.V[i]
		s += strconv.Itoa(v) + ": "
		for _, w := range g.Adj(v) {
			s += "(" + strconv.Itoa(w.vertex) + ", " + strconv.Itoa(w.weight) + ") "
		}
		s += "\n"
	}
	return s
}

func main() {
	g := NewGraph()
	g.AddVertex(1)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(8)
	g.AddVertex(10)
	g.AddVertex(11)
	g.AddVertex(13)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 10, 1)
	g.AddEdge(1, 11, 1)
	g.AddEdge(1, 13, 1)
	g.AddEdge(3, 11, 1)
	g.AddEdge(4, 8, 1)
	g.AddEdge(4, 10, 1)
	g.AddEdge(4, 13, 1)
	println(g.ToString())
	g.Merge(4, 1)
	println(g.ToString())
}
