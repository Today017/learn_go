package main

import (
	"fmt"
	"net/url"
)

func test() {
	u, _ := url.Parse("http://example.com/article/123?name=golang&age=10")
	queryMap := u.Query()

	fmt.Println(queryMap)
	fmt.Println(queryMap.Get("name"))
	fmt.Println(queryMap.Get("age"))
}
