package main

import (
	"fmt"
	"goRpc/jsonRpc"
)

func main() {
	fmt.Println("calling Service")
	jsonRpc.JsonRpcSer()
}
