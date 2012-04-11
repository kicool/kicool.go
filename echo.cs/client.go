package main

// Connect to JSONRPC Server and send command-line args to Echo

import (
	"flag"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
)

var destAddr = flag.String("d", "127.0.0.1:1234", "server addr")
var message = flag.String("m", "", "message send to server")

func main() {
	flag.Parse()
	log.SetPrefix("echo.client: ")

	conn, e := net.Dial("tcp", *destAddr)
	if e != nil {
		log.Fatalf("Could not connect: %s\n", e)
		os.Exit(1)
	}

	client := jsonrpc.NewClient(conn)

	var reply string

	log.Println("Sending:", *message)
	client.Call("RPCFunc.Echo", *message, &reply)
	log.Println("Reply :", reply)
}
