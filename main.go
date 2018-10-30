package main

import (
	"fmt"
	exp "github.com/rapidclock/web-octopus/experimental"
	"log"
	"net/http"
	"time"
)

const (
	HomeUrl     = "https://en.wikipedia.org/wiki/Main_Page"
	LessLinkUrl = "https://vorozhko.net/get-all-links-from-html-page-with-go-lang"
	Url1        = "https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery"
	Url2        = "https://www.devdungeon.com/content/web-scraping-go"
)

func main() {
	//exp.Test_makeLinksAbsolute()
	//runPipeline()
	runPipelineWithOptions()
}

func checkPipelineA() {
	exp.Temp()
	crawler := exp.NewEngine()
	crawler.Consume(HomeUrl)
	crawler.Consume("Chicken Nugget")
	crawler.Consume("Tina kicks ass!!")
	fmt.Printf("1. %v\n", crawler.IsRunning())
	crawler.TurnOff()
	<-time.After(200 * time.Millisecond)
	fmt.Printf("2. %v\n", crawler.IsRunning())
	crawler.RestartEngine()
	<-time.After(100 * time.Millisecond)
	fmt.Printf("3. %v\n", crawler.IsRunning())
	crawler.Consume("Hello Poppet!!")
	fmt.Printf("4. %v\n", crawler.IsRunning())
	crawler.TurnOff()
}

func checkParsing() {
	resp, err := http.Get(HomeUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	results := exp.GetLinks(resp.Body)
	fmt.Println(len(results))
	results = exp.ProcessLinks(HomeUrl, results)
	fmt.Println(len(results))
	results = exp.RemoveDuplicates(results)
	fmt.Println(len(results))
	results = exp.ValidateLinks(results)
	fmt.Println(len(results))
	for _, v := range results {
		fmt.Println(v)
	}
}

func runPipeline() {
	crawler := exp.NewMonster()
	opAdapterPipe := exp.GetOutputAdapterPipe()
	crawler.BuildSystem(opAdapterPipe)
	crawler.StartCrawling(LessLinkUrl)
}

func runPipelineWithOptions() {
	opt := &exp.Options{
		MaxDepth: 1,
	}
	crawler := exp.NewMonsterWithOptions(opt)
	opAdapterPipe := exp.GetOutputAdapterPipe()
	crawler.BuildSystem(opAdapterPipe)
	crawler.StartCrawling(LessLinkUrl)
}
