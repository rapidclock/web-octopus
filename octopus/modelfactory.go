package octopus

import "sync"

const (
	defaultMaxDepth   int64 = 2
	anchorTag               = "a"
	anchorAttrb             = "href"
	defaultTimeToQuit       = 5
	defaultCrawlLimit int64 = -1
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

func GetDefaultCrawlOptions() *CrawlOptions {
	return &CrawlOptions{
		MaxCrawlDepth:      defaultMaxDepth,
		MaxCrawlLinks:      defaultCrawlLimit,
		StayWithinBaseHost: false,
		CrawlRatePerSec:    -1,
		RespectRobots:      false,
		IncludeBody:        true,
		OpAdapter:          nil,
		ValidProtocols:     []string{"http", "https"},
		TimeToQuit:         defaultTimeToQuit,
	}
}

func MakeNodeChSet(nodeCh chan<- *Node, quitCh chan<- int) *NodeChSet {
	return &NodeChSet{
		NodeCh: nodeCh,
		StdChannels: &StdChannels{
			QuitCh: quitCh,
		},
	}
}
