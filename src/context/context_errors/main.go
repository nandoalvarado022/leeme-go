package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func myfunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Ctx is kicking in with error:%+v\n", ctx.Err())
			return
		default:
			time.Sleep(15 * time.Second)
			fmt.Printf("I was not canceled\n")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3*time.Second))
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		myfunc(ctx)
	}()

	wg.Wait()
	fmt.Printf("In main, ctx err is %+v\n", ctx.Err())
}
