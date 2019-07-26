package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func BenchmarkClosure(b *testing.B) {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		f()
	}
}

func BenchmarkIterator(b *testing.B) {
	for i := 0; i < 10; i++ {
		fib(i)
	}
}
