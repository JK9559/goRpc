package jsonRpc

import (
	"fmt"
	"github.com/JK9559/goRpcApi/api"
	"log"
	"net/rpc/jsonrpc"
)

func JsonRpcCli() {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8082")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := api.ArithReq{A: 9, B: 2}
	var res api.ArithRsp

	err = conn.Call("Arith.Multiply", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, res.Pro)

	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("%d / %d = %d ...... %d\n", req.A, req.B, res.Quo, res.Rem)
}
