package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*3)

	// go func() {
	// 	time.Sleep(time.Millisecond * 200)
	// 	cancel()
	// }()

	parse(ctx)
}

func parse(ctx context.Context){
	for{
		select{
		case <-time.After(time.Second*2):
			fmt.Println("parse completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return
		}
	}
}