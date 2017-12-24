package main

import (
	"fmt"
	"net/http"
)

func main() {
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.ishuhui.com/",
		"https://www.bilibili.com/",
	}

	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
