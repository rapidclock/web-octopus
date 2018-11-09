package octopus

import (
	"fmt"
	"time"
)

func (o *octopus) makeIngestPipe(inChSet *ingestPipeChSet, opChSet *NodeChSet) {
	go channelConnector(inChSet, opChSet, o.timeToQuit, o.masterQuitCh)
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
		// case i := <-inChSet.QuitCh:
			// 	{
			// 		nodeOpChSet.QuitCh <- i
			// 		masterQuitCh <- i
			// 	}
		}
	}
}

func channelConnector(inChSet *ingestPipeChSet, opChSet *NodeChSet,
	timeOut time.Duration, masterQuitCh chan int) {
	for {
		select {
		case node := <-inChSet.NodeCh:
			opChSet.NodeCh <- node
		case i := <-inChSet.QuitCh:
			{
				fmt.Println("Quit Received on Ingest Channel")
				opChSet.QuitCh <- i
				masterQuitCh <- i
			}
		case <-time.After(timeOut * time.Second):
			{
				fmt.Println("Timeout Triggered in Ingest Channel")
				opChSet.QuitCh <- 1
				return
			}
		}
	}
}
