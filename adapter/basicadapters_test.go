package adapter

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	oct "github.com/rapidclock/web-octopus/octopus"
)

func TestFileWriterAdapterWritesOutput(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "crawl.log")

	adapter := &FileWriterAdapter{FilePath: filePath}
	chSet := adapter.Consume()
	chSet.NodeCh <- &oct.Node{NodeInfo: &oct.NodeInfo{Depth: 2, UrlString: "https://example.com"}}
	chSet.QuitCh <- 1

	deadline := time.Now().Add(500 * time.Millisecond)
	for {
		data, err := os.ReadFile(filePath)
		if err == nil && len(data) > 0 {
			if got := string(data); got != "2 - https://example.com\n" {
				t.Fatalf("unexpected file content: %q", got)
			}
			break
		}
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for file content: %v", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
