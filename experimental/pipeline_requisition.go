package experimental

import "net/http"

func (m *Monster) MakeRequisitionPipe(parsePipe chan<- *Node, opAdapterPipe chan<- *Node) chan<- *Node {
	requisitionPipe := make(chan *Node)
	isDepthLimited := false
	if m.MaxDepth > 0 {
		isDepthLimited = true
	}
	go func(isDepthLimited bool, maxDepth int) {
		for {
			select {
			case node := <-requisitionPipe:
				{
					go depthLimitedRequest(isDepthLimited, maxDepth, node, parsePipe, opAdapterPipe)
				}
			}
		}
	}(isDepthLimited, m.MaxDepth)
	return requisitionPipe
}

func makeRequest(node *Node, parsePipe chan<- *Node, opAdapterPipe chan<- *Node) {
	resp, err := http.Get(node.UrlStr)
	if err == nil && resp.StatusCode == 200 {
		node.Body = resp.Body
		parsePipe <- node
		if opAdapterPipe != nil {
			opAdapterPipe <- node
		}
	}
}

func depthLimitedRequest(isDepthLimited bool, maxDepth int,
	node *Node, parsePipe chan<- *Node, opAdapterPipe chan<- *Node) {
	if !isDepthLimited || node.Depth <= maxDepth {
		makeRequest(node, parsePipe, opAdapterPipe)
	}
}
