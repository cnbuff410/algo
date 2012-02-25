package main

import (
    "fmt"
    "strconv"
)

type digraph struct {
    V int
    E int
    adj [][]int
}

func NewDigraph(capacity int) *digraph {
    a := make([][]int, capacity)
    length := len(a)
    for i := 0; i < length; i++ {
        a[i] = make([]int, 0)
    }
    return &digraph{V:capacity, E:0, adj:a}
}

func (g *digraph) AddEdge(src int, dst int) {
    g.E++
    g.adj[src] = append(g.adj[src], dst)
}

func (g *digraph) Adj(v int) []int {
    return g.adj[v]
}

func (g *digraph) Reverse() *digraph {
    newg := NewDigraph(g.V)
    for i := 0; i < g.V; i++ {
        for _, v := range g.adj[i] {
            newg.adj[v] = append(newg.adj[v], i)
        }
    }
    return newg
}

func (g *digraph) ToString() string {
    s := strconv.Itoa(g.V) + " vertices, " + strconv.Itoa(g.E) + " edges\n";
    for v := 0; v <g.V; v++ {
        s += strconv.Itoa(v) + ": ";
        for _, w := range g.Adj(v) {
            s += strconv.Itoa(w) + " ";
        }
        s += "\n";
    }
    return s;
}

func main() {
    // Construction
    g := NewDigraph(13)
    g.AddEdge(4,2)
    g.AddEdge(2,3)
    g.AddEdge(3,2)
    g.AddEdge(6,0)
    g.AddEdge(0,5)
    g.AddEdge(2,0)
    g.AddEdge(11,12)
    g.AddEdge(12,9)
    g.AddEdge(9,10)
    g.AddEdge(9,11)
    g.AddEdge(8,9)
    g.AddEdge(10,12)
    g.AddEdge(11,4)
    g.AddEdge(4,3)
    g.AddEdge(3,5)
    g.AddEdge(7,8)
    g.AddEdge(8,7)
    g.AddEdge(5,4)
    g.AddEdge(0,1)
    g.AddEdge(6,4)
    g.AddEdge(6,9)
    g.AddEdge(7,6)
    fmt.Println(g.ToString())

    // Test reverse
    //newg := g.Reverse()
    //fmt.Println(newg.ToString())

    // Test DFS
    fmt.Println("Result for DFS")
    fmt.Println(DirectedDFS(g, 0))

    // Test BFS
    fmt.Println("Result for BFS")
    fmt.Println(DirectedBFS(g, 3))

    //Test DAG
    if c := DAG(g); c != nil {
            fmt.Println("Found cycle: ")
            for e := c.Front(); e != nil; e = e.Next() {
                    fmt.Printf("%v -> ", e.Value)
            }
            fmt.Printf("\n")
    }
}
