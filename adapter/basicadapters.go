package adapter

import (
	"fmt"
	"io"
	"log"
	"os"

	oct "github.com/rapidclock/web-octopus/octopus"
)

// StdOpAdapter is an output adapter that just prints the output onto the
// screen.
//
// Sample Output Format is:
//
//	LinkNum - Depth - Url
type StdOpAdapter struct{}

func (s *StdOpAdapter) Consume() *oct.NodeChSet {
	listenCh := make(chan *oct.Node)
	quitCh := make(chan int, 1)
	listenChSet := &oct.NodeChSet{
		NodeCh: listenCh,
		StdChannels: &oct.StdChannels{
			QuitCh: quitCh,
		},
	}
	go func() {
		i := 1
		for {
			select {
			case output := <-listenCh:
				fmt.Printf("%d - %d - %s\n", i, output.Depth, output.UrlString)
				i++
			case <-quitCh:
				return
			}
		}
	}()
	return listenChSet
}

// FileWriterAdapter is an output adapter that writes the output to a
// specified file.
// Sample Output Format is:
//
//	Depth - Url
type FileWriterAdapter struct {
	FilePath string
}

func (fw *FileWriterAdapter) Consume() *oct.NodeChSet {
	listenCh := make(chan *oct.Node)
	quitCh := make(chan int, 1)
	listenChSet := &oct.NodeChSet{
		NodeCh: listenCh,
		StdChannels: &oct.StdChannels{
			QuitCh: quitCh,
		},
	}
	fw.writeToFile(listenCh, quitCh)
	return listenChSet
}

func (fw *FileWriterAdapter) writeToFile(listenCh chan *oct.Node,
	quitCh chan int) {
	fp, err := fw.getFilePointer()
	if err != nil {
		log.Printf("failed to open output file %q: %v", fw.FilePath, err)
		// Start a fallback goroutine to drain listenCh so producers do not block.
		go func() {
			for {
				select {
				case <-listenCh:
					// Discard messages; file is not available.
				case <-quitCh:
					return
				}
			}
		}()
		return
	}
	go func() {
		defer fp.Close()
		for {
			select {
			case output := <-listenCh:
				_, err = fmt.Fprintf(fp, "%d - %s\n", output.Depth,
					output.UrlString)
				if err != nil {
					log.Println("File Error - ", err)
				}
			case <-quitCh:
				return
			}
		}
	}()
}

func (fw *FileWriterAdapter) getFilePointer() (w io.WriteCloser, err error) {
	w, err = os.OpenFile(fw.FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	return
}
