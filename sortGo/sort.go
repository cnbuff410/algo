package main

import (
	"fmt"
	"sort"
)

var (
	aux itemSlice = make(itemSlice, 5000000) // Used for merge sort
)

type itemSlice []int

func (il itemSlice) Len() int {
	return len(il)
}

func (il itemSlice) Less(i, j int) bool {
	return il[i] < il[j]
}

func (il itemSlice) Lesse(i, j int) bool {
	return il[i] <= il[j]
}

func (il itemSlice) Swap(i, j int) {
	il[i], il[j] = il[j], il[i]
}

func min(i1, i2 int) int {
	switch {
	case i1 < i2:
		return i1
	case i1 > i2:
		return i2
	}
	return -1
}

func (il itemSlice) Show() {
	for _, i := range il {
		fmt.Println(i)
	}
}

func (il itemSlice) IsSorted() bool {
	Size := len(il)
	if Size == 1 {
		return true
	} // Single item is sorted
	for i := 1; i < Size; i++ {
		if il.Less(i, i-1) {
			return false
		}
	}
	return true
}

func (il itemSlice) NativeSort() {
	sort.Sort(il)
}

func (il itemSlice) InsertionSort() {
	size := len(il)
	for i := 1; i < size; i++ {
		for j := i; j > 0 && il.Less(j, j-1); j-- {
			il.Swap(j, j-1)
		}
	}
}

func (il itemSlice) ShellSort() {
	size := len(il)
	step := 1
	for step < size/3 {
		step = 3*step + 1 // Find the maximum step
	}
	for step >= 1 { // h sort
		for i := step; i < size; i++ {
			for j := i; (j >= step) && (il.Less(j, j-step)); j -= step {
				il.Swap(j, j-1)
			}
		}
		step /= 3
	}
}

func (il itemSlice) merge(low, mid, high int) {
	l := il[low : high+1]
	if l.IsSorted() {
		return
	}
	copy(aux[low:high+1], l)
	i := low
	j := mid + 1
	for k := low; k <= high; k++ {
		switch {
		case i > mid:
			il[k] = aux[j]
			j++
		case j > high:
			il[k] = aux[i]
			i++
		case aux.Less(j, i):
			il[k] = aux[j]
			j++
		default:
			il[k] = aux[i]
			i++
		}
	}
}

func (il itemSlice) MergeSortTD(low int, high int) { // Top-down approach
	if low >= high {
		return
	}
	mid := low + (high-low)/2
	il.MergeSortTD(low, mid)
	il.MergeSortTD(mid+1, high)
	il.merge(low, mid, high)
}

func (il itemSlice) MergeSortBU(low int, high int) { // Bottom-up approach
	size := len(il)
	for step := 1; step < size; step *= 2 {
		for i := 0; i < size-step; i += step * 2 {
			il.merge(i, i+step-1, min(size-1, i+step*2-1))
		}
	}
}

func (il itemSlice) partition(low, high int) int {
	pivot := low
	i := low
	j := high
	for {
		if i == pivot {
			i++
		}
		if j == pivot {
			j--
		}
		for il.Lesse(i, pivot) {
			i++
			if i == high+1 {
				break
			}
		}
		for il.Lesse(pivot, j) {
			j--
			if j == low {
				break
			}
		}
		if i >= j {
			break
		}
		il.Swap(i, j)
	}
	il.Swap(pivot, j)
	return j
}

func (il itemSlice) QuickSort(low, high int) {
	if low >= high {
		return
	}
	middle := il.partition(low, high)
	il.QuickSort(low, middle-1)
	il.QuickSort(middle+1, high)
}

type heap struct {
	il    itemSlice
	count int
}

func PQ(capacity int) *heap {
	il := make(itemSlice, capacity)
	count := 0
	return &heap{il, count}
}

func (h *heap) Sink(k int) {
	for 2*k < h.il.Len() {
		j := k * 2
		if j < h.il.Len() && h.il.Less(j, j+1) {
			j++
		}
		if !h.il.Less(k, j) {
			break
		}
		h.il.Swap(k, j)
		k = j
	}
}

func (h *heap) Swim(k int) {
	for k > 1 && h.il.Less(k/2, k) {
		h.il.Swap(k/2, k)
		k = k / 2
	}
}

func (h *heap) Insert(v int) {
	if h.count < h.il.Len() {
		h.count++
		h.il[h.count] = v
		h.Swim(h.count)
	}
}

func (h *heap) Max() int {
	return h.il[1]
}

func (h *heap) Delmax() int {
	max := h.il[1]
	h.il.Swap(1, h.count)
	h.il[h.count] = 0
	h.count--
	h.Sink(1)
	return max
}

func (h *heap) IsEmpty() bool {
	return h.count == 0
}

func main() {
	fmt.Println("hello")
}
