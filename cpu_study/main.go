package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func test_put(cli clientv3.Client, key string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		// handle error!
	}
	// use the response
	// fmt.Println(resp)
}

func test_get(cli clientv3.Client, key string, TTL int) {
	fmt.Println("start")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(TTL)*time.Millisecond)
	// ctx, cancel := context.WithCancel(context.Background())
	// resp, err := cli.Get(ctx, "a", clientv3.WithPrefix())
	_, err := cli.Get(ctx, key, clientv3.WithPrefix())
	defer cancel()
	file, _ := os.Create("output.txt")
	defer file.Close()
	if err != nil {
		// handle error!
		file.WriteString("err")
	}
	// use the response
	file.WriteString("resp")
}

func test_watch(cli clientv3.Client, key string, TTL int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(TTL)*time.Second)
	rch := cli.Watch(ctx, key)
	file, _ := os.Create("output.txt")
	defer file.Close()
	for wresp := range rch {
		for _, ev := range wresp.Events {
			line := fmt.Sprintf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			file.WriteString(line)
		}
	}
	defer cancel()
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

	// keys := [4]string{"Penn", "Teller", "aba", "ababa"}
	// num, _ := strconv.Atoi(os.Args[1])
	// TTL, _ := strconv.Atoi(os.Args[2])
	// for i := 0; i < num; i++ {
	// 	go test_watch(*cli, "key", TTL)
	// }

	// start := time.Now()
	// for time.Since(start) < time.Duration(2*TTL)*time.Second {
	// 	test_put(*cli, "key", time.Now().String())
	// 	time.Sleep(100 * time.Microsecond)
	// }

	ttl, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < 10; i++ {
		go test_get(*cli, strconv.Itoa(i), ttl)
		time.Sleep(time.Second)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("All goroutines have finished.")
}
