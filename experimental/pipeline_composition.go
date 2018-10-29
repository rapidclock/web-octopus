package experimental

func MakeCompositionPipe(undupPipe chan<- *Node) chan<- *ReqProp {
	compositionPipe := make(chan *ReqProp)
	go func() {
		for {
			select {
			case req := <-compositionPipe:
				go structurize(req, undupPipe)
			}
		}
	}()
	return compositionPipe
}

func structurize(reqProp *ReqProp, undupPipe chan<- *Node) {
	node := &Node{
		reqProp,
		nil,
	}
	undupPipe <- node
}
