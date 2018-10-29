package experimental

func MakeCompositionPipe(cleanPipe chan<- *Node) chan<- *ReqProp {
	compositionPipe := make(chan *ReqProp)
	go func() {
		for {
			select {
			case req := <-compositionPipe:
				go structurize(req, cleanPipe)
			}
		}
	}()
	return compositionPipe
}

func structurize(reqProp *ReqProp, cleanPipe chan<- *Node) {
	node := &Node{
		reqProp,
		nil,
	}
	cleanPipe <- node
}
