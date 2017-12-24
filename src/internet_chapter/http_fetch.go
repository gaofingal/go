package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, er := http.Get("http://www.bilibili.com")
	if er != nil {
		log.Fatalf("Get: %v", er)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Get: %v", er)
	}

	fmt.Printf("Got: %q", string(data))
}
