package octopus

import (
	"log"
	"net/http"
)

func (o *octopus) makeInvalidUrlFilterPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(validateUrl, outChSet)
}

func validateUrl(node *Node, outChSet *NodeChSet) {
	resp, err := http.Head(node.UrlString)
	if err == nil && resp == nil {
		log.Fatal("WOW resp is nill although err is not")
	}
	if err == nil && resp != nil && resp.StatusCode == 200 {
		outChSet.NodeCh <- node
	}
	// log.Printf("%v\n", err)
}
