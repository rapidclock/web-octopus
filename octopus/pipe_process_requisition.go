package octopus

import (
	"net/http"
)

func (o *octopus) makePageRequisitionPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(makePageRequest, outChSet)
}

func makePageRequest(node *Node, outChSet *NodeChSet) {
	resp, err := http.Get(node.UrlString)
	if err == nil && resp.StatusCode == 200 {
		node.Body = resp.Body
		outChSet.NodeCh <- node
	}
}
