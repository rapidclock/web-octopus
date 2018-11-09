package octopus

import (
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
		case i := <-inChSet.QuitCh:
			{
				nodeOpChSet.QuitCh <- i
				masterQuitCh <- i
			}
		}
	}
}

func channelConnector(inChSet *ingestPipeChSet, opChSet *NodeChSet,
	timeOut time.Duration, masterQuitCh chan int) {
	// timeOutTimer := time.NewTimer(timeOut)
	for {
		// timeOutCh = time.After(timeOut * time.Second)
		// timeOutCh = time.NewTimer(timeOut)
		select {
		case node := <-inChSet.NodeCh:
			opChSet.NodeCh <- node
			// if !timeOutTimer.Stop() {
			// 	<-timeOutTimer.C
			// }
			// log.Println("abc")
			// timeOutTimer.Reset(timeOut)
			// case i := <-inChSet.QuitCh:
			// 	{
			// 		fmt.Println("Quit Received on Ingest Channel")
			// opChSet.QuitCh <- i
			// 		masterQuitCh <- i
			// 		if !timeOutTimer.Stop() {
			// 			<-timeOutTimer.C
			// 		}
			// 		timeOutTimer.Reset(timeOut)
			// 	}
			// case <-timeOutTimer.C:
			// 	fmt.Println("Timeout Triggered in Ingest Channel")
			// 	opChSet.QuitCh <- 1
			// 	masterQuitCh <- 1
			// 	return

		}
	}
}
