func Fibo(num int) int {
	if num < 2 {
		return 1
	}
	return Fibo(num - 1) + Fibo(num - 2)
}

//cpu 所有核占用100%
func Benchmark_all_cpu(b *testing.B) {
	f, _: = os.Create("./tmp/all_cpu_100")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	b.ResetTimer()
	for i:=0; i < b.N; i++ {
		go Fibo(30)
	}
}