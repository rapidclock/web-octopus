package octopus

import (
	"sync/atomic"
)

func (o *octopus) makeCrawlLinkCountLimitPipe(inChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(o.checkWithinLimit, inChSet, "Crawl Link Limit")
}

func (o *octopus) checkWithinLimit(node *Node, outChSet *NodeChSet) {
	if v := atomic.AddInt64(&o.crawledUrlCounter,
		1); v <= o.MaxCrawledUrls {
		outChSet.NodeCh <- node
	} else {
		outChSet.QuitCh <- 1
		o.masterQuitCh <- 1
	}
}
