package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func constructGraph(path string) (*Graph, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	graph := NewGraph()
	buf := bufio.NewReader(f)
	line, isPrefix, err := buf.ReadLine()
	for !isPrefix && err == nil {
		s := string(line)
		s = strings.TrimSpace(s)
		strList := strings.Fields(s)
		d1, _ := strconv.Atoi(strList[0])
                graph.AddVertex(d1)
		for i := 1; i < len(strList); i++ {
			d2, _ := strconv.Atoi(strList[i])
                        graph.AddVertex(d2)
			graph.AddEdge(d1, d2, 1)
		}
		line, isPrefix, err = buf.ReadLine()
	}
	if isPrefix {
		return nil, errors.New("buffer size to small")
	}
	return graph, nil
	//if err != nil { return (nil, err) }
}

// Sum of weight for vertex ourside of A
func w(g *Graph, A []int, v int) int {
	weight := 0
	for _, v1 := range A {
		weight += g.Weight(v, v1)
	}
	return weight
}

func mincutPhase(g *Graph, a int) int {
	loop := g.VNum()
	A := make([]int, 0) // Subset of graph vertex
	A = append(A, a)
	VminusA := g.Vertex()

	// Check which vertex outside of A is most tightly connected with A
	for loop > 1 {
		vCandidate := -1
		vCandidateWeight := -1
		// Check all vertices in V - A, pick up the one with biggest w
		for _, v := range VminusA {
			weight := w(g, A, v)
			if weight > vCandidateWeight {
				vCandidate = v
				vCandidateWeight = weight
			}
		}
		// Add to A the most tightly connected vertex
		A = append(A, vCandidate)
		// Del the candidate from V-A
		var i, n int
		for i, n = range VminusA {
			if n == vCandidate {
				break
			}
		}
		VminusA = append(VminusA[:i], VminusA[i+1:]...)
		loop--
	}
	size := len(A)
	t, s := A[size-1], A[size-2]

	// Calculate the cut of this phase
	mincut := 0
	for _, v := range g.Adj(t) {
		mincut += g.Weight(t, v.vertex)
	}
	// Shrink G by merging the two vertices added last
	g.Merge(t, s)
	return mincut
}

func mincutGraph(g *Graph, a int) int {
	mincut := len(g.Adj(a))
	loop := g.VNum()
	for loop > 1 {
		mincutCandidate := mincutPhase(g, a)
		if mincutCandidate < mincut {
			mincut = mincutCandidate
		}
		loop--
	}
	return mincut
}

/*func main() {*/
	/*g, _ := constructGraph("kargerAdj.txt")*/
	/*println("Construct graph finished")*/
	/*println("mincut is ", mincutGraph(g, 1))*/
/*}*/
