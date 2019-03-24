package octopus

import "sync"

const (
	defaultMaxDepth       int64  = 2
	anchorTag                    = "a"
	anchorAttrb                  = "href"
	defaultTimeToQuit            = 30
	defaultLinkCrawlLimit int64  = -1
	defaultCrawlRateLimit int64  = -1
	defaultRequestTimeout uint64 = 15
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

func createNode(parentUrlStr, urlStr string, depth int64) *Node {
	return &Node{
		NodeInfo: &NodeInfo{
			ParentUrlString: parentUrlStr,
			UrlString:       urlStr,
			Depth:           depth,
		},
		Body: nil,
	}
}

// Returns an instance of CrawlOptions with the values set to sensible defaults.
func GetDefaultCrawlOptions() *CrawlOptions {
	return &CrawlOptions{
		MaxCrawlDepth:         defaultMaxDepth,
		MaxCrawledUrls:        defaultLinkCrawlLimit,
		StayWithinBaseHost:    false,
		CrawlRatePerSec:       defaultCrawlRateLimit,
		CrawlBurstLimitPerSec: defaultCrawlRateLimit,
		RespectRobots:         false,
		IncludeBody:           true,
		OpAdapter:             nil,
		ValidProtocols:        []string{"http", "https"},
		TimeToQuit:            defaultTimeToQuit,
	}
}

// Utility function to create a NodeChSet given a created Node and Quit Channel.
func MakeNodeChSet(nodeCh chan<- *Node, quitCh chan<- int) *NodeChSet {
	return &NodeChSet{
		NodeCh: nodeCh,
		StdChannels: &StdChannels{
			QuitCh: quitCh,
		},
	}
}

// Utility to create a NodeChSet and get full access to the Quit & Node Channel.
func MakeDefaultNodeChSet() (*NodeChSet, chan *Node, chan int) {
	nodeCh := make(chan *Node)
	quitCh := make(chan int)
	return MakeNodeChSet(nodeCh, quitCh), nodeCh, quitCh
}
