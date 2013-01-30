// A http sample using httpbin.org
package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var (
	u = &url.URL{Host: "httpbin.org", Scheme: "http"}
	h = flag.Bool("h", false, "show help for httpbin.org")
)

func help(u *url.URL) {
	log.Println("help for ", u)

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln("http.Get", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ioutil.ReadAll", err)
	}

	log.Println(string(body))
}

func getIP(u *url.URL) {
	log.Println("get origin IP")

	u.Path = "/ip"

	log.Println("get from ", u)

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

	type IPJson struct {
		Origin string
	}
	var ipjson IPJson
	err = json.Unmarshal(body, &ipjson)
	if err != nil {
		log.Println(err)
	}

	log.Println("Your ip is", ipjson.Origin)
}

func main() {
	flag.Parse()

	if *h {
		help(u)
	}

	getIP(u)

}
