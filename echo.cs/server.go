package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RPCFunc uint8

func (*RPCFunc) Echo(arg *string, result *string) error {
	log.Print("Arg passed: " + *arg)
	*result = ">" + *arg + "<"
	return nil
}

func main() {
	log.Print("Starting Server...")
	l, err := net.Listen("tcp", "localhost:1234")
	defer l.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("listening on: ", l.Addr())
	rpc.Register(new(RPCFunc))
	for {
		log.Print("waiting for connections ...")
		conn, err := l.Accept()
		if err != nil {
			log.Printf("accept error: %s", conn)
			continue
		}
		log.Printf("connection started: %v", conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}
}
