package conn

import (
	"bufio"
	"errors"
	"net"
	"time"
)

var (
	ErrEmptyRequest = errors.New("request is empty")

	CRLF = `\r\n`
)

//Conn 实现net.conn接口，实现 命令行到redis protocol 的转换，
//write的 接收参数是命令行的字符串的[]byte,
//read的，返回值是命令行字符串的[]byte
type Conn struct {
	innerBuffer []byte //环形队列保存未解析完的内容
	head        int
	tail        int
	conn        *net.Conn
	reader      *bufio.Reader
}

func parseRequest(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return nil, ErrEmptyRequest
	}
	//strings.Split(util.String(b))
	return nil, nil
}

func (c Conn) Read(b []byte) (n int, err error) {
	panic("implement me")
}

func (c Conn) Write(b []byte) (n int, err error) {
	panic("implement me")
}

func (c Conn) Close() error {
	panic("implement me")
}

func (c Conn) LocalAddr() net.Addr {
	panic("implement me")
}

func (c Conn) RemoteAddr() net.Addr {
	panic("implement me")
}

func (c Conn) SetDeadline(t time.Time) error {
	panic("implement me")
}

func (c Conn) SetReadDeadline(t time.Time) error {
	panic("implement me")
}

func (c Conn) SetWriteDeadline(t time.Time) error {
	panic("implement me")
}
