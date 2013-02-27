package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	file = flag.String("f", "", "file to hash")
)

func main() {
	flag.Parse()

	//open file
	b, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	h := md5.New()
	_, err = h.Write(b)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	s := h.Sum(nil)
	fmt.Printf("%x  %s\n", s, *file)

	os.Exit(0)
}
