package main

import (
    "testing"
    "fmt"
)

func Test_ugraph(t *testing.T) {
    graph := NewGraph()
    for i:=0; i<13; i++ {
            graph.AddVertex(i)
    }
    graph.AddEdge(0,1,1)
    graph.AddEdge(0,2,1)
    graph.AddEdge(0,5,1)
    graph.AddEdge(0,6,1)
    graph.AddEdge(3,4,1)
    graph.AddEdge(3,5,1)
    graph.AddEdge(4,6,1)
    graph.AddEdge(7,8,1)
    graph.AddEdge(9,10,1)
    graph.AddEdge(9,11,1)
    graph.AddEdge(9,12,1)
    graph.AddEdge(11,12,1)
    fmt.Println(graph.ToString())
    if graph.VNum() != 13 {t.Errorf("Wrong vertices number")}
    if graph.ENum() != 12 {t.Errorf("Wrong edges number")}
    if graph.Degree(1) != 1 {t.Errorf("Wrong degree for vertex 1")}
}
