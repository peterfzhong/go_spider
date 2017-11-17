package main

import ("fmt"
	"net/http"
	"io/ioutil"
	iconv "github.com/djimenez/iconv-go"
	goquery "github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
)

type Spider  struct {

}

func  (spider* Spider) HttpGet(url string)(content string, statusCode int){
	//var content	string = ""
	resp, err1 := http.Get(url)
	if err1 != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	content = string(data)

	content, _  = iconv.ConvertString(content, "gbk", "utf-8")

	return
}

func (spider *Spider) Translate(src string)(dest string){
	text, err := iconv.ConvertString(src, "gbk", "utf-8")
	dest = src
	if err!= nil{
		fmt.Println("err in convert string, src: ", src)
		return
	}

	dest = text

	return
}

func (spider* Spider) Search()(){
	for i := 1; i <= 24; i++{
		url := "http://cl.ciko.pw/thread0806.php?fid=20&search=&page=" + fmt.Sprintf("%d", i)
		fmt.Println(url)

		wg.Add(1)
		go spider.SearchGo(url)

	}
	wg.Wait()
}

func (spider* Spider) SearchGo(url string){
	defer wg.Done()
	doc, err := goquery.NewDocument(url)

	if err != nil{
		fmt.Println("error in query ", url, err)
		return
	}

	fmt.Println(doc.Html())

	doc.Find("h3").Each(func(i int, s *goquery.Selection) {
		title := spider.Translate(s.Text())
		fmt.Println(title)

		srcUrl, exists := s.Find("a").Attr("href")
		if !exists {
			fmt.Println("not find href")
		}

		srcUrl = "http://cl.ciko.pw/" + srcUrl
		content := spider.SearchNovel(srcUrl)

		spider.SaveNovelTxt(title, content)

	})
}

func (spider* Spider) SearchNovel(url string)(content string){
	fmt.Println("url: ", url, "\r\n\r\n")

	doc, err := goquery.NewDocument(url)
	content = url + "\r\n\r\n"
	if err != nil{
		fmt.Println("error in query ", url, err)
		return
	}

	doc.Find("div.tpc_content").Each(func(i int, s *goquery.Selection) {
		if len(s.Text()) > 50{
			content += spider.Translate(s.Text()) + "\r\n\r\n"
			content = strings.Replace(content, "  ", "\r\n", -1)
			content = strings.Replace(content, "　　", "\r\n", -1)
		}
	})
	//fmt.Println(content)

	totalPage := 0
	doc.Find("#last").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("href")
		len := len(val)
		totalPage, _ = strconv.Atoi(val[strings.Index(val, "page=")+5: len])
	})
	fmt.Println("total page: ", totalPage)

	content += spider.SearchNovelPage(url, totalPage)

	return
}

func (spider* Spider) SearchNovelPage(baseUrl string, totalPage int)(content string){
	start := strings.LastIndex(baseUrl, "/") + 1
	end := strings.Index(baseUrl, ".html")

	if start == -1 || end == -1{
		fmt.Println("baseUrl not used: ", baseUrl)
		return
	}

	fmt.Println("base url; ", baseUrl)
	fmt.Println(start, ", ", end)
	seedId := baseUrl[start: end]
	baseSeedUrl := fmt.Sprintf("http://cl.ciko.pw/read.php?tid=%s",seedId)

	for i := 2; i <= totalPage; i++{
		url := fmt.Sprintf("%s&page=%d", baseSeedUrl, i)
		fmt.Println(url)

		doc, err := goquery.NewDocument(url)
		if err != nil{
			fmt.Println("error in query ", url, err)
			return
		}
		fmt.Println(doc.Html())

		doc.Find("div.tpc_content").Each(func(i int, s *goquery.Selection) {
			if len(s.Text()) > 50 {
				content += spider.Translate(s.Text()) + "\r\n\r\n"
				content = strings.Replace(content, "  ", "\r\n", -1)
				content = strings.Replace(content, "　　", "\r\n", -1)
				fmt.Println("xxxxxxxxxxxxxxxxxxxx, ", content)
			}
		})

		//fmt.Println(content)

	}
	return
}

func (spider* Spider) SaveNovelTxt(title string, content string)  {
	file_name := "./novel/" + title + ".txt"
	err := ioutil.WriteFile(file_name, []byte(content), 0666)
	if err != nil{
		fmt.Println("error in ioutil.WriteFile, ", file_name)
	}
}


