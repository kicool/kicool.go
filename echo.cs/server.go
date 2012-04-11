package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var addr = flag.String("i", "127.0.0.1:1234", "listen address and port")

type RPCFunc uint8

func (*RPCFunc) Echo(arg *string, result *string) error {
	log.Print("Arg passed: " + *arg)
	*result = ">" + *arg + "<"
	return nil
}

func main() {
	flag.Parse()
	log.SetPrefix("echo.server: ")

	log.Print("Starting Server...")
	l, err := net.Listen("tcp", *addr)
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
			log.Fatal("accept error:", conn, err)
			continue
		}
		log.Print("connection started:", conn.LocalAddr(), conn.RemoteAddr())
		go jsonrpc.ServeConn(conn)
	}
}
