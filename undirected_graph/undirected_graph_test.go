package main

import (
    "testing"
    "fmt"
)

func Test_ugraph(t *testing.T) {
    graph := New(13)
    graph.AddEdge(0,1)
    graph.AddEdge(0,2)
    graph.AddEdge(0,5)
    graph.AddEdge(0,6)
    graph.AddEdge(3,4)
    graph.AddEdge(3,5)
    graph.AddEdge(4,6)
    graph.AddEdge(7,8)
    graph.AddEdge(9,10)
    graph.AddEdge(9,11)
    graph.AddEdge(9,12)
    graph.AddEdge(11,12)
    fmt.Println(graph.ToString())
    if graph.V() != 13 {t.Errorf("Wrong vertices number")}
    if graph.E() != 12 {t.Errorf("Wrong edges number")}
    if graph.Degree(1) != 1 {t.Errorf("Wrong degree for vertex 1")}
}
