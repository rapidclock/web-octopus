package octopus

import (
	"net/http"
)

func (o *octopus) makeFilterUrlValidationPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(validateUrl, outChSet)
}

func validateUrl(node *Node, outChSet *NodeChSet) {
	resp, err := http.Head(node.UrlString)
	if err == nil && resp.StatusCode == 200 {
		outChSet.NodeCh <- node
	}
	// log.Printf("%v\n", err)
}
