package octopus

import (
	"io"
	"strings"
	"testing"
	"time"
)

func TestParseHtmlPageExtractsAnchorLinks(t *testing.T) {
	outNodeCh := make(chan *Node, 2)
	outQuitCh := make(chan int, 1)
	out := MakeNodeChSet(outNodeCh, outQuitCh)

	node := &Node{
		NodeInfo: &NodeInfo{UrlString: "https://example.com", Depth: 1},
		Body:     io.NopCloser(strings.NewReader(`<html><body><a href="/a">A</a><a href="https://other/b">B</a></body></html>`)),
	}

	parseHtmlPage(node, out)

	got := make([]string, 0, 2)
	for i := 0; i < 2; i++ {
		select {
		case n := <-outNodeCh:
			got = append(got, n.UrlString)
		case <-time.After(200 * time.Millisecond):
			t.Fatal("timed out waiting for parsed link")
		}
	}

	if got[0] != "/a" || got[1] != "https://other/b" {
		t.Fatalf("unexpected parsed links: %#v", got)
	}
}
