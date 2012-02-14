package main

import (
    "strconv"
)

type Graph struct {
    v int       // Number of vertices
    e int       // Number of edges
    adj [][]int   // Adjacency lists
}

/* create a V-vertex graph with no edges */
func New(V int) *Graph {
    s := make([][]int, V)
    length := len(s)
    for i := 0; i < length; i++ {
        s[i] = make([]int, 0)
    }
    return &Graph{v:V, e:0, adj:s}
}

/* Number of vertices */
func (g *Graph) V() int {
    return g.v
}

/* Number of edges */
func (g *Graph) E() int {
    return g.e
}

/* Degree of graph */
func (g *Graph) Degree(v int) int {
    return len(g.adj[v]);
}

/* Add edge v-w */
func (g *Graph) AddEdge(v int, w int) {
    g.adj[v] = append(g.adj[v], w)
    g.adj[w] = append(g.adj[w], v)
    (g.e)++
}

/* Vertices adjacent to v */
func (g *Graph) Adj(v int) []int {
    return g.adj[v]
}

func (g *Graph) IsConnected(s int, v int) bool {
    for _, vertex := range g.adj[s] {
        if vertex == v {return true}
    }
    return false
}

func (g *Graph) ToString() string {
    nv := g.V()
    s := strconv.Itoa(nv) + " vertices, " + strconv.Itoa(g.E()) + " edges\n";
    for v := 0; v <nv; v++ {
        s += strconv.Itoa(v) + ": ";
        for _, w := range g.Adj(v) {
            s += strconv.Itoa(w) + " ";
        }
        s += "\n";
    }
    return s;
}
