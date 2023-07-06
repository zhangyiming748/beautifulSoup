package main

import (
	"fmt"
	"github.com/zhangyiming748/BeautifulSoup"
)

func main() {
	BeautifulSoup.SetDebug(true)
	BeautifulSoup.Header("Accept-Language", "en")
	html := "https://telegra.ph"
	proxy := "http://127.0.0.1:8889"
	withProxy, err := BeautifulSoup.GetWithProxy(html, proxy)
	if err != nil {
		fmt.Println("Get 失败")

	} else {
		fmt.Println("Get 成功")
		fmt.Println(withProxy)
	}
}
