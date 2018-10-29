package experimental

import "io"

type Node struct {
	*ReqProp
	Body io.ReadCloser
}

type ReqProp struct {
	ParentUrl string
	UrlStr    string
}

type Monster struct {
	listenPipe chan string
	compPipe   chan<- *ReqProp
}
