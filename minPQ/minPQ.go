package main

import "fmt"

type minPQ struct {
	list  []int
	count int
}

func newMinPQ(capacity int) *minPQ {
	list := make([]int, capacity+1) // Start from 1
	count := 0
	return &minPQ{list, count}
}

func (pq *minPQ) Sink(k int) {
	for k * 2 <= pq.count {
		j := k * 2
		if (j < pq.count && pq.list[j] > pq.list[j+1]) { j++ }
		if !(pq.list[k] > pq.list[j]) { break }
		pq.list[k], pq.list[j] = pq.list[j], pq.list[k]
		k = j
	}
}

func (pq *minPQ) Swim(k int) {
	for k > 1 && pq.list[k/2] > pq.list[k] {
		pq.list[k/2], pq.list[k] = pq.list[k], pq.list[k/2] // swap
		k = k / 2
	}
}

func (pq *minPQ) Insert(v int) {
        if pq.count >= len(pq.list) - 1 { pq.resize(2 * pq.count) }
        pq.count++
        pq.list[pq.count] = v
        pq.Swim(pq.count)
}

func (pq *minPQ) Min() int {
	return pq.list[1]
}

func (pq *minPQ) Delmin() int {
        if pq.count ==0 { panic("Priority queue underflow") }
        pq.list[1], pq.list[pq.count] = pq.list[pq.count], pq.list[1] // swap
	min := pq.list[pq.count]
        pq.list = append(pq.list[:pq.count], pq.list[pq.count+1:]...)
	pq.count--
	pq.Sink(1)
        // For optimization, we can resize the list down here if it's too empty
	return min
}

func (pq *minPQ) IsEmpty() bool {
	return pq.count == 0
}

func (pq *minPQ) resize(capacity int) {
        l := make([]int, capacity)
        for i := 1; i <= pq.count; i++ { l[i] = pq.list[i] }
        pq.list= l;
}

func (pq *minPQ) String() string {
        return fmt.Sprintf("There are %v items in queue: %v", pq.count, pq.list)
}

func main() {
        // Test
        pq := newMinPQ(9)
        pq.Insert(1)
        pq.Insert(7)
        pq.Insert(2)
        pq.Insert(5)
        pq.Insert(8)
        pq.Insert(10)
        pq.Insert(3)
        pq.Insert(13)
        pq.Insert(11)
        pq.Insert(17)
        pq.Insert(12)
        pq.Insert(15)
        pq.Insert(18)
        pq.Insert(110)
        pq.Insert(13)
        pq.Insert(113)
        pq.Insert(115)
        pq.Insert(117)
        pq.Delmin()
        fmt.Println(pq.String())
}
