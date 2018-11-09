package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/rapidclock/web-octopus/adapter"
	exp "github.com/rapidclock/web-octopus/experimental"
	oct "github.com/rapidclock/web-octopus/octopus"
)

const (
	HomeUrl     = "https://en.wikipedia.org/wiki/Main_Page"
	LessLinkUrl = "https://vorozhko.net/get-all-links-from-html-page-with-go-lang"
	Url1        = "https://benjamincongdon.me/blog/2018/03/01/Scraping-the-Web-in-Golang-with-Colly-and-Goquery"
	Url2        = "https://www.devdungeon.com/content/web-scraping-go"
)

func main() {
	// exp.Test_makeLinksAbsolute()
	// runPipeline()
	// runPipelineWithOptions()
	// octopusTest()
	stupid()
}

func stupid() {
	// resp, err := http.Head(HomeUrl)
	// if err == nil && resp == nil {
	// 	log.Fatal("WOW resp is nill although err is not")
	// }
	// if err == nil && resp != nil && resp.StatusCode == 200 {
	// 	fmt.Printf("%s\n", resp.Status)
	// }
	resp, err := http.Head("https://en.wikipedia.org/wiki/Main_Page")
	fmt.Println("A")
	if err == nil && resp == nil {
		log.Fatal("WOW resp is nill although err is not")
	}
	fmt.Println("B")
	if err == nil && resp.StatusCode == 200 {
		fmt.Printf("\nXX%s\n", resp.Status)
	}
	fmt.Println("C")
}
func octopusTest() {
	outputAdapter := &adapter.StdOpAdapter{}
	// outputAdapter := &adapter.FileWriterAdapter{"crawl_output.txt"}

	crawlOpt := oct.GetDefaultCrawlOptions()
	crawlOpt.MaxCrawlDepth = 3
	crawlOpt.OpAdapter = outputAdapter

	octopus := oct.New(crawlOpt)
	octopus.SetupSystem()
	octopus.BeginCrawling(Url1)
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
	crawler.StartCrawling(HomeUrl)
}
