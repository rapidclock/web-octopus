package octopus

import (
	"io"
	"sync"
)

// octopus is a concurrent web crawler.
// It has an inbuilt parser based of html.NewTokenizer to collect all links in a web-page.
// It also has a CrawlOptions structure to initialize setting specific
// to an instance of the crawler.
type octopus struct {
	*CrawlOptions
	visited      *sync.Map
	isReady      bool
	adapterChSet *NodeChSet
}

// CrawlOptions is used to house options for crawling.
// You can specify depth of exploration for each link,
// if crawler should ignore other hostnames (except from base host).
// MaxCrawlDepth - Indicates the maximum depth that will be crawled,
// for each new link.
// MaxCrawlLinks - Specifies the Maximum Number of Unique Links that will be crawled.
// Note : When combined with DepthPerLink, it will combine both.
// Use -1 to indicate infinite links to be crawled (only bounded by depth of traversal).
// IncludeBody - Include the response Body in the crawled NodeInfo (for further processing).
// OpAdapter is a user specified concrete implementation of an Output Adapter. The crawler
// will pump output onto the implementation's channel returned by its Consume method.
// CrawlRate is the rate at which requests will be made.
// RespectRobots (unimplemented) choose whether to respect robots.txt or not.
type CrawlOptions struct {
	MaxCrawlDepth      int64
	MaxCrawlLinks      int64
	StayWithinBaseHost bool
	CrawlRatePerSec    int64
	RespectRobots      bool
	IncludeBody        bool
	OpAdapter          *OutputAdapter
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

type StdChannels struct {
	QuitCh chan<- int
	// logCh     chan<- string
	// errorCh   chan<- string
}

type NodeChSet struct {
	NodeCh chan<- *Node
	*StdChannels
}

type StringChSet struct {
	StrCh chan<- string
	*StdChannels
}

// OutputAdapter is the interface for the Adapter that is used to handle
// output from the Octopus Crawler.
// The contract stipulates that the crawler provides the channel
// to listen for a quit command.
// The crawler pumps its output onto the returned channel of the Consume method.
// Implementers of the interface should listen on this channel for output from
// the crawler.
type OutputAdapter interface {
	Consume() *NodeChSet
}
