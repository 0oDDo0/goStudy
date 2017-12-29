package main

import (
	"context"
	"log"
	"time"
)
func C(ctx context.Context) {
	times := 1
	for {
		select {
		case <-ctx.Done():
			log.Println("C Done")
			return
		case <-time.After(1 * time.Second):
			log.Println("C times=", times)
			times++
		}
	}
}

func B(ctx context.Context) {
	times := 1
	ctx, _ = context.WithCancel(ctx)
	go C(ctx)
	for {
		select {
		case <-ctx.Done():
			log.Println("B Done")
			return
		case <-time.After(1 * time.Second):
			log.Println("B times=", times)
			times++
		}
	}
}

func A(ctx context.Context) {
	times := 1
	go B(ctx)
	for {
		select {
		case <-ctx.Done():
			log.Println("A Done")
			return
		case <-time.After(1 * time.Second):
			log.Println("A times=", times)
			times++
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(10 * time.Second)
		cancel()
	}()
	A(ctx)
	time.Sleep(5 * time.Second)
}
