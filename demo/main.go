package main

import (
    "fmt"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    // 是否阻塞的开关
    var flag bool
    testCh := make(chan int)

    go func() {
        // pprof 监听的端口 8080
        if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
            fmt.Printf("%v", err)
        }
    }()

    go func() {
        // 永远为 false 的条件
        if flag {
            <-testCh
        }

        // 死循环
        select {}
    }()

    // 每秒执行100次
    tick := time.Tick(time.Second / 100)
    for range tick {
        ch1(testCh)
    }
}

func ch1(ch chan<- int) {
    go func() {
        defer fmt.Println("ch1 stop")
        // 给 ch channel 写入数据
        ch <- 0
    }()
}