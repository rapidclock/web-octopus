package octopus

import "sync"

const (
	defaultMaxDepth int16 = 2
	anchorTag             = "a"
	anchorAttrb           = "href"
)

// NewWithDefaultOptions - Create an Instance of the Octopus with the default CrawlOptions.
func NewWithDefaultOptions() *octopus {
	oct := &octopus{
		CrawlOptions: getDefaultCrawlOptions(),
		visited:      new(sync.Map),
		isBuilt:      false,
	}
	oct.setup()
	return oct
}

// New - Create an Instance of the Octopus with the given CrawlOptions.
func New(opt *CrawlOptions) *octopus {
	oct := &octopus{
		CrawlOptions: opt,
		visited:      new(sync.Map),
		isBuilt:      false,
	}
	return oct
}

func getDefaultCrawlOptions() *CrawlOptions {
	return &CrawlOptions{
		MaxDepthCrawled:    -1,
		MaxLinksCrawled:    -1,
		StayWithinBaseHost: false,
		CrawlRatePerSec:    -1,
		RespectRobots:      false,
		IncludeBody:        true,
		OpAdapter:          nil,
	}
}
