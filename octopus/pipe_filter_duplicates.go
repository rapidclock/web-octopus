package octopus

func (o *octopus) makeDuplicateUrlFilterPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(o.filterDuplicates, outChSet, "Dupliacate Filter")
}

func (o *octopus) filterDuplicates(node *Node, outChSet *NodeChSet) {
	if _, visited := o.visited.Load(node.UrlString); !visited {
		o.visited.Store(node.UrlString, true)
		outChSet.NodeCh <- node
	}
}
