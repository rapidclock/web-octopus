# web-octopus
[![Build Status](https://travis-ci.com/rapidclock/web-octopus.svg?token=hJhLfHtyz41UyuLTTdFx&branch=master)](https://travis-ci.com/rapidclock/web-octopus)
<br>
A concurent web crawler to crawl the web.

## Current Features:
- Depth Limited Crawling
- User specified valid protocols
- User buildable adapters that the crawler feeds output to.
- Filter Duplicates.
- Filter URLs that fail a HEAD request.
- User specifiable max timeout between two successive url requests.


### Sample Implementation Snippet

```go
package main

import (
	"github.com/rapidclock/web-octopus/adapter"
	"github.com/rapidclock/web-octopus/octopus"
)

func main() {
	opAdapter := &adapter.StdOpAdapter{}
	
	options := octopus.GetDefaultCrawlOptions()
	options.MaxCrawlDepth = 3
	options.TimeToQuit = 10
	options.OpAdapter = opAdapter
	
	crawler := octopus.New(options)
	crawler.SetupSystem()
	crawler.BeginCrawling("https://www.example.com")
}
```