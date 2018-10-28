package experimental

func MakeCompositionPipe(undupPipe chan<- *Node) chan<- *ReqProp {
	compositionPipe := make(chan *ReqProp)
	go func() {
		for {
			select {
			case req := <-compositionPipe:
				go func(reqProp *ReqProp) {
					node := &Node{
						reqProp,
						nil,
					}
					undupPipe <- node
				}(req)
			}
		}
	}()
	return compositionPipe
}
