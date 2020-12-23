package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// 主函数入口
func main() {
	fmt.Print("interesting start!\n")
	http.HandleFunc("/getbook", Index)
	fmt.Print(http.ListenAndServe("0.0.0.0:1766", nil))

}
func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var result []byte
	switch randb(2) {
	case 0:
		// 豆瓣
		result = reParseJSON(doubanMain())
	case 1:
		// 浙江新华书店 http://ningbo.zxhsd.com
		result = reParseJSON(nbxhsdMain())
	}
	fmt.Fprint(w, string(result))
}

// struct -> json
func reParseJSON(v interface{}) []byte {
	textbyte, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return textbyte
}

// 随机方法
func randb(r int) int {
	rand.Seed(time.Now().UnixNano())
	res := rand.Intn(r)
	if res == r {
		res -= 1
	}
	return res
}

// 发送Get请求
func getHTML(link string, host string) *goquery.Document {
	var client = &http.Client{}
	r, err := http.NewRequest("GET", link, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Accept", "text/html")
	r.Header.Add("Accept-Language", "en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	r.Header.Add("Cookie", `bid=u-d5askqrT8; gr_user_id=e75422bb-f3a5-47c8-a08d-124425314379; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03=f29d9a55-e90c-475c-aece-414df2fdeb82; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03_f29d9a55-e90c-475c-aece-414df2fdeb82=false; gr_cs1_f29d9a55-e90c-475c-aece-414df2fdeb82=user_id%3A0; _vwo_uuid_v2=D3F9864D988439E45E37BBF2747F83FB8|22a3c044a53ddab8a66a5cb7dca023ad; viewed="4913065_4913064_2201813"`)
	r.Header.Add("Host", host)
	r.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36`)
	r.Header.Add("sec-ch-ua", `"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	r.Header.Add("sec-ch-ua-mobile", "?0")
	r.Header.Add("Sec-Fetch-Dest", "document")
	r.Header.Add("Sec-Fetch-Mode", "navigate")
	r.Header.Add("Sec-Fetch-Site", "none")
	r.Header.Add("Sec-Fetch-User", "?1")
	r.Header.Add("Upgrade-Insecure-Requests", "1")
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}
	return doc
}

// 发送POST请求(application/x-www-form-urlencoded)
func postHTML(link string, data string) *goquery.Document {
	res, err := http.Post(link,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		panic(err)
	}
	return doc
}
