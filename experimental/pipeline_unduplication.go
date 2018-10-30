package experimental

import "sync"

func (m *Monster) MakeUnduplicationPipe(validationPipe chan<- *Node) chan<- *Node {
	var visitMap sync.Map
	undupPipe := make(chan *Node)
	go func() {
		for {
			select {
			case node := <-undupPipe:
				{
					if _, visited := visitMap.Load(node.UrlStr); visited {
						break
					}
					visitMap.Store(node.UrlStr, true)
					validationPipe <- node
				}
			}
		}
	}()
	return undupPipe
}
