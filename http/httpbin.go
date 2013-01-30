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

func getIP(u *url.URL) string {
	ip, err := getJsonResp(u, "/ip")
	if err == nil {
		log.Println("Your IP:", ip["origin"])
		return ip["origin"]
	}
	return ""
}

func getUA(u *url.URL) string {
	ua, err := getJsonResp(u, "/user-agent")
	if err == nil {
		log.Println("Your UA:", ua["user-agent"])
		return ua["user-agent"]
	}
	return ""

}

func getJsonResp(u *url.URL, p string) (map[string]string, error) {
	u.Path = p

	log.Println("getJsonResp from ", u)

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatalln("http.Get", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("ioutil.ReadAll", err)
		return nil, err
	}

	//log.Println(string(body))

	jsonBlob := make(map[string]string)
	err = json.Unmarshal(body, &jsonBlob)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Response", jsonBlob)
	return jsonBlob, nil
}

func main() {
	flag.Parse()

	log.SetFlags(log.Lmicroseconds)

	if *h {
		help(u)
	}

	getIP(u)

	getUA(u)

}
