package octopus

import (
	"sync/atomic"
)

func (o *octopus) makeLimitCrawlPipe(inChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(o.checkWithinLimit, inChSet)
}

func (o *octopus) checkWithinLimit(node *Node, outChSet *NodeChSet) {
	if v := atomic.AddInt64(&o.crawledUrlCounter,
		1); v < o.MaxCrawledUrls {
		outChSet.NodeCh <- node
	} else {
		outChSet.QuitCh <- 1
	}
}
