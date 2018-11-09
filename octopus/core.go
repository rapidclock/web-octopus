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
		o.timeToQuit = time.Duration(o.TimeToQuit)
	} else {
		log.Fatalln("TimeToQuit is not greater than 0")
	}
}

func (o *octopus) SetupSystem() {
	o.setupOctopus()

	ingestCh := make(chan *Node)
	ingestQuitCh := make(chan int, 1)
	ingestChSet := MakeNodeChSet(ingestCh, ingestQuitCh)
	ingestStrCh := make(chan string)
	inPipeChSet := &ingestPipeChSet{
		ingestCh,
		ingestStrCh,
		ingestQuitCh,
	}

	outAdapterChSet := o.OpAdapter.Consume()

	pageParseChSet := o.makeParseNodeFromHtmlPipe(ingestChSet)
	depthLimitChSet := o.makeCrawlDepthFilterPipe(pageParseChSet)
	distributorChSet := o.makeDistributorPipe(depthLimitChSet, outAdapterChSet)
	pageReqChSet := o.makePageRequisitionPipe(distributorChSet)
	invUrlFilterChSet := o.makeInvalidUrlFilterPipe(pageReqChSet)
	dupFilterChSet := o.makeDuplicateUrlFilterPipe(invUrlFilterChSet)
	protoFilterChSet := o.makeUrlProtocolFilterPipe(dupFilterChSet)
	linkAbsChSet := o.makeLinkAbsolutionPipe(protoFilterChSet)

	o.makeIngestPipe(inPipeChSet, linkAbsChSet)

	o.inpUrlStrChan = ingestStrCh
	o.masterQuitCh = make(chan int, 1)
	o.isReady = true
}

func (o *octopus) BeginCrawling(baseUrlStr string) {
	if !o.isReady {
		log.Fatal("Call BuildSystem first to setup Octopus")
	}
	go func() {
		o.inpUrlStrChan <- baseUrlStr
	}()
	for {
		select {
		case urlStr := <-o.inpUrlStrChan:
			{
				o.inpUrlStrChan <- urlStr
			}
		case <-o.masterQuitCh:
			{
				fmt.Println("Master Kill Switch Activated")
				return
			}
		}
	}
}
