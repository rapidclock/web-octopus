package octopus

import (
	"io"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// octopus is a concurrent web crawler.
// It has an inbuilt parser based of html.NewTokenizer to collect all links in a web-page.
// It also has a CrawlOptions structure to initialize setting specific
// to an instance of the crawler.
type octopus struct {
	*CrawlOptions
	visited           *sync.Map
	isReady           bool
	adapterChSet      *NodeChSet
	isValidProtocol   map[string]bool
	timeToQuit        time.Duration
	inputUrlStrChan   chan string
	masterQuitCh      chan int
	crawledUrlCounter int64
	rateLimiter       *rate.Limiter
	requestTimeout    uint64
}

// CrawlOptions is used to house options for crawling.
//
// You can specify depth of exploration for each link,
// if crawler should ignore other host names (except from base host).
//
// 	MaxCrawlDepth - Indicates the maximum depth that will be crawled,
// 	for each new link.
//
// 	MaxCrawledUrls - Specifies the Maximum Number of Unique Links that will be crawled.
// 	Note : When combined with DepthPerLink, it will combine both.
// 	Use -1 to indicate infinite links to be crawled (only bounded by depth of traversal).
//
// 	StayWithinBaseHost - (unimplemented) Ensures crawler stays within the
// 	level 1 link's hostname.
//
// 	CrawlRatePerSec - is the rate at which requests will be made (per second).
// 	If this is negative, Crawl feature will be ignored. Default is negative.
//
// 	CrawlBurstLimitPerSec - Represents the max burst capacity with which requests
// 	can be made. This must be greater than or equal to the CrawlRatePerSec.
//
// 	RespectRobots (unimplemented) choose whether to respect robots.txt or not.
//
// 	IncludeBody - (unimplemented) Include the response Body in the crawled
// 	NodeInfo (for further processing).
//
// 	OpAdapter is a user specified concrete implementation of an Output Adapter. The crawler
// 	will pump output onto the implementation's channel returned by its Consume method.
//
// 	ValidProtocols - This is an array containing the list of url protocols that
// 	should be crawled.
//
// 	TimeToQuit - represents the total time to wait between two new nodes to be
// 	generated before the crawler quits. This is in seconds.
type CrawlOptions struct {
	MaxCrawlDepth         int64
	MaxCrawledUrls        int64
	StayWithinBaseHost    bool
	CrawlRatePerSec       int64
	CrawlBurstLimitPerSec int64
	RespectRobots         bool
	IncludeBody           bool
	OpAdapter             OutputAdapter
	ValidProtocols        []string
	TimeToQuit            int64
}

// NodeInfo is used to represent each crawled link and its associated crawl depth.
type NodeInfo struct {
	ParentUrlString string
	UrlString       string
	Depth           int64
}

// Node encloses a NodeInfo and its Body (HTML) Content.
type Node struct {
	*NodeInfo
	Body io.ReadCloser
}

// StdChannels are used to hold the standard set of channels that are used
// for special operations. Will include channels for Logging, Statistics,
// etc. in the future.
type StdChannels struct {
	QuitCh chan<- int
	// logCh     chan<- string
	// errorCh   chan<- string
}

// NodeChSet is the standard set of channels used to build the concurrency
// pipelines in the crawler.
type NodeChSet struct {
	NodeCh chan<- *Node
	*StdChannels
}

type ingestPipeChSet struct {
	NodeCh chan *Node
	StrCh  chan string
	QuitCh chan int
}

// OutputAdapter is the interface that has to be implemented in order to
// handle outputs from the octopus crawler.
//
// The octopus will call the OutputAdapter.Consume(
// ) method and deliver all relevant output and quit signals on the channels
// included in the received NodeChSet.
//
// This implies that it is the responsibility of the user who implements
// OutputAdapter to handle processing the output of the crawler that is
// delivered on the NodeChSet.NodeCh.
//
// Implementers of the interface should listen to the included channels in
// the output of Consume() for output from the crawler.
type OutputAdapter interface {
	Consume() *NodeChSet
}
