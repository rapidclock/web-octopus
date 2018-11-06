package octopus

import (
	"net/http"
)

func (o *octopus) makePageRequisitionPipe(outChSet *NodeChSet,
	adapterChSet *NodeChSet) *NodeChSet {
	return stdForkNodeFunc(makePageRequest, outChSet, adapterChSet)
}

func makePageRequest(node *Node, outChSet *NodeChSet, adapterChSet *NodeChSet) {
	resp, err := http.Get(node.UrlString)
	if err == nil && resp.StatusCode == 200 {
		node.Body = resp.Body
		outChSet.NodeCh <- node
		if adapterChSet != nil {
			adapterChSet.NodeCh <- node
		}
	}
}
