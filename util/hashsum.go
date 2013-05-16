package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"os"
)

var (
	file = flag.String("f", "", "file to hash")
	algo = flag.String("h", "md5", "which hash algo to use")
)

func main() {
	flag.Parse()

	//open file
	b, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var h hash.Hash
	switch *algo {
	case "sha1":
		h = sha1.New()
	case "sha256":
		h = sha256.New()
	case "sha224":
		h = sha256.New224()
	case "sha512":
		h = sha512.New()
	case "sha384":
		h = sha512.New384()
	case "md5":
		h = md5.New()
	default:
		log.Println("unsported hash algorithm", *algo)
		os.Exit(3)
	}

	_, err = h.Write(b)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	s := h.Sum(nil)
	fmt.Printf("%x  %s\n", s, *file)

	os.Exit(0)
}
