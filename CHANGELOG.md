# Changelog

All notable changes to this project are documented in this file.

## [1.4.1] - 2026-02-10

### Changed
- Updated publish automation to run **only** when a GitHub Release is published (not on every tag push).
- Added strict release-tag validation in publish workflow and kept VERSION/tag consistency checks.

## [1.4.0] - 2026-02-10

### Added
- Added GitHub Actions CI workflow for pull requests and default-branch pushes with module, formatting, vet, static analysis, test, and race checks.
- Added GitHub Actions publish workflow on version tags to trigger Go proxy and pkg.go.dev indexing.

### Changed
- Removed legacy Travis CI configuration in favor of GitHub Actions.

## [1.3.0] - 2026-02-10

### Added
- Added parser-focused unit coverage to validate anchor extraction behavior.

### Changed
- Restored robust HTML parsing using `golang.org/x/net/html` tokenizer instead of regex-based extraction.
- Restored mature rate limiting with `golang.org/x/time/rate` for predictable throttling semantics.
- Added module `replace` directives to GitHub mirrors for `golang.org/x/*` to improve fetch reliability in restricted environments.
- Fixed distributor shutdown to stop forwarding after quit signal without closing inbound channels owned by upstream producers.

## [1.1.0] - 2026-02-10

### Added
- Go module support (`go.mod`) with explicit dependency versions.
- Unit tests for crawler defaults and factory helpers.
- Unit tests for rate-limit validation and timeout setup behavior.
- Unit tests for pipeline helper behavior (link absolution, duplicate filter, timeout propagation).
- Unit test coverage for `adapter.FileWriterAdapter` write behavior.
- Comprehensive README covering installation, architecture, configuration, adapters, testing, and release policy.

### Changed
- Replaced `log.Fatal` behavior in core library paths with non-process-terminating logic (`panic` for invalid setup value and early return for nil or file-open failure paths).
- Improved file adapter open flags to use write-only, create, and truncate semantics for predictable output.

### Notes
- This version focuses on modernization and maintainability without changing the fundamental crawler pipeline design.
