package experimental

import "io"

type ReqBody io.ReadCloser

type Node struct {
	*ReqProp
	Body *ReqBody
}

type ReqProp struct {
	ParentUrl string
	UrlStr    string
}
