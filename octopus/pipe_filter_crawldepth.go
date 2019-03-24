package octopus

func (o *octopus) makeCrawlDepthFilterPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(o.filterByUrlDepth, outChSet, "Crawl Depth Filter")
}

func (o *octopus) filterByUrlDepth(node *Node, outChSet *NodeChSet) {
	if node.Depth < o.MaxCrawlDepth {
		outChSet.NodeCh <- node
	}
}
