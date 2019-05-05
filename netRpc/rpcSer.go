package netRpc

import (
	"github.com/JK9559/goRpcApi/api"
	"net"
	"net/http"
	"net/rpc"
)

type Alice int

func (al *Alice) Multiply(args *api.Args, res *int) error {
	*res = args.A * args.B
	return nil
}

func StartSer() {
	al := new(Alice)
	rpc.Register(al)
	rpc.HandleHTTP()
	l, _ := net.Listen("tcp", ":1234")
	http.Serve(l, nil)
}
