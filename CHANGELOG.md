# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v2.1.0] - 2026-07-12

### Removed

- **darwin/amd64 (Intel) pre-built binary.** macOS releases now ship
  **arm64 only**, per the org-wide policy (darwin is Apple-Silicon only; no
  universal binaries). Intel Mac users can build from source.

### Changed

- **Linux release archives are now `.tar.gz`** (darwin/windows remain `.zip`),
  per `nlink-jp/.github` CONVENTIONS.md §Release Archive Standard.
- **`README.md` + `LICENSE` are now bundled** in every release archive
  alongside the binary (previously the archives contained only the binary).
- **darwin code-signature identifier** is now the canonical `json-filter`
  (was `json-filter-darwin-arm64`), set via `codesign -i` so it stays stable
  after the archived binary is renamed to its canonical name.
- **Dropped the `-s -w` linker strip flags**, aligning `LDFLAGS` with the
  org-standard form (`-X main.version=…`) used across the util-series. Release
  binaries now retain their symbol table / DWARF (marginally larger). This
  also fixes a Windows cross-build failure with the previous strip flags.

No change to the binary's behaviour — a packaging / build-config release.

## [v2.0.1] - 2026-05-22

### Changed

- **Releases are now Developer ID signed and Apple-notarized.**
  Darwin release zips (`json-filter-v2.0.1-darwin-{amd64,arm64}.zip`)
  carry full Apple Developer ID Application signatures and Apple
  notarization tickets. End users on macOS no longer need to
  bypass Gatekeeper with right-click → Open or
  `xattr -d com.apple.quarantine` on first launch. Local users
  who place `json-filter` under Dropbox-synced (or any other
  FileProvider-managed) paths are no longer killed by macOS's
  ad-hoc + provenance distrust policy. Pipeline:
  `scripts/codesign-darwin.sh` + `scripts/notarize-darwin.sh`,
  driven by `make package`. See `nlink-jp/.github`
  CONVENTIONS.md §"Code Signing and Notarization (macOS)" for
  the org-wide convention this implements.

## [v2.0.0] - 2026-04-13

### Changed
- Replaced regex-based JSON extraction with recursive descent parser from [nlk/jsonfix](https://github.com/nlink-jp/nlk)
- Go version requirement updated to 1.26

### Added
- 20+ JSON repair capabilities: markdown code fences, single quotes, trailing commas, unquoted keys, comments, Python literals, double-escaped JSON, and more
- Unit tests for `extractAndValidateJSON` covering both legacy and new functionality

## [v1.2.0] - 2026-03-28

### Changed
- Migrated to nlink-jp organisation — module path updated to `github.com/nlink-jp/json-filter`
- Standardised Makefile: `dist/` output, `build` / `build-all` / `package` / `test` / `clean` targets; added `linux/arm64` platform
- Updated README to follow organisation conventions (description → features → installation → usage → building)

## [v1.1.0] - 2025-10-02

### Added
- Add support for JSON arrays as a valid input format.

[Unreleased]: https://github.com/nlink-jp/json-filter/compare/v2.0.0...HEAD
[v2.0.0]: https://github.com/nlink-jp/json-filter/compare/v1.2.0...v2.0.0
[v1.2.0]: https://github.com/nlink-jp/json-filter/compare/v1.1.0...v1.2.0
[v1.1.0]: https://github.com/nlink-jp/json-filter/compare/v1.0.0...v1.1.0
[v1.0.0]: https://github.com/nlink-jp/json-filter/releases/tag/v1.0.0

## [v1.0.0] - 2025-08-28

### Added
- Initial project setup.
- `README.md`, `CHANGELOG.md`, and `LICENSE` files.
- Japanese version of `README.md` (`README.ja.md`).

### Changed
- Renamed `json-filter.go` to `main.go` for Go conventions.
- Initialized Go module (`go.mod`).
- Simplified `universal-mac` target in `Makefile` to build only macOS universal binary without building other OS binaries.
- Changed binary name from `json-filter-cli` to `json-filter`.
- Modified `release` target in `Makefile` to correctly build and package all binaries, including universal macOS binary.
- Changed release package output directory from `bin/release` to `bin`.
