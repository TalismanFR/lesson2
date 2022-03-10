package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, finish := context.WithCancel(context.Background())

	result := make(chan int)

	for i := 0; i <= 10; i++ {
		go worker(ctx, i, result)
	}

	foundBy := <-result
	fmt.Println("Result found by ", foundBy)
	finish()

	time.Sleep(time.Second)
}

func worker(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond

	fmt.Println(workerNum, "slepp ", waitTime)
	select {
	case <-time.After(waitTime):
		fmt.Println("worker ", workerNum, "done")
		out <- workerNum
	case <-ctx.Done():
		fmt.Println("Signal finish recived worker ", workerNum)
		return
	}
}
