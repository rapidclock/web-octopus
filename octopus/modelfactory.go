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
		CrawlOptions: GetDefaultCrawlOptions(),
		visited:      new(sync.Map),
		isReady:      false,
	}
	return oct
}

// New - Create an Instance of the Octopus with the given CrawlOptions.
func New(opt *CrawlOptions) *octopus {
	oct := &octopus{
		CrawlOptions: opt,
		visited:      new(sync.Map),
		isReady:      false,
	}
	return oct
}

func GetDefaultCrawlOptions() *CrawlOptions {
	return &CrawlOptions{
		MaxCrawlDepth:      -1,
		MaxCrawlLinks:      -1,
		StayWithinBaseHost: false,
		CrawlRatePerSec:    -1,
		RespectRobots:      false,
		IncludeBody:        true,
		OpAdapter:          nil,
	}
}
