package main
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_Sort(t *testing.T) {
	/* Test priority queue first */
	pq := PQ(10) // Priority queue with size 10
	pq.Insert(16)
	pq.Insert(17)
	pq.Insert(5)
	pq.Delmax()
	pq.Insert(25)
	pq.Insert(1)
	pq.Insert(13)
	pq.Delmax()
	pq.Insert(16)
	pq.Insert(12)
	pq.Insert(5)
	pq.Delmax()
	result := []int{0, 16, 13, 12, 1, 5, 5, 0, 0, 0}
	for i := 0; i < pq.il.Len(); i++ {
		if result[i] != pq.il[i] {
			t.Errorf("Priority queue is wrong")
			break
		}
	}

	filename := "./small"
	f, ferr := os.Open(filename)
	if ferr != nil {
		panic(ferr)
	}
	list := make(itemSlice, 0)
	br := bufio.NewReader(f)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			break
		}
		strList := strings.Fields(strings.TrimSpace(string(line)))
		for _, s := range strList {
			i, _ := strconv.Atoi(s)
			list = append(list, i)
		}
	}
	totalSize := len(list)

	// Builtin sort
	localList1 := make(itemSlice, totalSize)
	copy(localList1, list)
	start := time.Nanoseconds()
	localList1.NativeSort()
	end := time.Nanoseconds()
	println("Ellapsed time for built in sort: ", (end-start)/1000000.0, " ms")
	if localList1.IsSorted() != true {
		t.Errorf("Sorting is wrong")
	}

	// Insertion sort
        localList2 := make(itemSlice, totalSize)
        copy(localList2, list)
        start = time.Nanoseconds()
        localList2.InsertionSort()
        end = time.Nanoseconds()
        t.Log("Ellapsed time for insertion sort: ", (end-start)/1000000.0, " ms")
        if localList2.IsSorted() != true {
        t.Errorf("Sorting is wrong")
        }

	// Shell sort
        localList3 := make(itemSlice, totalSize)
        copy(localList3, list)
        start = time.Nanoseconds()
        localList3.ShellSort()
        end = time.Nanoseconds()
        println("Ellapsed time for shell sort: ", (end-start)/1000000.0, " ms")
        if localList3.IsSorted() != true {
        t.Errorf("Sorting is wrong")
        }

	// Merge sort top-down
	localList4 := make(itemSlice, totalSize)
	copy(localList4, list)
	start = time.Nanoseconds()
	localList4.MergeSortTD(0, len(localList4)-1)
	end = time.Nanoseconds()
	println("Ellapsed time for merge sort top-down: ", (end-start)/1000000.0, " ms")
	if localList4.IsSorted() != true {
		t.Errorf("Sorting is wrong")
	}

	// Merge sort bottom-up
	localList5 := make(itemSlice, totalSize)
	copy(localList5, list)
	start = time.Nanoseconds()
	localList5.MergeSortBU(0, len(localList5)-1)
	end = time.Nanoseconds()
	println("Ellapsed time for merge sort bottom-up: ", (end-start)/1000000.0, " ms")
	if localList5.IsSorted() != true {
		t.Errorf("Sorting is wrong")
	}

	// Quick sort
	localList6 := make(itemSlice, totalSize)
	copy(localList6, list)
	start = time.Nanoseconds()
	localList6.QuickSort(0, len(localList6)-1)
	end = time.Nanoseconds()
	println("Ellapsed time for quick sort: ", (end-start)/1000000.0, " ms")
	if localList6.IsSorted() != true {
		t.Errorf("Sorting is wrong")
	}
}
