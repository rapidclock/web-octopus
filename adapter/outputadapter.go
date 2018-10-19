package adapter

import (
	"fmt"
	"io"
	"log"
	"os"

	oct "github.com/rapidclock/web-octopus/octopus"
)

// OutputAdapter is the interface for the Adapter that is used to handle
// output from the Octopus Crawler.
// The contract stipulates that the crawler provides the channel
// to listen for a quit command.
// The crawler pumps its output onto the returned channel of the Consume method.
// Implementers of the interface should listen on this channel for output from
// the crawler.
type OutputAdapter interface {
	Consume(quitCh <-chan bool) chan<- oct.CrawlOutput
}

// StdOpAdapter is an output adapter that just prints the output onto the screen.
type StdOpAdapter struct{}

func (s *StdOpAdapter) Consume(quitCh <-chan bool) chan<- oct.CrawlOutput {
	listenCh := make(chan oct.CrawlOutput)
	go func() {
		for {
			select {
			case output := <-listenCh:
				fmt.Printf("%d - %s\n", output.Depth, output.URLString)
			case <-quitCh:
				return
			}
		}
	}()
	return listenCh
}

// FileWriterAdapter is an output adapter that writes the output to a specified file.
type FileWriterAdapter struct {
	FilePath string
}

func (fw *FileWriterAdapter) Consume(quitCh <-chan bool) chan<- oct.CrawlOutput {
	listenCh := make(chan oct.CrawlOutput)
	fw.writeToFile(quitCh, listenCh)
	return listenCh
}

func (fw *FileWriterAdapter) writeToFile(quitCh <-chan bool, ch <-chan oct.CrawlOutput) {
	fp, err := fw.getFilePointer()
	if err != nil {
		fp.Close()
		log.Fatal(err)
	}
	go func() {
		defer fp.Close()
		for {
			select {
			case output := <-ch:
				fmt.Fprintf(fp, "%d - %s\n", output.Depth, output.URLString)
			case <-quitCh:
				return
			}
		}
	}()
}

func (fw *FileWriterAdapter) getFilePointer() (w io.WriteCloser, err error) {
	w, err = os.OpenFile(fw.FilePath, os.O_RDWR|os.O_CREATE, 0755)
	return
}
