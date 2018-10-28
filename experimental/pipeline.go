package experimental

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
)

func BuildSystem() chan<- string {
	linkValidatePipe := LinkIngestPipeline()
	dupRemovalPipe := DuplicateRemovalPipeline(linkValidatePipe)
	processPipe := ProcessingPipeline(dupRemovalPipe)
	pageParsePipe := PageParsePipeline(processPipe)
	ingestPipe := LinkIngestPipeline(pageParsePipe)
	return linkValidatePipe
}

func LinkIngestPipeline(parsePipe chan<- io.ReadCloser) chan<- string {
	ingestPipe := make(chan string)
	go func() {
		for {
			select {
			case link := <-ingestPipe:
				go getPage(link, parsePipe)
			}
		}
	}()
	return ingestPipe
}

func getPage(link string, parsePipe chan<- io.ReadCloser) {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode == 200 {
		parsePipe <- resp.Body
	}
}

func PageParsePipeline(processPipe chan<- string) chan<- io.ReadCloser {
	parsePipeline := make(chan io.ReadCloser)
	go func() {
		for {
			reqBody := <-parsePipeline
			go parsePage(reqBody, processPipe)
		}
	}()
	return parsePipeline
}

func parsePage(body io.ReadCloser, processPipe chan<- string) {
	defer body.Close()
	z := html.NewTokenizer(body)
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
						processPipe <- attr.Val
					}
				}
			}
		}
	}
}

func ProcessingPipeline(dupRemovalPipe chan<- string) chan<- string {
	processPipe := make(chan string)
	go func() {
		for {
			select {
			case link := <-processPipe:
				return;
			}
		}
	}()
	return processPipe
}

func DuplicateRemovalPipeline() chan<- string {

}

func LinkValidationPipeline() chan<- string {

}

func StructurizePipeline() {
	go func() {


	}()
}