/*
Package octopus implements a concurrent web crawler.
The octopus uses a pipeline of channels to implement a non-blocking web crawler.
The octopus also provides user configurable options that can be used to
customize the behaviour of the crawler.

# Features

Current Features of the crawler include:
 1. User specifiable Depth Limited Crawling
 2. User specified valid protocols
 3. User buildable adapters that the crawler feeds output to.
 4. Filter Duplicates.
 5. Filter URLs that fail a HEAD request.
 6. User specifiable max timeout between two successive url requests.
 7. User specifiable Max Number of Links to be crawled.

# Pipeline Overview

The overview of the Pipeline is given below:
 1. Ingest
 2. Link Absolution
 3. Protocol Filter
 4. Duplicate Filter
 5. Invalid Url Filter (Urls whose HEAD request Fails)
    (5x) (Optional) Crawl Rate Limiter.
    [6]. Make GET Request
    7a. Send to Output Adapter
    7b. Check for Timeout (gap between two output on this channel).
 8. Max Links Crawled Limit Filter
 9. Depth Limit Filter
 10. Parse Page for more URLs.

Note: The output from 7b. is fed to 8.

	1 -> 2 -> 3 -> 4 -> 5 -> (5x) -> [6] -> 7b -> 8 -> 9 -> 10 -> 1
*/
package octopus
