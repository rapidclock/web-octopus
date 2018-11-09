package main

import (
	"github.com/rapidclock/web-octopus/adapter"
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
	octopusTest()
}

func octopusTest() {
	outputAdapter := &adapter.StdOpAdapter{}
	// outputAdapter := &adapter.FileWriterAdapter{"crawl_output.txt"}

	crawlOpt := oct.GetDefaultCrawlOptions()
	crawlOpt.OpAdapter = outputAdapter

	octopus := oct.New(crawlOpt)
	octopus.SetupSystem()
	octopus.BeginCrawling(Url1)
}
