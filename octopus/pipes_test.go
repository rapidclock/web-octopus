package octopus

import (
	"sync"
	"testing"
	"time"
)

func TestMakeLinkAbsolute(t *testing.T) {
	nodeCh := make(chan *Node, 1)
	quitCh := make(chan int, 1)
	out := MakeNodeChSet(nodeCh, quitCh)

	node := createNode("https://example.com/guide/", "../about", 2)
	makeLinkAbsolute(node, out)

	select {
	case got := <-nodeCh:
		if got.UrlString != "https://example.com/about" {
			t.Fatalf("expected absolute url, got %q", got.UrlString)
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("timed out waiting for output node")
	}
}

func TestFilterDuplicates(t *testing.T) {
	nodeCh := make(chan *Node, 2)
	quitCh := make(chan int, 1)
	o := &octopus{visited: new(sync.Map)}
	out := MakeNodeChSet(nodeCh, quitCh)

	node := createNode("", "https://example.com", 1)
	o.filterDuplicates(node, out)
	o.filterDuplicates(node, out)

	select {
	case <-nodeCh:
	default:
		t.Fatal("expected first node to pass duplicate filter")
	}

	select {
	case <-nodeCh:
		t.Fatal("expected second node to be filtered")
	default:
	}
}

func TestConnectWithTimeout(t *testing.T) {
	listenNodeCh := make(chan *Node, 1)
	listenQuitCh := make(chan int, 1)
	outNodeCh := make(chan *Node, 1)
	outQuitCh := make(chan int, 1)

	go connectWithTimeout(listenNodeCh, listenQuitCh, MakeNodeChSet(outNodeCh, outQuitCh), 25*time.Millisecond)

	select {
	case <-outQuitCh:
	case <-time.After(200 * time.Millisecond):
		t.Fatal("expected timeout quit signal")
	}
}
