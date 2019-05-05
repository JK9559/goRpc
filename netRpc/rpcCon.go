package netRpc

import (
	"fmt"
	"github.com/JK9559/goRpcApi/api"
	"net/rpc"
)

func StartCon() {
	client, _ := rpc.DialHTTP("tcp", "39.97.162.199:1234")
	args := api.Args{A: 10, B: 8}
	var res int
	client.Call("Alice.Multiply", args, &res)
	fmt.Println(args.A, args.B, res)
}
