package conn

import (
	"bufio"
	"errors"
	"github.com/1071496910/go-redis-client/util"
	"net"
	"strconv"
	"strings"
)

var (
	ErrEmptyRequest    = errors.New("request is empty")
	ErrInvalidResponse = errors.New("invalid response")

	CRLF = "\r\n"
)

//Conn 实现net.conn接口，实现 命令行到redis protocol 的转换，
//write的 接收参数是命令行的字符串的[]byte,
//read的，返回值是命令行字符串的[]byte
type Conn struct {
	net.Conn
	reader       *bufio.Reader
	lastResponse []byte
}

func NewConn(addr string) (*Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(conn)

	return &Conn{
		Conn:   conn,
		reader: reader,
	}, nil
}

func parseRequest(b []byte) ([]byte, error) {
	if len(b) == 0 {
		return nil, ErrEmptyRequest
	}
	result := ""
	requestTokens := strings.Split(util.String(b), " ")
	l := len(requestTokens)
	result = result + "*" + strconv.Itoa(l) + CRLF
	for i := 0; i < l; i++ {
		tokenLen := len(requestTokens[i])
		result = result + "$" + strconv.Itoa(tokenLen) + CRLF
		result = result + requestTokens[i] + CRLF
	}
	return util.Bytes(result), nil
}

func parseOneLineResponse(firstByte byte, r *bufio.Reader) ([]byte, error) {

	result := []byte{}
	switch firstByte {
	case '+':
		fallthrough
	case ':':
		content, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}
		result = append(result, content...)
		return result, nil
	case '-':
		content, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}
		result = append(result, content...)
		return result, errors.New(util.String(result))
	case '$':
		lstr, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}
		l, err := strconv.Atoi(util.String(lstr))
		if err != nil {
			return nil, err
		}
		content, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}
		if len(content) != l {
			return nil, ErrInvalidResponse
		}
		return content, nil
	}
	return nil, nil
}

func parseResponse(r *bufio.Reader) ([]byte, error) {
	firstByte, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	result := []byte{}
	switch firstByte {
	case '+':
		fallthrough
	case ':':
		fallthrough
	case '-':
		fallthrough
	case '$':
		return parseOneLineResponse(firstByte, r)
	case '*':
		lstr, _, err := r.ReadLine()
		if err != nil {
			return nil, err
		}
		l, err := strconv.Atoi(util.String(lstr))
		if err != nil {
			return nil, err
		}
		for i := 0; i < l; i++ {
			firstByte, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
			tmpResult, err := parseOneLineResponse(firstByte, r)
			if err != nil {
				return nil, err
			}
			result = append(result, ' ')
			result = append(result, tmpResult...)
		}
		return result, nil
	}
	return nil, nil
}

func (c *Conn) Recv() ([]byte, error) {
	result, err := parseResponse(c.reader)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Conn) Send(b []byte) (n int, err error) {
	request, err := parseRequest(b)
	if err != nil {
		return -1, err
	}
	n, err = c.Conn.Write(request)
	return
}
