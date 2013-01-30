// request yahoo search
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url" //Package url parses URLs and implements query escaping. See RFC 3986.
	"path/filepath"
	"flag"
)

var (
	keyword = flag.String("s", "yahoo", "keyword to search with yahoo")
)

func main() {
	flag.Parse()

	u, err := url.ParseRequestURI("http://search.cn.yahoo.com/s?q=")
	if err != nil {
		log.Fatalln("url.Parse", err)
	}

	q := u.Query()
	q.Set("q", *keyword)
	u.RawQuery = q.Encode()

	fmt.Println(u)

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln("http.Get", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ioutil.ReadAll", err)
	}

	//log.Println(string(body))
	path, _ := filepath.Abs("./resp.html")
	err = ioutil.WriteFile(path, body, 0666)
	if err != nil {
		log.Fatalln("ioutil.WriteFile", err)
	}
}
