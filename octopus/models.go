package octopus

import "io"

// Node is used to represent each crawled link and its associated depth of crawl.
type Node struct {
	URLString string
	Depth     int
}

// webOctopus is a concurrent version of webSpider.
// It has an inbuilt parser based of htmlparser.Parser to collect all links in a web-page.
// It also has a CrawlOptions structure to initialize setting specific
// to an instance of the crawler.
type webOctopus struct {
	CrawlOptions
	visited map[Node]bool
}

// CrawlOptions is used to house options for crawling.
// You can specify depth of exploration for each link,
// if crawler should ignore other hostnames (except from base host).
// MaxLinksCrawled - Specifies the Maximum Number of Unique Links that will be crawled.
// Note : When combined with DepthPerLink, it will combine both.
// Use -1 to indicate infinite links to be crawled (only bounded by depth of traversal).
// IncludeBody - Include the response Body in the crawled Node (for further processing).
// OpAdapter is a user specified concrete implementation of an Output Adapter. The crawler
// will pump output onto the implementation's channel returned by its Consume method.
type CrawlOptions struct {
	DepthPerLink       int16
	MaxLinksCrawled    int64
	StayWithinBaseHost bool
	BaseURLString      string
	IncludeBody        bool
	OpAdapter          OutputAdapter
}

type CrawlOutput struct {
	Node
	Body io.ReadCloser
}
