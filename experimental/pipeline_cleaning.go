package experimental

import (
	"fmt"
	"net/url"
)

func MakeLinkCleaningPipe(undupPipe chan<- *Node) chan<- *Node {
	cleanPipe := make(chan *Node)
	go func() {
		for {
			select {
			case node := <-cleanPipe:
				go makeLinksAbsolute(node, undupPipe)
			}
		}
	}()
	return cleanPipe
}

func makeLinksAbsolute(node *Node, undupPipe chan<- *Node) {
	if node.ParentUrl != "" {
		linkUrl, err := url.Parse(node.UrlStr)
		if err != nil {
			return
		}
		if !linkUrl.IsAbs() {
			baseUrl, err := url.Parse(node.ParentUrl)
			if err != nil {
				return
			}
			absLinkUrl := baseUrl.ResolveReference(linkUrl)
			node.UrlStr = absLinkUrl.String()
		}
	}
	undupPipe <- node
}

func Test_makeLinksAbsolute() {
	node := &Node{
		&ReqProp{
			"https://en.wikipedia.org/wiki/Main_Page",
			"/wiki/Caijia_language",
		},
		nil,
	}
	outputChannel := make(chan *Node)
	go makeLinksAbsolute(node, outputChannel)
	newNode := <-outputChannel
	fmt.Println(newNode.UrlStr)
	fmt.Println(newNode.ParentUrl)
}
