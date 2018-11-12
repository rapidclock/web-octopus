package main

import (
	"github.com/rapidclock/web-octopus/adapter"
	"github.com/rapidclock/web-octopus/octopus"
)

func main() {
	opAdapter := &adapter.StdOpAdapter{}
	options := octopus.GetDefaultCrawlOptions()
	options.OpAdapter = opAdapter
	options.MaxCrawledUrls = 150
	crawler := octopus.New(options)
	crawler.SetupSystem()
	crawler.BeginCrawling("https://www.macrumors.com")
}
