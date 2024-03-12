// Todo oom
package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"
)

func RandStr(n int) string {
	var letters = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make_int32(n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return int32_to_str(b)
}

func make_int32(n int) []int32 {
	return make([]int32, n)
}

func int32_to_str(b []int32) string {
	return string(b)
}

func Benchmark_RandStr_go(b *testing.B) {
	f, _ := os.Create("./tmp/mem1")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go RandStr(100000)
	}
}

// use pool
func Benchmark_RandStr(b *testing.B) {
	f, _ := os.Create("./tmp/mem2")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			RandStr(100000)
		}
	})
}

// go + ch block -> oom
var ch3 = make(chan string, 8)

func channel_block() {
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ch3 <- letters
}

func Benchmark_channel_block(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go channel_block()
	}
}
