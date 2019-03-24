# web-octopus
[![Build Status](https://travis-ci.com/rapidclock/web-octopus.svg?token=hJhLfHtyz41UyuLTTdFx&branch=master)](https://travis-ci.com/rapidclock/web-octopus)
<br>
A concurent web crawler to crawl the web.

## Current Features:
- Depth Limited Crawling
- User specified valid protocols
- User buildable adapters that the crawler feeds output to.
- Filter Duplicates. (Default, Non-Customizable)
- Filter URLs that fail a HEAD request. (Default, Non-Customizable)
- User specifiable max timeout between two successive url requests.
- Max Number of Links to be crawled.


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

### List of customizations

Customizations can be made by supplying the crawler an instance of `CrawlOptions`. The basic structure is shown below, with a brief explanation for each option.

```go
type CrawlOptions struct {
	MaxCrawlDepth      int64 // Max Depth of Crawl, 0 is the initial link.
	MaxCrawledUrls     int64 // Max number of links to be crawled in total.
	StayWithinBaseHost bool // [Not-Implemented-Yet]
	CrawlRate          int64 // Max Rate at which requests can be made (req/sec).
	CrawlBurstLimit    int64 // Max Burst Capacity (should be atleast the crawl rate).
	RespectRobots      bool // [Not-Implemented-Yet]
	IncludeBody        bool // Include the Request Body (Contents of the web page) in the result of the crawl.
	OpAdapter          OutputAdapter // A user defined crawl output handler (See next section for info).
	ValidProtocols     []string // Valid protocols to crawl (http, https, ftp, etc.)
	TimeToQuit         int64 // Timeout (seconds) between two attempts or requests, before the crawler quits.
}
```

A default instance of the `CrawlOptions` can be obtained by calling `octopus.GetDefaultCrawlOptions()`. This can be further customized by overriding individual properties.

### Output Adapters

An Output Adapter is the final destination of a crawler processed request. The output of the crawler is fed here, according to the customizations made before starting the crawler through the `CrawlOptions` attached to the crawler.

The `OutputAdapter` is a Go Interface, that has to be implemented by your(user-defined) processor.

```go
type OutputAdapter interface {
	Consume() *NodeChSet
}
```

The user has to implement the `Consume()` method that returns a __*pointer*__ to a `NodeChSet`. The `NodeChSet` is described below. The crawler uses the returned channel to send the crawl output. The user can start listening for output from the crawler.

**Note** : If the user chooses to implement their custom `OutputAdapter` **REMEMBER** to listen for the output on another go-routine. Otherwise you might block the crawler from running. Atleast begin the crawling on another go-routine before you begin processing output.

The structure of the `NodeChSet` is given below.

```go
type NodeChSet struct {
	NodeCh chan<- *Node
	*StdChannels
}

type StdChannels struct {
	QuitCh chan<- int
}

type Node struct {
	*NodeInfo
	Body io.ReadCloser
}

type NodeInfo struct {
	ParentUrlString string
	UrlString       string
	Depth           int64
}
```

You can use the utility function `MakeDefaultNodeChSet()` to get a `NodeChSet` built for you. This also returns the `Node` and quit channels. Example given below:

```go
var opNodeChSet *NodeChSet
var nodeCh chan *Node
var quitCh chan int
// above to demo the types. One can easily use go lang type erasure.
opNodeChSet, nodeCh, quitCh = MakeDefaultNodeChSet()
```

The user should supply the custom OutputAdapter as an argument to the `CrawlOptions`.

#### Default Output Adapters:

We supply two default Adapters for you to try out. They are not meant to be feature rich, but you can still use them. Their primary purpose is meant to be a demonstration of how to build and use a `OutputAdapter`.

1. `adapter.StdOpAdapter` : Writes the crawled output (only links, not body) to the standard output.
1. `adapter.FileWriterAdapter` : Writes the crawled output (only links, not body) to a supplied file.

#### Implementation of the `adapter.StdOpAdapter`:
We have supplied the implementation of `adapter.StdOpAdapter` below to get a rough idea of what goes into building your own adapter.

```go
// StdOpAdapter is an output adapter that just prints the output onto the
// screen.
//
// Sample Output Format is:
// 	LinkNum - Depth - Url
type StdOpAdapter struct{}

func (s *StdOpAdapter) Consume() *oct.NodeChSet {
	listenCh := make(chan *oct.Node)
	quitCh := make(chan int, 1)
	listenChSet := &oct.NodeChSet{
		NodeCh: listenCh,
		StdChannels: &oct.StdChannels{
			QuitCh: quitCh,
		},
	}
	go func() {
		i := 1
		for {
			select {
			case output := <-listenCh:
				fmt.Printf("%d - %d - %s\n", i, output.Depth, output.UrlString)
				i++
			case <-quitCh:
				return
			}
		}
	}()
	return listenChSet
}
```
