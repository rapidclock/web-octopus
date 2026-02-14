# web-octopus
[![Go Reference](https://pkg.go.dev/badge/github.com/rapidclock/web-octopus.svg)](https://pkg.go.dev/github.com/rapidclock/web-octopus@v1.3.0)


A concurrent, channel-pipeline web crawler in Go.

> This release modernizes the project for current Go module workflows, testing expectations, and maintainability standards.

## Table of contents

- [Highlights](#highlights)
- [Installation](#installation)
- [Quick start](#quick-start)
- [Architecture](#architecture)
- [Configuration reference](#configuration-reference)
- [Output adapters](#output-adapters)
- [Testing](#testing)
- [Versioning and release](#versioning-and-release)
- [Compatibility notes](#compatibility-notes)

## Highlights

- Uses Go modules (`go.mod`) instead of legacy `go get`-only workflow.
- Includes automated unit tests for crawler defaults, validation behavior, pipeline helpers, and adapter output.
- Improved adapter safety around file handling and error paths.
- Expanded docs with architecture details and operational guidance.

## Installation

```bash
go get github.com/rapidclock/web-octopus@v1.3.0
```

Import packages:

```go
import (
    "github.com/rapidclock/web-octopus/adapter"
    "github.com/rapidclock/web-octopus/octopus"
)
```

## Quick start

```go
package main

import (
	"github.com/rapidclock/web-octopus/adapter"
	"github.com/rapidclock/web-octopus/octopus"
)

func main() {
	opAdapter := &adapter.StdOpAdapter{}

	options := octopus.GetDefaultCrawlOptions()
	options.MaxCrawlDepth = 3
	options.TimeToQuit = 10
	options.CrawlRatePerSec = 5
	options.CrawlBurstLimitPerSec = 8
	options.OpAdapter = opAdapter

	crawler := octopus.New(options)
	crawler.SetupSystem()
	crawler.BeginCrawling("https://www.example.com")
}
```

## Architecture

`web-octopus` uses a staged channel pipeline. Nodes (URLs + metadata) flow through filter and processing stages:

1. Ingest
2. Link absolution
3. Protocol filter
4. Duplicate filter
5. URL validation (`HEAD`)
6. Optional rate limiter
7. Page requisition (`GET`)
8. Distributor
   - Output adapter stream
   - Max delay watchdog stream
9. Max crawled links limiter (optional)
10. Crawl depth filter
11. HTML parsing back into ingest

This design allows localized extension by replacing adapters and modifying options, while preserving high concurrency.

## Configuration reference

`CrawlOptions` controls crawler behavior:

- `MaxCrawlDepth int64` — max depth for crawled nodes.
- `MaxCrawledUrls int64` — max total unique URLs; `-1` means unlimited.
- `CrawlRatePerSec int64` — request rate limit, negative to disable.
- `CrawlBurstLimitPerSec int64` — burst capacity for rate limiting.
- `IncludeBody bool` — include body in crawled node (currently internal pipeline behavior).
- `OpAdapter OutputAdapter` — required output sink.
- `ValidProtocols []string` — accepted URL schemes (e.g., `http`, `https`).
- `TimeToQuit int64` — max idle seconds before automatic quit.

### Defaults

Use:

```go
opts := octopus.GetDefaultCrawlOptions()
```

Default values are tuned for local experimentation:

- Depth: `2`
- Max links: `-1` (unbounded)
- Rate limit: disabled
- Protocols: `http`, `https`
- Timeout gap: `30s`

## Output adapters

The crawler emits processed nodes through the `OutputAdapter` interface:

```go
type OutputAdapter interface {
    Consume() *NodeChSet
}
```

### Built-in adapters

1. `adapter.StdOpAdapter`
   - Prints `count - depth - URL` to stdout.
2. `adapter.FileWriterAdapter`
   - Writes `depth - URL` lines to a file.

### Writing a custom adapter

Create channels, return `*octopus.NodeChSet`, and consume nodes in a goroutine. Always handle quit signals to avoid goroutine leaks.

## Testing

Run the full test suite:

```bash
go test ./...
```

Recommended local checks before release:

```bash
go test ./... -race
go vet ./...
```


## CI/CD

This repository uses GitHub Actions (not Travis CI):

- **CI workflow** (`.github/workflows/ci.yml`) runs automatically on PR open/sync/reopen and on pushes to the default branch. It validates module tidiness, formatting, vet/staticcheck, and test suites (including race detection).
- **Publish workflow** (`.github/workflows/publish.yml`) runs only when a GitHub **Release** is published (excluding prereleases), validates tag/version alignment, and triggers indexing on both the Go proxy and pkg.go.dev so new versions are discoverable quickly.

Release flow:

1. Update `VERSION` and `CHANGELOG.md`.
2. Merge to default branch.
3. Create and push tag `vX.Y.Z` matching `VERSION`.
4. Publish a GitHub Release for that tag.
5. GitHub Actions publish workflow handles Go portal refresh calls.

## Versioning and release

- Project follows semantic versioning.
- Current release in this repository: **v1.3.0**.
- See `CHANGELOG.md` for release notes.

## Compatibility notes

- Legacy examples using old `go get` package paths still map to the same module path.
- Existing adapters remain source-compatible.
