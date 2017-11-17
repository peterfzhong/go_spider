package main

import (
	"net/http"
	"io/ioutil"
	//"fmt"
	iconv "github.com/djimenez/iconv-go"
)

type SpiderBase struct {

}

func  (spider* SpiderBase) HttpGet(url string)(content string, statusCode int){
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
	//fmt.Println(content)
	//content, _  = iconv.ConvertString(content, "gbk", "utf-8")

	return
}

func (spider *SpiderBase) Translate(src string)(dest string){
	text, err := iconv.ConvertString(src, "gbk", "utf-8")
	dest = src
	if err!= nil{
		//fmt.Println("err in convert string, src: ", src)
		return
	}

	dest = text

	return
}
