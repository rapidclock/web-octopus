# Changelog

All notable changes to this project are documented in this file.

## [1.3.0] - 2026-02-10

### Added
- Added GitHub Actions CI workflow for pull requests and default-branch pushes with module, formatting, vet, static analysis, test, and race checks.
- Added GitHub Actions publish workflow driven by GitHub Release publication to trigger Go proxy and pkg.go.dev indexing.

### Changed
- Removed legacy Travis CI configuration in favor of GitHub Actions.
- Hardened CI reliability by adding workflow concurrency cancellation, running module/static checks on the primary Go version, and pinning staticcheck.
- Hardened publish reliability by validating release tag format against `VERSION`, checking out the exact release tag, skipping prereleases, and adding retry logic for Go proxy/pkg.go.dev refresh calls.

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
