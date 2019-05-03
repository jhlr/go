package table

import (
	"math/rand"
	"testing"
)

func benchmark(b *testing.B, t Interface) {
	var p []int
	add := func(b *testing.B) {
		for k, e := range p {
			t.Add(k).Set(e)
		}
	}
	node := func(b *testing.B) {
		for _, i := range p {
			if t.Node(i) == nil {
				b.Fatal("not found")
			}
		}
	}
	remove := func(b *testing.B) {
		for _, i := range p {
			if !t.Remove(i) {
				b.Fatal("not found", i)
			}
		}
	}
	b.Run("Add", func(b *testing.B) {
		p = rand.Perm(b.N)
		b.ResetTimer()
		add(b)
	})
	b.Run("Node", func(b *testing.B) {
		p = rand.Perm(b.N)
		add(b)
		b.ResetTimer()
		node(b)
	})
	b.Run("Remove", func(b *testing.B) {
		p = rand.Perm(b.N)
		add(b)
		b.ResetTimer()
		remove(b)
	})
}

func BenchmarkTree(b *testing.B) {
	f := CompareFunc(func(ka, kb interface{}) int {
		return ka.(int) - kb.(int)
	})
	benchmark(b, NewTree(f))
}

func BenchmarkHash(b *testing.B) {
	f := HashFunc(func(k interface{}) uint64 {
		return uint64(k.(int))
	})
	benchmark(b, NewHash(47, f))
}
