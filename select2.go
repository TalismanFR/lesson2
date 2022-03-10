package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	cancalCh := make(chan struct{})
	dataCh := make(chan int)

	go scanWait(cancalCh)

	go func(cancelCh <-chan struct{}, dataCh chan int) {
		val := 0
		for {
			select {
			case dataCh <- val:
				val++
				time.Sleep(time.Second)
			case <-cancelCh:
				close(dataCh)
				return
			}
		}
	}(cancalCh, dataCh)

	for val := range dataCh {
		fmt.Println("Read ", val)
	}
}

func scanWait(cancelCh chan<- struct{}) {
	os.Stdin.Read(make([]byte, 1))
	cancelCh <- struct{}{}
}
