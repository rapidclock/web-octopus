package octopus

func stdLinearNodeFunc(stdFunc func(n *Node, chSet *NodeChSet),
	outChSet *NodeChSet) *NodeChSet {
	listenCh := make(chan *Node)
	listenQuitCh := make(chan int, 1)
	listenChSet := &NodeChSet{
		NodeCh: listenCh,
		StdChannels: &StdChannels{
			QuitCh: listenQuitCh,
		},
	}
	go func() {
		defer close(listenCh)
		defer close(listenQuitCh)
		for {
			select {
			case node := <-listenCh:
				{
					go stdFunc(node, outChSet)
				}
			case <-listenQuitCh:
				{
					outChSet.QuitCh <- 1
					return
				}
			}
		}
	}()
	return listenChSet
}

func stdForkNodeFunc(stdFunc func(n *Node, chSetA *NodeChSet,
	chSetB *NodeChSet), outChSetA *NodeChSet, outChSetB *NodeChSet) *NodeChSet {
	listenCh := make(chan *Node)
	listenQuitCh := make(chan int, 1)
	listenChSet := &NodeChSet{
		NodeCh: listenCh,
		StdChannels: &StdChannels{
			QuitCh: listenQuitCh,
		},
	}
	go func() {
		defer close(listenCh)
		defer close(listenQuitCh)
		for {
			select {
			case node := <-listenCh:
				{
					go stdFunc(node, outChSetA, outChSetB)
				}
			case <-listenQuitCh:
				{
					if outChSetA != nil {
						outChSetA.QuitCh <- 1
					}
					if outChSetB != nil {
						outChSetB.QuitCh <- 1
					}
					return
				}
			}
		}
	}()
	return listenChSet
}
