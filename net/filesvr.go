package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	webroot = flag.String("root", "./", "web root directory")
	addr    = flag.String("addr", ":1981", "ipv4:port to listen")
)

func main() {
	flag.Parse()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(*webroot))))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Panicln("ListenAndServe:", err)
	}
}
