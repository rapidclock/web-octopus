package experimental

func (m *Monster) MakeCompositionPipe(cleanPipe chan<- *Node) chan<- *ReqProp {
	compositionPipe := make(chan *ReqProp)
	structFunc := structurize
	if m.MaxDepth > 0 {
		structFunc = structurizeWithDepth
	}
	go func() {
		for {
			select {
			case req := <-compositionPipe:
				go structFunc(m.MaxDepth, req, cleanPipe)
			}
		}
	}()
	return compositionPipe
}

func structurize(maxDepth int, reqProp *ReqProp, cleanPipe chan<- *Node) {
	node := &Node{
		reqProp,
		nil,
	}
	cleanPipe <- node
}

func structurizeWithDepth(maxDepth int, reqProp *ReqProp, cleanPipe chan<- *Node) {
	if reqProp.Depth <= maxDepth {
		structurize(maxDepth, reqProp, cleanPipe)
	}
}
