package octopus

func (o *octopus) makeFilterCrawlDepthPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(o.filterByUrlDepth, outChSet)
}

func (o *octopus) filterByUrlDepth(node *Node, outChSet *NodeChSet) {
	if node.Depth < o.MaxCrawlDepth {
		outChSet.NodeCh <- node
	}
}
