# AGENTS.md — json-filter

## Project summary

CLI filter that extracts, validates, prettifies, and repairs JSON from
arbitrary text streams. Powered by nlk/jsonfix. Part of util-series.

## Build commands

```bash
make build          # Build → dist/json-filter
make test           # Run all tests
make build-all      # Cross-compile for 5 platforms
make clean          # Remove dist/
```

## Module path

`github.com/nlink-jp/json-filter`

## Key structure

```
json-filter/
├── main.go          ← single-file implementation
├── main_test.go     ← unit tests
├── go.mod
└── Makefile
```

## Key dependencies

- `github.com/nlink-jp/nlk/jsonfix` — JSON extraction and repair (recursive descent parser)

## Gotchas

- `jsonfix.Extract` returns compact JSON; prettification (2-space indent) is
  applied by `extractAndValidateJSON` via `json.Indent`.
- `--bypass` outputs original input to stdout AND error to stderr on failure,
  then exits with code 1.
- No external configuration or environment variables required.
