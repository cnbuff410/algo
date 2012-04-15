/** Weighted directed graph **/

package main

import (
	"fmt"
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

type Digraph struct {
	vnum int           // Number of vertices
	enum int           // Number of edges
	adj  map[int][]*VW // Adjacent lists
	V    []int         // Vertices
}

func NewDigraph() *Digraph {
	a := make(map[int][]*VW)
	v := make([]int, 0)
	return &Digraph{V: v, adj: a, enum: 0, vnum: 0}
}

/* Number of vertices */
func (g *Digraph) VNum() int {
	return g.vnum
}

/* Number of edges */
func (g *Digraph) ENum() int {
	return g.enum
}

/* Degree of graph */
func (g *Digraph) Degree(v int) int {
	return len(g.adj[v])
}

/* Return the weight of edge from u to v */
func (g *Digraph) Weight(u, v int) int {
	if g.IsConnected(u, v) {
		for _, n := range g.adj[u] {
			if n.vertex == v {
				return n.weight
			}
		}
	}
	return 0
}

/* Return the whole vertices of the graph */
func (g *Digraph) Vertex() []int {
	vs := make([]int, g.vnum)
	for i := 0; i < g.vnum; i++ {
		vs[i] = g.V[i]
	}
	return vs
}

/* Add vertex v */
func (g *Digraph) AddVertex(v int) {
	if !g.Contains(v) {
		g.V = append(g.V, v)
		g.vnum++
	}
}

/* Delete vertex v */
func (g *Digraph) DelVertex(v int) {
	var index, v1 int
	// Remove v from the adjacent list of its connected vertex
	for _, n := range g.Vertex() {
		s := g.adj[n]
		for index = 0; index < len(s); index++ {
			n1 := s[index]
			if n1.vertex == v {
				break
			}
		}
		if index < len(s) { // Find one
			g.adj[n] = append(s[:index], s[index+1:]...)
		}
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

func (g *Digraph) AddEdge(src, dst, weight int) {
	if !g.IsConnected(src, dst) {
		g.adj[src] = append(g.adj[src], &VW{dst, weight})
		g.enum++
	}
}

/* Delete edge from src to dst */
func (g *Digraph) DelEdge(src, dst int) {
	if g.IsConnected(src, dst) {
		var index int
		var n *VW
		for index, n = range g.adj[src] {
			if n.vertex == dst {
				break
			}
		}
		g.adj[src] = append(g.adj[src][:index], g.adj[src][index+1:]...)
		(g.enum)--
	}
}

func (g *Digraph) Adj(v int) []*VW {
	return g.adj[v]
}

func (g *Digraph) Reverse() *Digraph {
	newg := NewDigraph()
	for _, src := range g.V {
                // Add same vertex
                newg.AddVertex(src)
                // Add reversed edge
		for _, dst := range g.adj[src] {
			newg.adj[dst.vertex] = append(newg.adj[dst.vertex], &VW{src, dst.weight})
		}
	}
	return newg
}

/* Check if two vertex are connected to each other */
func (g *Digraph) IsConnected(src, dst int) bool {
	for _, n := range g.adj[src] {
		if n.vertex == dst {
			return true
		}
	}
	return false
}

/* Check if given vertex is in graph already */
func (g *Digraph) Contains(s int) bool {
	for _, n := range g.V {
		if n == s {
			return true
		}
	}
	return false
}

func (g *Digraph) ToString() string {
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
	// Construction
	g := NewDigraph()
	for i := 0; i < 13; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(4, 2, 1)
	g.AddEdge(2, 3, 2)
	g.AddEdge(6, 0, 1)
	g.AddEdge(0, 5, 1)
	g.AddEdge(2, 0, 1)
	g.AddEdge(11, 12, 3)
	g.AddEdge(12, 9, 3)
	g.AddEdge(9, 10, 4)
	g.AddEdge(9, 11, 4)
	g.AddEdge(8, 9, 5)
	g.AddEdge(10, 12, 6)
	g.AddEdge(11, 4, 1)
	g.AddEdge(4, 3, 1)
	g.AddEdge(3, 5, 1)
	g.AddEdge(7, 8, 1)
	g.AddEdge(8, 7, 1)
	g.AddEdge(5, 4, 1)
	g.AddEdge(0, 1, 1)
	g.AddEdge(6, 4, 1)
	g.AddEdge(6, 9, 1)
	g.AddEdge(7, 6, 1)
	fmt.Println(g.ToString())

	// Test reverse
	newg := g.Reverse()
	fmt.Println(newg.ToString())

	// Test DFS
	fmt.Println("Result for DFS")
	fmt.Println(DirectedDFS(g, 6))

	// Test BFS
	fmt.Println("Result for BFS")
	fmt.Println(DirectedBFS(newg, 6))

	//Test DAG
	if c := DAG(g); c != nil {
		fmt.Println("Found cycle: ")
		for e := c.Front(); e != nil; e = e.Next() {
			fmt.Printf("%v -> ", e.Value)
		}
		fmt.Printf("\n")
	}
}
