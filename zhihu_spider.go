package main

import (
	"fmt"
	goquery "github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type ZhihuSpider struct{
	SpiderBase
	base_url string
}

func (spider* ZhihuSpider) SearchTopic(id string, pageNumber string)(){
	url := spider.base_url + "/topic/" + id + "/top-answers?page=" + pageNumber
	doc, err := goquery.NewDocument(url)

	if err != nil{
		fmt.Println("error in query ", url, err)
		return
	}
	//fmt.Println(doc.Html())

	doc.Find(".zm-item-rich-text").Each(func(i int, s *goquery.Selection) {
		title := spider.Translate(s.Text())

		srcUrl, exists := s.Attr("data-entry-url")
		if !exists {
			fmt.Println("not find href")
		}

		srcUrl = spider.base_url + srcUrl
		spider.SearchQuestion(srcUrl)


	})
}

func (spider* ZhihuSpider) SearchQuestion(url string){
	doc, err := goquery.NewDocument(url)

	if err != nil{
		fmt.Println("error in query ", url, err)
		return
	}
	//fmt.Println(url)


	//return
	//
	//doc.Find(".zm-item-rich-text").Each(func(i int, s *goquery.Selection) {
	//	title := spider.Translate(s.Text())
	//	fmt.Println(title)
	//
	//	srcUrl, exists := s.Find("a").Attr("data-entry-url")
	//	if !exists {
	//		fmt.Println("not find href")
	//	}
	//
	//	srcUrl = spider.base_url + srcUrl
	//	spider.SearchQuestion(srcUrl)
	//
	//
	//})
}

func (spider* ZhihuSpider) ParseQuestion(result string){

}

func (spider* ZhihuSpider) SearchComment(id string){
	
}

func (spider* ZhihuSpider) ParseComment(result string){

}

func run_spider_zhihu()  {
	var spider ZhihuSpider
	spider.base_url = "http://www.zhihu.com"
	for pageNumber := 1; pageNumber < 50; pageNumber++ {
		spider.SearchTopic("19552204", strconv.Itoa(pageNumber))
	}
}
