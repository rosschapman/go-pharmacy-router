package main

import (
	"testing"
)

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestMain(t *testing.T) {
	t.Run("test main method", func(t *testing.T) {
		main()
	})
}
