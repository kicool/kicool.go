package main

import (
	"net/url"
	"fmt"
	"github.com/kicool/Normalize-URL"
	"flag"
)

var (
	arg = flag.String("url", "", "url to mormalize")
)

func main() {
	flag.Parse()

	u, err:=url.Parse(*arg)
	if err == nil {
		normalize.Normalize(u)
		fmt.Println(*arg)
		fmt.Println(u)
	}
}


