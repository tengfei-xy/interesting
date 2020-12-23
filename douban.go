package main

import (
	"fmt"
)

type douban struct {
	Type       string `json:"type"`
	bTagLink   string
	bRangeLink string
	BookName   string `json:"book"`
	BookLink   string `json:"link"`
}

func doubanMain() douban {
	var book douban

	fmt.Println("搜索对象:豆瓣")

	// 获取书籍类型
	book.getType()

	// 获取书籍信息
	book.getInfo()
	return book
}

const doubanHost string = "book.douban.com"

// 获取书籍类型
func (book *douban) getType() {
	doc := getHTML("https://book.douban.com/tag/?view=type", doubanHost)
	tag := doc.Find("div[class=article]").Find("table[class=tagCol]").Find("td>a")
	r := randb(tag.Length())

	book.bTagLink, _ = tag.Eq(r).Attr("href")
	book.Type = tag.Eq(r).Text()

	fmt.Printf("书籍类型:%s\n", book.Type)
}

// 获取书籍信息
func (book *douban) getInfo() {
	// 由于豆瓣只能获得前1k书籍内容
	doc := getHTML("https://book.douban.com"+fmt.Sprintf("%s?start=%d&type=T", book.bTagLink, randb(999)), doubanHost)
	info := doc.Find("#subject_list").Find("div[class=info]")
	r := randb(info.Length())
	book.BookName, _ = info.Eq(r).Find("a").Eq(0).Attr("title")
	book.BookLink, _ = info.Eq(r).Find("a").Eq(0).Attr("href")

	fmt.Printf("书籍名称:%s(%s)\n\n", book.BookName, book.BookLink)
}
