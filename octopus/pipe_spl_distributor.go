package octopus

// makeDistributorPipe - Distributes any node received on its listen channel
// to the list of channels passed into this.
// Basically this behaves like a repeater or a hub.
func (o *octopus) makeDistributorPipe(outChSetList ...*NodeChSet) (
	listenChSet *NodeChSet) {
	listenCh := make(chan *Node)
	listenQuitCh := make(chan int, 1)
	listenChSet = &NodeChSet{
		NodeCh: listenCh,
		StdChannels: &StdChannels{
			QuitCh: listenQuitCh,
		},
	}
	go distribute(listenCh, listenQuitCh, outChSetList...)
	return
}

func distribute(listenCh chan *Node, listenQuitCh chan int,
	outChSetList ...*NodeChSet) {
	for {
		select {
		case node := <-listenCh:
			{
				for _, outChSet := range outChSetList {
					if outChSet != nil {
						outChSet.NodeCh <- node
					}
				}
			}
		case <-listenQuitCh:
			for _, outChSet := range outChSetList {
				if outChSet != nil {
					outChSet.QuitCh <- 1
				}
			}
			return
		}
	}
}
