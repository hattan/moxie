package main

import (
	"os"
	"testing"
)

var args = make([]string, 1000)

func setup() {

	for i := 0; i < 1000; i++ {
		args[i] = "a"
	}
}
func TestMain(m *testing.M) {
	for i := 0; i < 1000; i++ {
		args[i] = "a"
	}

	code := m.Run()

	os.Exit(code)
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getStringJoin(args)
	}
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getStringLoop(args)
	}
}

func BenchmarkLoop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getStringLoop2(args)
	}
}
