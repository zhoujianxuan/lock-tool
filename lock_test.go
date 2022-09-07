package lock_tool

import (
	"math/rand"
	"testing"
)

func BenchmarkLockTool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TestWork()
		}
	})
}

func BenchmarkLockToolLock(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TestWorkLock()
		}
	})
}

func BenchmarkLockToolPrivate(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			key := rand.Int()
			TestWorkPrivate(key%100)
		}
	})
}