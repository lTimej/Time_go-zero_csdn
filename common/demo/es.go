package demo

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

package main

import (
"context"
"fmt"
"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
	err    error
)

func init() {
	client, err = elastic.NewClient(elastic.SetURL("http://172.20.16.20:9200"))
	if err != nil {
		// Handle error
		fmt.Println(err, "链接es错误")
		return
	}
}

func Create(indices string) {
	exists, err := client.IndexExists(indices).Do(context.Background())
	if err != nil {
		// Handle error
	}
	if !exists {
		// Index does not exist yet.
		fmt.Println("索引不存在")
		mapping := `
		{
			"mappings":{
				"properties":{
					"user":{
						"type":"keyword"
					},
					"message":{
						"type":"keyword"
					},
					"retweets":{
						"type":"keyword"
					},
					"tags":{
						"type":"keyword"
					},
					"suggest_field":{
						"type":"completion"
					}
				}
			}
		}
		`
		createIndex, err := client.CreateIndex(indices).Body(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			fmt.Println(createIndex, "创建索引失败")
		}
	}
}

type Hello struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Desc     string `json:"desc"`
}

func Get() {
	ctx := context.Background()
	suggester := elastic.NewCompletionSuggester("my-suggest").Fuzziness(0).
		Text("hel").Field("desc").SkipDuplicates(true)
	searchSource := elastic.NewSearchSource().
		Suggester(suggester).
		FetchSource(false).
		TrackScores(true)
	searchResult, err := client.Search().
		Index("t1").
		SearchSource(searchSource).
		Do(ctx)
	if err != nil {
		fmt.Println(err, "====")
	}
	fmt.Println(searchResult.Suggest["my-suggest"][0].Text)
	fmt.Println(searchResult.Suggest["my-suggest"][0].Offset)
	fmt.Println(searchResult.Suggest["my-suggest"][0].Length)
	fmt.Println(searchResult.Suggest["my-suggest"][0].Options)
	options := searchResult.Suggest["my-suggest"][0].Options
	for _, option := range options {
		fmt.Println(option.Text)
		//fmt.Println(option.Source)
	}
	//var ttyp Hello
	//for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
	//	t := item.(Hello)
	//	fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
	//}
}


