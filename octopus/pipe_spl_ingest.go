package octopus

func (o *octopus) makeIngestPipe(inChSet *ingestPipeChSet, opChSet *NodeChSet) {
	go channelConnector(inChSet, opChSet)
	go setupStringIngestPipe(inChSet, opChSet, o.masterQuitCh)
}

func setupStringIngestPipe(inChSet *ingestPipeChSet, nodeOpChSet *NodeChSet,
	masterQuitCh chan int) {
	for {
		select {
		case str := <-inChSet.StrCh:
			{
				nodeOpChSet.NodeCh <- createNode("", str, 1)
			}
		case i := <-inChSet.QuitCh:
			{
				nodeOpChSet.QuitCh <- i
				masterQuitCh <- i
			}
		}
	}
}

func channelConnector(inChSet *ingestPipeChSet, opChSet *NodeChSet) {
	for node := range inChSet.NodeCh {
		opChSet.NodeCh <- node
	}
}
