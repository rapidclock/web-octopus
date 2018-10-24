package experimental

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
)

func GetLinks(body io.Reader) []string {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

func ProcessLinks(baseLink string, inpLinks []string) [] string {
	baseUrl, err := url.Parse(baseLink)
	if err != nil {
		log.Fatal(err)
	}
	var results []string
	for _, link := range inpLinks {
		Url, err := url.Parse(link)
		if err != nil {
			continue
		}
		Url = convertToAbsolutePath(baseUrl, Url)
		results = append(results, Url.String())
	}
	return results
}

func convertToAbsolutePath(baseUrl *url.URL, childUrl *url.URL) *url.URL {
	if !childUrl.IsAbs() {
		childUrl = baseUrl.ResolveReference(childUrl)
	}
	return childUrl
}

func RemoveDuplicates(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

func ValidateLinks(links []string) []string {
	j := 0
	for _, link := range links {
		if IsWorkingLink(link) {
			links[j] = link
			j++
		}
	}
	return links[:j]
}

func IsWorkingLink(link string) bool {
	resp, err := http.Head(link)
	if err != nil {
		return false
	}
	if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
		return false
	}
	return true
}
