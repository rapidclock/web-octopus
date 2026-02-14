package octopus

import "time"

func (o *octopus) makeRateLimitingPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(o.rateLimit, outChSet, "Crawl Rate Limit")
}

func (o *octopus) rateLimit(node *Node, outChSet *NodeChSet) {
	for {
		if r := o.rateLimiter.Reserve(); !r.OK() {
			continue
		} else {
			time.Sleep(r.Delay())
			outChSet.NodeCh <- node
			return
		}
	}
}
