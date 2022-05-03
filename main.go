package main

import (
	"my_grpc/grpc_practice/client"
	"my_grpc/grpc_practice/server"
	"time"
)

func main() {
	go server.Start()
	time.Sleep(1 * time.Second)
	client.Start()

	var b chan bool
	<-b
}
