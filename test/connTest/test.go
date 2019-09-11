package main

import (
	"fmt"
	"github.com/1071496910/go-redis-client/conn"
	"github.com/1071496910/go-redis-client/util"
)

func main() {
	c, err := conn.NewConn("127.0.0.1:6379")
	if err != nil {
		panic(err)
	}

	//_ ,err = c.Send(util.Bytes("LRANGE l 0 4"))
	_, err = c.Send(util.Bytes("GET b"))
	if err != nil {
		panic(err)
	}
	fmt.Println("after write")

	_, err = c.Recv()
	if err != nil {
		panic(err)
	}
}
