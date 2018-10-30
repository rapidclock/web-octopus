package experimental

import "io"

type Node struct {
	*ReqProp
	Body io.ReadCloser
}

type ReqProp struct {
	ParentUrl string
	UrlStr    string
	Depth     int
}

type Options struct {
	MaxDepth int
}

type Monster struct {
	*Options
	listenPipe chan string
	compPipe   chan<- *ReqProp
}
