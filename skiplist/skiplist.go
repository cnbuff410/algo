package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxInt = int(^uint(0) >> 1)
	MinInt = -MaxInt - 1
)

type Node struct {
	Value int
	Next  []*Node
}

type Skiplist struct {
	Head   *Node
	Level  int
	Length int
}

func New(maxLevel int) *Skiplist {
	// Allocation head points
	head := &Node{MinInt, make([]*Node, maxLevel)}
	return &Skiplist{Level: 0, Head: head, Length: 0} // Return header
}

func (sl *Skiplist) Insert(value int) {
	// Calculate level
	level := 0
	t := time.Now()
	rand.Seed((int64)(t.Nanosecond()))
	for rand.Float64() < 0.5 {
		level++
	}
	if level+1 > sl.Level {
		sl.Level = level + 1
	}

	// Init new node
	newNode := &Node{Value: value, Next: make([]*Node, level+1)}

	// Find the place to insert
	for i := sl.Level - 1; i >= 0; i-- {
		// Go down if value is bigger or ptr is the last one element
		var ptr *Node
		for ptr = sl.Head; ptr.Next[i] != nil; ptr = ptr.Next[i] {
			if ptr.Next[i].Value > value {
				break
			}
		}
		// Insert
		if i <= level {
			newNode.Next[i] = ptr.Next[i]
			ptr.Next[i] = newNode
		}
	}
	sl.Length++
}

func (sl *Skiplist) Delete(value int) {
	if !sl.Contains(value) {
		var prePtr *Node
		for i := sl.Level - 1; i >= 0; i-- {
			for ptr := sl.Head; ptr.Next[i] != nil; ptr = ptr.Next[i] {
				if ptr.Value == value {
					prePtr.Next[i] = ptr.Next[i].Next[i]
					break
				}
				if ptr.Next[i].Value > value {
					break
				}
			}
		}
	}
}

func (sl *Skiplist) Print() {
	for i := sl.Level - 1; i >= 0; i-- {
		fmt.Printf("Level %d: ", i)
		for ptr := sl.Head; ptr.Next[i] != nil; ptr = ptr.Next[i] {
			fmt.Printf("%v ", ptr.Next[i].Value)
		}
		fmt.Printf("\n\n")
	}
}

func (sl *Skiplist) Contains(value int) bool {
	var i int
	for i = sl.Level - 1; i >= 0; i-- {
		for ptr := sl.Head; ptr.Next[i] != nil; ptr = ptr.Next[i] {
			if ptr.Next[i].Value == value {
				return true
			}
			if ptr.Next[i].Value > value {
				break
			}
		}
	}
	return false
}

func (sl *Skiplist) Len() int {
	return sl.Length
}
