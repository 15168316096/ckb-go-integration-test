package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime/pprof"
	"strconv"
	"testing"
)

// Todo cpu
func Fibo(num int) int {
	if num < 2 {
		return 1
	}
	return Fibo(num-1) + Fibo(num-2)
}

func Fibo_no_recur(n int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		b = a + b
		a = b - a
	}
	return
}

// cpu 所有核占用100%
func Benchmark_all_cpu(b *testing.B) {
	f, err := os.Create("./tmp/all_cpu_100")
	if err != nil {
		b.Fatalf("Error creating file: %v", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go Fibo(30)
	}
}

func Benchmark_all_cpu_2(b *testing.B) {
	f, err := os.Create("./tmp/all_cpu_2")
	if err != nil {
		b.Fatalf("Error creating file: %v", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go Fibo_no_recur(30)
	}
}

func Benchmark_all_cpu_3(b *testing.B) {
	f, err := os.Create("./tmp/all_cpu_3")
	if err != nil {
		b.Fatalf("Error creating file: %v", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibo_no_recur(30)
	}
}

func Benchmark_all_cpu_4(b *testing.B) {
	f, err := os.Create("./tmp/all_cpu_100")
	if err != nil {
		b.Fatalf("Error creating file: %v", err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Fibo_no_recur(30000000)
	}
}

// Todo io wait
func writeFile(num int) {
	tmp := "测试写文件测试写文件测试写文件测试写文件测试写文件测试写文件\n"
	filename := "./tmp/output" + strconv.Itoa(num) + ".txt"
	for i := 0; i < num; i++ {
		tmp += tmp
	}
	d1 := []byte(tmp)
	err2 := ioutil.WriteFile(filename, d1, 0666)
	if err2 != nil {
		// Use t.Fatalf instead of b.Fatalf
		fmt.Printf("Error creating file: %v", err2)
	}
}

func Benchmark_io_test(b *testing.B) {
	f, _ := os.Create("./tmp/iowait")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeFile(16)
	}
}

func Benchmark_io_test_2(b *testing.B) {
	f, _ := os.Create("./tmp/iowait2")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 10; i++ {
			go Fibo(32)
		}
		writeFile(16)
	}
}
