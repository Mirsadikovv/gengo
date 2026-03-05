package main

import "testing"

//BenchmarkAtomicMapStore-32    10000	    128909 ns/op	  181755 B/op	   17 allocs/op
//BenchmarkAtomicMapLoad-32    	720423896	1.651 ns/op	       0 B/op	       0 allocs/op
//BenchmarkMutexMapStore-32    	14464365	134.2 ns/op	      97 B/op	       0 allocs/op
//BenchmarkMutexMapLoad-32    	100000000	10.66 ns/op	       0 B/op	       0 allocs/op

func BenchmarkAtomicMapStore(b *testing.B) {
	b.ResetTimer()
	m := NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}

}

func BenchmarkAtomicMapLoad(b *testing.B) {
	m := NewSyncMap[int, int]()
	m.Store(1, 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Load(1)
	}
}
