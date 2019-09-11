package conn

import (
	"fmt"
	"github.com/1071496910/go-redis-client/util"
	"testing"
)

var (
	conn *Conn
)

func init() {
	var err error
	conn, err = NewConn("127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
}

func BenchmarkSet(b *testing.B) {

	for i := 0; i < b.N; i++ {
		req := fmt.Sprintf("SET %v %v", i, i)
		conn.Send(util.Bytes(req))
		conn.Recv()
		//fmt.Println("Get result:", util.String(conn.lastResponse))
	}
}
