package main

import (
	"fmt"
	"strconv"
	"strings"

	iconv "github.com/djimenez/iconv-go"
)

type nbxhsd struct {
	Type       string `json:"type"`
	kindId     string
	page       int
	pageTotal  int
	bTagLink   string
	bRangeLink string
	BookName   string `json:"book"`
	BookLink   string `json:"link"`
}

const nbxhsdHost string = "ningbo.zxhsd.com"

func nbxhsdMain() nbxhsd {
	var book nbxhsd

	fmt.Println("搜索对象:浙江新华书店")
	// 获取书籍类型
	book.getType()
	// 获取书籍类型总页数
	book.getPageTotal()

	// 获取书籍信息
	book.getInfo()
	return book
}

// 获取书籍类型
func (book *nbxhsd) getType() {
	doc := getHTML("http://ningbo.zxhsd.com/", nbxhsdHost)
	btype := doc.Find("div[class=catecon]>h3:first-child+ul>li>div a")
	r := randb(btype.Length())

	book.bTagLink, _ = btype.Eq(r).Attr("href")
	// http://www.zxhsd.com/search/book_search.jsp?kindId=AFAG0,AFAD0 -> AFAG0%2CAFAD0
	book.kindId = strings.Replace(strings.Split(book.bTagLink, "=")[1], ",", "0%2", 1)
	book.Type = cnstring(btype.Eq(r).Text())

	fmt.Printf("书籍类型:%s\n", book.Type)
	return
}

// 获取书籍类型总页数
func (book *nbxhsd) getPageTotal() {
	doc := getHTML(book.bTagLink, nbxhsdHost)
	book.pageTotal, _ = strconv.Atoi(strings.TrimSpace(doc.Find("div[class=meneame]>span").Text()))
	book.page = randb(book.pageTotal)
}

func (book *nbxhsd) getInfo() {
	data := fmt.Sprintf("sTotalPage=%d&isbn=&sm=&zz=&cbs=&ywdl=&ywxl=&jgfwks=&jgfwjs=&tempcbsjks=&tempcbsjjs=&sjsjks=&sjsjjs=&zkks=0&zkjs=100&kindId=%s&orderby=jlrq&category=&keyword=&curpage=%d", book.pageTotal, book.kindId, book.page)
	doc := postHTML("http://www.zxhsd.com/search/book_search.jsp", data)
	booklist := doc.Find("li[class=li_content]>dl>dt>a")
	seq := randb(booklist.Length())

	book.BookName = cnstring(booklist.Eq(seq).Text())
	link, _ := booklist.Eq(seq).Attr("href")
	book.BookLink = cnstring(link)
	fmt.Printf("书籍名称:%s(%s)\n\n", book.BookName, book.BookLink)
}
func cnstring(source string) string {
	output, err := iconv.ConvertString(source, "GB2312", "utf-8")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return output
}
