package octopus

import (
	"golang.org/x/net/html"
)

func (o *octopus) makeParseNodeFromHtmlPipe(outChSet *NodeChSet) *NodeChSet {
	return stdLinearNodeFunc(parseHtmlPage, outChSet)
}

func parseHtmlPage(node *Node, outChSet *NodeChSet) {
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
						outChSet.NodeCh <- &Node{
							NodeInfo: &NodeInfo{
								ParentUrlString: node.UrlString,
								UrlString:       attr.Val,
								Depth:           node.Depth + 1,
							},
							Body: nil,
						}
					}
				}
			}
		}
	}
}
