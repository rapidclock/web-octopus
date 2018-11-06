package octopus

type stdFunc func(*Node, *NodeChSet)

func stdLinearNodeFunc(stdFn stdFunc, outChSet *NodeChSet) *NodeChSet {
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
					go stdFn(node, outChSet)
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
