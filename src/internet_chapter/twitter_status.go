package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type status struct {
	Text string
}

type user struct {
	XMLName xml.Name
	Status  status
}

func main() {
	response, _ := http.Get("https://twitter.com/ymkfnuiwgij")
	user := user{xml.Name{"", "user"}, status{""}}
	data, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(data, &user)
	fmt.Printf("status: %s", user.Status.Text)
}
