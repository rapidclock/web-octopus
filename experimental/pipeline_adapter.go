package experimental

import "fmt"

func GetOutputAdapterPipe() chan<- *Node {
	opPipe := make(chan *Node)
	go HandleOutput(opPipe)
	return opPipe
}

func HandleOutput(opAdapterPipe <-chan *Node) {
	count := 0
	for {
		opNode := <-opAdapterPipe
		count++
		fmt.Printf("%d - %d - %v\n", count, opNode.Depth, opNode.UrlStr)
	}
}
