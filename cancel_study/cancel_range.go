package main

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func test_put(cli clientv3.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		// handle error!
	}
	// use the response
	fmt.Println(resp)
}

func test_get(cli clientv3.Client) {
	fmt.Println("start")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	// ctx, cancel := context.WithCancel(context.Background())
	// resp, err := cli.Get(ctx, "a", clientv3.WithPrefix())
	_, err := cli.Get(ctx, "1", clientv3.WithPrefix())
	defer cancel()
	if err != nil {
		// handle error!
		fmt.Println("err")
	}
	// use the response
	// fmt.Println(resp)
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Println("err")
	}
	defer cli.Close()

	// test_put(*cli)

	for i := 0; i < 100; i++ {
		go test_get(*cli)
		time.Sleep(time.Second)
	}

	// // Wait for all goroutines to finish
	time.Sleep(200 * time.Second)
	fmt.Println("All goroutines have finished.")
}
