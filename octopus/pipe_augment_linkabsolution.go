package octopus

import "net/url"

func (o *octopus) makeLinkAbsolutionPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(makeLinkAbsolute, outChSet, "Link Absolution")
}

func makeLinkAbsolute(node *Node, outChSet *NodeChSet) {
	if node == nil || outChSet == nil {
		return
	}
	if node.ParentUrlString != "" {
		linkUrl, err := url.Parse(node.UrlString)
		if err != nil {
			return
		}
		if !linkUrl.IsAbs() {
			baseUrl, err := url.Parse(node.ParentUrlString)
			if err != nil {
				return
			}
			absLinkUrl := baseUrl.ResolveReference(linkUrl)
			node.UrlString = absLinkUrl.String()
		}
	}
	outChSet.NodeCh <- node
}
