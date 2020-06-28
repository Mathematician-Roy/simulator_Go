package main
import "testing"

func BenchmarkStuff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		simulator_mainGame_goroutine(1000000, 10)
	}
}

