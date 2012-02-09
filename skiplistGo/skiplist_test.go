package main

import (
    "testing"
)

func Test_skiplist(t *testing.T) {
    sl := New(32)
    sl.Insert(3)
    sl.Insert(16)
    sl.Insert(17)
    sl.Insert(5)
    sl.Insert(25)
    sl.Insert(1)
    sl.Insert(13)
    sl.Insert(16)
    sl.Insert(12)
    sl.Insert(5)
    sl.Print()
}

func populate(n int) *L {
    l := New()
    for i := 0; i < n; i++ {
        l.Put(strconv.Itoa(i),i)
    }
    return l
}

func BenchmarkInsert(b *testing.B) {
    populate(b.N)
}

var l100000 = populate(1000000)
func BenchmarkLookup100000(b *testing.B) {
        l := l100000
        for i := 0; i < b.N; i++ {
                l.Get(strconv.Itoa(i%10000))
        }
}

var l10m = populateM(10)
func BenchmarkLookup10M(b *testing.B) {
        l := l10m
        for i := 0; i < b.N; i++ {
                l.Get(strconv.Itoa(i%10))
        }
}
