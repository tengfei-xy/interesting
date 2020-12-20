package main
import(
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
    "math/rand"
    "time"
)
type douban struct {
	Type string `json:"type"`
	bTagLink string
	bRangeLink string
	BookName string `json:"book"`
	BookLink string `json:"link"`
}

func interesting() douban{
	var book douban
	rand.Seed(time.Now().Unix())

	// 获取书籍类型
	book.getType()

	// 获取书籍信息
	book.getInfo()
	return book
}

// 获取书籍类型
func (book * douban)getType()  {
	doc := reqHtml("https://book.douban.com/tag/?view=type")
	tag := doc.Find("div[class=article]").Find("table[class=tagCol]").Find("td>a")
	r := randb(tag.Length())
	
	book.bTagLink,_ = tag.Eq(r).Attr("href")
	book.Type = tag.Eq(r).Text()

	fmt.Printf("书籍类型:%s\n",book.Type)
}

// 获取书籍信息
func (book * douban)getInfo(){
	// 由于豆瓣只能获得前1k书籍内容
	doc := reqHtml("https://book.douban.com" + fmt.Sprintf("%s?start=%d&type=T",book.bTagLink,randb(999)))
	info := doc.Find("#subject_list").Find("div[class=info]")
	r := randb(info.Length())
	book.BookName,_ = info.Eq(r).Find("a").Eq(0).Attr("title")
	book.BookLink,_ = info.Eq(r).Find("a").Eq(0).Attr("href")
	
	fmt.Printf("书籍名称:%s(%s)\n\n",book.BookName,book.BookLink)
}

// 随机方法
func randb(r int)int{
	res := rand.Intn(r)
	if res == r{
		res -= 1
	}
	return res
}
// 发送Get请求
func reqHtml(link string) *goquery.Document{
	var client			= &http.Client{}
	r, err := http.NewRequest("GET", link, nil)
	if err != nil {
		panic(err)
	}
	r.Header.Add("Accept","text/html")
	r.Header.Add("Accept-Language","en-US,en;q=0.9,zh-CN;q=0.8,zh;q=0.7")
	r.Header.Add("Cookie",`bid=u-d5askqrT8; gr_user_id=e75422bb-f3a5-47c8-a08d-124425314379; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03=f29d9a55-e90c-475c-aece-414df2fdeb82; gr_session_id_22c937bbd8ebd703f2d8e9445f7dfd03_f29d9a55-e90c-475c-aece-414df2fdeb82=false; gr_cs1_f29d9a55-e90c-475c-aece-414df2fdeb82=user_id%3A0; _vwo_uuid_v2=D3F9864D988439E45E37BBF2747F83FB8|22a3c044a53ddab8a66a5cb7dca023ad; viewed="4913065_4913064_2201813"`)
	r.Header.Add("Host","book.douban.com")
	r.Header.Add("User-Agent",`Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36`)
	r.Header.Add("sec-ch-ua",`"Google Chrome";v="87", " Not;A Brand";v="99", "Chromium";v="87"`)
	r.Header.Add("sec-ch-ua-mobile","?0")
	r.Header.Add("Sec-Fetch-Dest","document")
	r.Header.Add("Sec-Fetch-Mode","navigate")
	r.Header.Add("Sec-Fetch-Site","none")
	r.Header.Add("Sec-Fetch-User","?1")
	r.Header.Add("Upgrade-Insecure-Requests","1")
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