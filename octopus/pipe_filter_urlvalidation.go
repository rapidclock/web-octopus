package octopus

import (
	"net/http"
	"time"
)

func (o *octopus) makeInvalidUrlFilterPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(validateUrl, outChSet, "URL Validation")
}

func validateUrl(node *Node, outChSet *NodeChSet) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(defaultRequestTimeout),
	}
	resp, err := client.Head(node.UrlString)
	if err == nil && resp != nil && resp.StatusCode == 200 {
		outChSet.NodeCh <- node
	}
}
