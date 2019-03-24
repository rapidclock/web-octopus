package octopus

import (
	"net/http"
	"time"
)

func (o *octopus) makePageRequisitionPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(makePageRequest, outChSet, "URL Requisition")
}

func makePageRequest(node *Node, outChSet *NodeChSet) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(defaultRequestTimeout),
	}
	resp, err := client.Get(node.UrlString)
	if err == nil && resp.StatusCode == 200 {
		node.Body = resp.Body
		outChSet.NodeCh <- node
	}
}
