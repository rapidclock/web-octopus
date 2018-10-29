package experimental

import "golang.org/x/net/html"

func MakeParsingPipe() (chan<- *Node, chan<- chan<- *ReqProp) {
	parsePipe := make(chan *Node)
	compPipeChan := make(chan chan<- *ReqProp)
	go func() {
		compositionPipe := <-compPipeChan
		for {
			select {
			case node := <-parsePipe:
				{
					go parsePage(node, compositionPipe)
				}
			}
		}
	}()
	return parsePipe, compPipeChan
}

func parsePage(node *Node, compositionPipe chan<- *ReqProp) {
	defer node.Body.Close()
	z := html.NewTokenizer(node.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						compositionPipe <- &ReqProp{
							ParentUrl: node.UrlStr,
							UrlStr:    attr.Val,
						}
					}
				}
			}
		}
	}
}
