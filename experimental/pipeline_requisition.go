package experimental

import "net/http"

func MakeRequistionPipe(parsePipe chan<- *Node, opAdapterPipe chan<- *Node) chan<- *Node {
	requisitionPipe := make(chan *Node)
	go func() {
		for {
			select {
			case node := <-requisitionPipe:
				{
					go makeRequest(node, parsePipe, opAdapterPipe)
				}
			}
		}
	}()
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
