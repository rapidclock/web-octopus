package octopus

import (
	"fmt"
	"time"
)

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

	o.inputUrlStrChan = ingestStrCh
	o.masterQuitCh = make(chan int, 1)

	outAdapterChSet := o.OpAdapter.Consume()

	pageParseChSet := o.makeParseNodeFromHtmlPipe(ingestChSet)
	depthLimitChSet := o.makeCrawlDepthFilterPipe(pageParseChSet)
	maxDelayChSet := o.makeMaxDelayPipe(depthLimitChSet)

	distributorChSet := o.handleDistributorPipeline(maxDelayChSet, outAdapterChSet)

	pageReqChSet := o.makePageRequisitionPipe(distributorChSet)

	invUrlFilterChSet := o.handleRateLimitingPipeline(pageReqChSet)

	dupFilterChSet := o.makeDuplicateUrlFilterPipe(invUrlFilterChSet)
	protoFilterChSet := o.makeUrlProtocolFilterPipe(dupFilterChSet)
	linkAbsChSet := o.makeLinkAbsolutionPipe(protoFilterChSet)

	o.makeIngestPipe(inPipeChSet, linkAbsChSet)

	<-time.After(500 * time.Millisecond)
	o.isReady = true
}

func (o *octopus) handleDistributorPipeline(maxDelayChSet, outAdapterChSet *NodeChSet) *NodeChSet {
	var distributorChSet *NodeChSet
	if o.MaxCrawledUrls < 0 {
		distributorChSet = o.makeDistributorPipe(maxDelayChSet, outAdapterChSet)
	} else {
		maxLinksCrawledChSet := o.makeCrawlLinkCountLimitPipe(outAdapterChSet)
		distributorChSet = o.makeDistributorPipe(maxDelayChSet, maxLinksCrawledChSet)
	}
	return distributorChSet
}

func (o *octopus) handleRateLimitingPipeline(pageReqChSet *NodeChSet) *NodeChSet {
	var invUrlFilterChSet *NodeChSet
	if o.rateLimiter != nil {
		rateLimitingChSet := o.makeRateLimitingPipe(pageReqChSet)
		invUrlFilterChSet = o.makeInvalidUrlFilterPipe(rateLimitingChSet)
	} else {
		invUrlFilterChSet = o.makeInvalidUrlFilterPipe(pageReqChSet)
	}
	return invUrlFilterChSet
}

func (o *octopus) BeginCrawling(baseUrlStr string) {
	if !o.isReady {
		panic("Call BuildSystem first to setup Octopus")
	}
	go func() {
		o.inputUrlStrChan <- baseUrlStr
	}()
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
