package octopus

import (
	"golang.org/x/net/html"
)

func (o *octopus) makeHtmlParsingPipe(outChSet *NodeInfoChSet) *NodeChSet {
	listenCh := make(chan *Node)
	listenQuitCh := make(chan int, 1)
	listenChSet := &NodeChSet{
		NodeCh: listenCh,
		StdChannels: &StdChannels{
			QuitCh: listenQuitCh,
		},
	}
	go func() {
		defer close(listenCh)
		defer close(listenQuitCh)
		for {
			select {
			case node := <-listenCh:
				{
					go parseHtmlPage(node, outChSet)
				}
			case <-listenQuitCh:
				{
					outChSet.QuitCh <- 1
					return
				}
			}
		}
	}()
	return listenChSet
}

func parseHtmlPage(node *Node, outChSet *NodeInfoChSet) {
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
						outChSet.nodeInfoCh <- &NodeInfo{
							ParentUrlString: node.UrlString,
							UrlString:       attr.Val,
							Depth:           node.Depth + 1,
						}
					}
				}
			}
		}
	}
}
