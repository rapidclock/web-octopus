package octopus

import (
	"fmt"
	"time"
)

func (o *octopus) makeMaxDelayPipe(opNodeChSet *NodeChSet) *NodeChSet {
	listenNodeCh := make(chan *Node)
	listenQuitCh := make(chan int, 1)
	listenChSet := MakeNodeChSet(listenNodeCh, listenQuitCh)
	go connectWithTimeout(listenNodeCh, listenQuitCh, opNodeChSet, o.timeToQuit)
	return listenChSet
}

func connectWithTimeout(listenNodeCh <-chan *Node, listenQuitCh <-chan int,
	opNodeChSet *NodeChSet, timeoutDuration time.Duration) {
	timer := time.NewTimer(timeoutDuration)
	for {
		select {
		case node := <-listenNodeCh:
			opNodeChSet.NodeCh <- node
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(timeoutDuration)
		case i := <-listenQuitCh:
			opNodeChSet.QuitCh <- i
			return
		case <-timer.C:
			fmt.Println("Timeout Triggered in MaxDelayTimeout Channel")
			opNodeChSet.QuitCh <- 1
			return
		}
	}
}
