package main

import "testing"

func BenchmarkStuff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		executeOrder66()
	}
}
