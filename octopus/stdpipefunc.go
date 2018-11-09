package octopus

import (
	"fmt"
)

type stdFunc func(*Node, *NodeChSet)

func stdLinearNodeFunc(stdFn stdFunc, outChSet *NodeChSet) *NodeChSet {
	listenCh := make(chan *Node)
	listenQuitCh := make(chan int, 1)
	listenChSet := MakeNodeChSet(listenCh, listenQuitCh)
	go func() {
		for {
			select {
			case node := <-listenCh:
				{
					go stdFn(node, outChSet)
				}
			case <-listenQuitCh:
				{
					fmt.Println("Quit Received on Internal Channel")
					outChSet.QuitCh <- 1
					return
				}
			}
		}
	}()
	return listenChSet
}
