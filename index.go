package main

import (
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"fmt"
	"os"
	//"time"
	"path/filepath"
	"io/ioutil"
)

var (
	docId    uint64
)

const (
	text1 = `在苏黎世的FIFA颁奖典礼上，巴萨球星、阿根廷国家队队长梅西赢得了生涯第5个金球奖，继续创造足坛的新纪录`
	text2 = `12月6日，网上出现照片显示国产第五代战斗机歼-20的尾翼已经涂上五位数部队编号`
)

type NovelIndex struct {

}

func (index* NovelIndex) Index()(){
	// searcher是协程安全的
	searcher := engine.Engine{}

	searcher.Init(types.EngineInitOptions{
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.DocIdsIndex,
		},
		UsePersistentStorage:    true,
		PersistentStorageFolder: "./index",
		PersistentStorageShards: 8,
		//SegmenterDictionaries: "./dict/dictionary.txt",
		//StopTokenFile:         "./dict/stop_tokens.txt",
	})
	defer searcher.Close()

	docId++
	searcher.IndexDocument(docId, types.DocumentIndexData{Content: text1}, false)
	docId++
	searcher.IndexDocument(docId, types.DocumentIndexData{Content: text2}, false)

	searcher.FlushIndex()


	fmt.Printf("%#v\n", searcher.Search(types.SearchRequest{Text: "巴萨 梅西"}))
	fmt.Printf("%#v\n", searcher.Search(types.SearchRequest{Text: "战斗机 金球奖"}))
}

func (index *NovelIndex) IndexFile(srcpath string) () {
	searcher := engine.Engine{}

	searcher.Init(types.EngineInitOptions{
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.DocIdsIndex,
		},
		UsePersistentStorage:    true,
		PersistentStorageFolder: "/tmp/",
		PersistentStorageShards: 8,
		SegmenterDictionaries: "./dict/dictionary.txt",
		//StopTokenFile:         "./dict/stop_tokens.txt",
	})

	defer searcher.Close()

	var paths []string
	var docId uint64
	docId = 0
	//update index dynamically
	//srcpath := "./novel"
	filepath.Walk(srcpath, func(path string, f os.FileInfo, err error) error {
		fmt.Println("src file: ", path)
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		fc, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("read file:", path, "error:", err)
		}

		docId++
		fmt.Println("indexing file:", path, "... ...")
		searcher.IndexDocument(docId, types.DocumentIndexData{Content: string(fc)}, true)
		fmt.Println("indexed file:", path, " ok")
		paths = append(paths, path)

		return nil
	})

	searcher.FlushIndex()
	fmt.Println("%#v\n", searcher.Search(types.SearchRequest{Text: "老外"}))
}

func (index *NovelIndex) SearchText() () {
	searcher := engine.Engine{}

	searcher.Init(types.EngineInitOptions{
		IndexerInitOptions: &types.IndexerInitOptions{
			IndexType: types.DocIdsIndex,
		},
		UsePersistentStorage:    true,
		PersistentStorageFolder: "/tmp/",
		PersistentStorageShards: 8,
		SegmenterDictionaries: "./dict/dictionary.txt",
		//StopTokenFile:         "./dict/stop_tokens.txt",
	})

	defer searcher.Close()

	fmt.Println("%#v\n", searcher.Search(types.SearchRequest{Text: "老外"}))

}