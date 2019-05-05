package jsonRpc

import (
	"errors"
	"fmt"
	"github.com/JK9559/goRpcApi/api"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Arith struct {
}

func (ar *Arith) Multiply(req api.ArithReq, rsp *api.ArithRsp) error {
	rsp.Pro = req.A * req.B
	return nil
}

func (ar *Arith) Divide(req api.ArithReq, rsp *api.ArithRsp) error {
	if 0 == req.B {
		return errors.New("divide by zero")
	}
	rsp.Quo = req.A / req.B
	rsp.Rem = req.A % req.B
	return nil
}

// https://studygolang.com/articles/14336

func JsonRpcSer() {
	_ = rpc.Register(new(Arith))

	lis, err := net.Listen("tcp", "127.0.0.1:8082")
	if nil != err {
		log.Fatalln("fatal error: ", err)
	}

	_, _ = fmt.Fprintf(os.Stdout, "%s\n", "start connection")

	for {
		conn, err := lis.Accept()
		if nil != err {
			continue
		}

		go func(conn net.Conn) {
			_, _ = fmt.Fprintf(os.Stdout, "%s\n", "new client in coming")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

}
