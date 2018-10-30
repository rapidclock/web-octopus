package experimental

import (
	"log"
	"net/http"
)

func (m *Monster) MakeUrlValidationPipe(requisitionPipe chan<- *Node) chan<- *Node {
	validationPipe := make(chan *Node)
	go func() {
		for {
			select {
			case node := <-validationPipe:
				{
					go validateLink(requisitionPipe, node)
				}
			}
		}
	}()
	return validationPipe
}

func validateLink(requisitionPipe chan<- *Node, node *Node) {
	resp, err := http.Head(node.UrlStr)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("%v\n", err)
		return
	}
	requisitionPipe <- node
}
