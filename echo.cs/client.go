package main

// Connect to JSONRPC Server and send command-line args to Echo

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"flag"
)

var destAddr = flag.String("d", "127.0.0.1:1234", "server addr")
var message = flag.String("m", "", "message send to server")

func main() {
	flag.Parse()

	conn, e := net.Dial("tcp", *destAddr)
	if e != nil {
		fmt.Fprintf(os.Stderr, "Could not connect: %s\n", e)
		os.Exit(1)
	}
	client := jsonrpc.NewClient(conn)
	var reply string
	fmt.Printf("Sending: %s\n", *message)
	client.Call("RPCFunc.Echo", *message, &reply)
	fmt.Printf("Reply:  %s\n", reply)
}
