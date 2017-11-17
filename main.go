package main

import (
	"sync"
	//"fmt"
)
var wg sync.WaitGroup


func main() {
	//fmt.Print("Hello World\n")

	//spider := Spider{}
	//spider.Search()

	//index := NovelIndex{}
	//index.IndexFile("./novel")
	////index.SearchText()

	//go_spider_run()

	run_spider_zhihu()
}


