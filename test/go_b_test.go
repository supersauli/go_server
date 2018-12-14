package test

import "testing"

func Add(max int) int {
	var val int
	for i := 0; i < max; i++ {
		val += i
	}
	return val
}

//Benchmark这个是必须的
func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i)
	}
}

func Benchmark_TimeAdd(b *testing.B) {
	//b.N = 100000
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Add(i)
	}
}
