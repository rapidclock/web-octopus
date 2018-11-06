package octopus

import (
	"log"
	"time"
)

func (o *octopus) setupValidProtocolMap() {
	o.isValidProtocol = make(map[string]bool)
	for _, protocol := range o.ValidProtocols {
		o.isValidProtocol[protocol] = true
	}
}

func (o *octopus) SetupSystem() {
	o.setupValidProtocolMap()

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
	o.masterQuitCh = ingestQuitCh
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
		case <-time.After(10 * time.Second):
			{
				o.masterQuitCh <- 1
				return
			}
		}
	}
}
