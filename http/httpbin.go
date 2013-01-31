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
		s, ok := ip["origin"].(string)
		if ok {
			log.Println("Your IP:", ip["origin"])
			return s
		}
	}
	return ""
}

func getUA(u *url.URL) string {
	ua, err := getJsonResp(u, "/user-agent")
	if err == nil {
		s, ok := ua["user-agent"].(string)
		if ok {
			log.Println("Your UA:", s)
			return s
		}
	}
	return ""
}

type RespJsonType map[string]interface{}

func getJsonResp(u *url.URL, p string) (RespJsonType, error) {
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

	log.Println("Response:", string(body))

	jsonBlob := make(RespJsonType)
	err = json.Unmarshal(body, &jsonBlob)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return jsonBlob, nil
}

func fakeUA(ua string) {
	req, err := http.NewRequest("GET", "http://httpbin.org/user-agent", nil)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v", req)

	req.Header.Set("User-Agent", ua)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func getHeader(u *url.URL) interface{} {
	header, err := getJsonResp(u, "/headers")
	if err == nil {
		log.Println("Your Headers:", header["headers"])
		return header["headers"]
	}
	return nil
}

func getGet(u *url.URL) interface{} {
	getData, err := getJsonResp(u, "/get")
	if err == nil {
		log.Println("Your get data:", getData)
		return getData
	}
	return nil
}

func main() {
	flag.Parse()

	log.SetFlags(log.Lmicroseconds)

	if *h {
		help(u)
	}

	getIP(u)

	getUA(u)

	fakeUA("Golang httpbin")
	getUA(u)

	getHeader(u)

	getGet(u)

}
