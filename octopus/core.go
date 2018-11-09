package octopus

import (
	"fmt"
	"log"
	"time"
)

func (o *octopus) setupOctopus() {
	o.setupValidProtocolMap()
	o.setupTimeToQuit()
}

func (o *octopus) setupValidProtocolMap() {
	o.isValidProtocol = make(map[string]bool)
	for _, protocol := range o.ValidProtocols {
		o.isValidProtocol[protocol] = true
	}
}

func (o *octopus) setupTimeToQuit() {
	if o.TimeToQuit > 0 {
		o.timeToQuit = time.Duration(o.TimeToQuit) * time.Second
	} else {
		log.Fatalln("TimeToQuit is not greater than 0")
	}
}

func (o *octopus) SetupSystem() {
	o.isReady = false
	o.setupOctopus()

	ingestNodeCh := make(chan *Node)
	ingestQuitCh := make(chan int, 1)
	ingestStrCh := make(chan string)

	ingestChSet := MakeNodeChSet(ingestNodeCh, ingestQuitCh)
	inPipeChSet := &ingestPipeChSet{
		ingestNodeCh,
		ingestStrCh,
		ingestQuitCh,
	}

	outAdapterChSet := o.OpAdapter.Consume()

	pageParseChSet := o.makeParseNodeFromHtmlPipe(ingestChSet)
	depthLimitChSet := o.makeCrawlDepthFilterPipe(pageParseChSet)
	// maxDelayChSet := o.makeMaxDelayPipe(depthLimitChSet)
	// distributorChSet := o.makeDistributorPipe(maxDelayChSet, outAdapterChSet)
	distributorChSet := o.makeDistributorPipe(depthLimitChSet, outAdapterChSet)
	pageReqChSet := o.makePageRequisitionPipe(distributorChSet)
	invUrlFilterChSet := o.makeInvalidUrlFilterPipe(pageReqChSet)
	dupFilterChSet := o.makeDuplicateUrlFilterPipe(invUrlFilterChSet)
	protoFilterChSet := o.makeUrlProtocolFilterPipe(dupFilterChSet)
	linkAbsChSet := o.makeLinkAbsolutionPipe(protoFilterChSet)


	o.makeIngestPipe(inPipeChSet, linkAbsChSet)

	o.inputUrlStrChan = ingestStrCh
	o.masterQuitCh = make(chan int, 1)

	<-time.After(500 * time.Millisecond)
	o.isReady = true
}

func (o *octopus) BeginCrawling(baseUrlStr string) {
	if !o.isReady {
		log.Fatal("Call BuildSystem first to setup Octopus")
	}
	go func() {
		o.inputUrlStrChan <- baseUrlStr
	}()
	// for {
	// 	select {
	// 	// case urlStr := <-o.inputUrlStrChan:
	// 	// 	{
	// 	// 		o.inputUrlStrChan <- urlStr
	// 	// 	}
	// 	case <-o.masterQuitCh:
	// 		{
	// 			fmt.Println("Master Kill Switch Activated")
	// 			return
	// 		}
	// 	}
	// }
	<-o.masterQuitCh
	fmt.Println("Master Kill Switch Activated")
}

func (o *octopus) GetInputUrlStrChan() chan<- string {
	if o.isReady {
		return o.inputUrlStrChan
	} else {
		return nil
	}
}

func (o *octopus) GetMasterQuitChan() chan<- int {
	if o.isReady {
		return o.masterQuitCh
	} else {
		return nil
	}
}
