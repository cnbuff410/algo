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

func populate(n int) *Skiplist {
	l := New(n)
	for i := 0; i < n; i++ {
		l.Insert(i)
	}
	return l
}

var l100 = populate(100)

func BenchmarkLookup100(b *testing.B) {
	l := l100
	for i := 0; i < b.N; i++ {
		l.Contains(i)
	}
}

var l1000 = populate(1000)

func BenchmarkLookup1000(b *testing.B) {
	l := l1000
	for i := 0; i < b.N; i++ {
		l.Contains(i)
	}
}

var l10000 = populate(10000)

func BenchmarkLookup10000(b *testing.B) {
	l := l10000
	for i := 0; i < b.N; i++ {
		l.Contains(i)
	}
}

var l100000 = populate(100000)

func BenchmarkLookup100000(b *testing.B) {
	l := l100000
	for i := 0; i < b.N; i++ {
		l.Contains(i)
	}
}
