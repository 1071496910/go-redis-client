package main

import (
	"bufio"
	"bytes"
	"fmt"
	"time"
)

func main() {
	buf := bytes.NewBuffer([]byte("hello world\r\nasdfa\r\nadfasdf\r\n"))
	r := bufio.NewReader(buf)
	for {
		time.Sleep(time.Second)
		line, _, err := r.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Print(string(line))
		fmt.Println("a")
	}

}
