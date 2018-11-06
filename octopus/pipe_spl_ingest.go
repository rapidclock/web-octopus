package octopus

func (o *octopus) makeIngestPipe(inChSet *ingestPipeChSet, opChSet *NodeChSet) {
	go channelConnector(inChSet, opChSet)
	go setupStringIngestPipe(inChSet, opChSet)
}

func setupStringIngestPipe(inChSet *ingestPipeChSet, nodeOpChSet *NodeChSet) {
	for {
		select {
		case str := <-inChSet.StrCh:
			{
				nodeOpChSet.NodeCh <- createNode("", str, 1)
			}
		case i := <-inChSet.QuitCh:
			{
				nodeOpChSet.QuitCh <- i
			}
		}
	}
}

func channelConnector(inChSet *ingestPipeChSet, opChSet *NodeChSet) {
	for {
		select {
		case node := <-inChSet.NodeCh:
			opChSet.NodeCh <- node
		case i := <-inChSet.QuitCh:
			opChSet.QuitCh <- i
		}
	}
}
