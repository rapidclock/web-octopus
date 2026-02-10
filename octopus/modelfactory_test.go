package octopus

import "testing"

func TestGetDefaultCrawlOptions(t *testing.T) {
	opts := GetDefaultCrawlOptions()

	if opts.MaxCrawlDepth != defaultMaxDepth {
		t.Fatalf("expected default max depth %d, got %d", defaultMaxDepth, opts.MaxCrawlDepth)
	}
	if opts.MaxCrawledUrls != defaultLinkCrawlLimit {
		t.Fatalf("expected default max crawled urls %d, got %d", defaultLinkCrawlLimit, opts.MaxCrawledUrls)
	}
	if opts.CrawlRatePerSec != defaultCrawlRateLimit || opts.CrawlBurstLimitPerSec != defaultCrawlRateLimit {
		t.Fatalf("unexpected crawl rate defaults: rate=%d burst=%d", opts.CrawlRatePerSec, opts.CrawlBurstLimitPerSec)
	}
	if opts.TimeToQuit != defaultTimeToQuit {
		t.Fatalf("expected default timeout %d, got %d", defaultTimeToQuit, opts.TimeToQuit)
	}
}

func TestMakeDefaultNodeChSet(t *testing.T) {
	chSet, nodeCh, quitCh := MakeDefaultNodeChSet()
	if chSet == nil || chSet.StdChannels == nil {
		t.Fatal("expected non-nil channel set")
	}
	if chSet.NodeCh == nil || chSet.QuitCh == nil {
		t.Fatal("expected non-nil channels on set")
	}
	if nodeCh == nil || quitCh == nil {
		t.Fatal("expected returned concrete channels to be non-nil")
	}
}
