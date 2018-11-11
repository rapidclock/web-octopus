/*
Package octopus implements a concurrent web crawler.
The octopus uses a pipeline of channels to implement a non-blocking web crawler.
The octopus also provides user configurable options that can be used to
customize the behaviour of the crawler.

Features

Current Features of the crawler include:
	1. Depth Limited Crawling
	2. User specified valid protocols
	3. User buildable adapters that the crawler feeds output to.
	4. Filter Duplicates.
	5. Filter URLs that fail a HEAD request.
	6. User specifiable max timeout between two successive url requests.
 */
package octopus
