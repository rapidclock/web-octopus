package octopus

import (
	"net/url"
)

func (o *octopus) makeUrlProtocolFilterPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(o.filterByProtocol, outChSet,
		"Link Protocol Filter")
}

func (o *octopus) filterByProtocol(node *Node, outChSet *NodeChSet) {
	if node.UrlString != "" {
		linkUrl, err := url.Parse(node.UrlString)
		if err == nil && o.isValidProtocol[linkUrl.Scheme] {
			outChSet.NodeCh <- node
		}
	}
}
