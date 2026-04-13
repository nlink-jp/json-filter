# json-filter

A command-line filter that extracts, validates, prettifies, and repairs JSON from arbitrary text streams — ideal for processing LLM outputs, API responses, and log data.

## Features

- **JSON extraction**: Identifies and extracts JSON objects or arrays embedded in larger text streams using a recursive descent parser ([nlk/jsonfix](https://github.com/nlink-jp/nlk))
- **Automatic prettification**: Formats valid JSON with 2-space indentation
- **Robust JSON repair**: Fixes 20+ common issues found in LLM output, API responses, and log data:
  - Markdown code fences (`` ```json ... ``` ``)
  - Single-quoted keys and values
  - Trailing commas
  - Unquoted keys
  - Missing closing braces/brackets (including deeply nested)
  - Comments (`//`, `/* */`, `#`)
  - Python-style literals (`True`, `False`, `None`)
  - Double-escaped JSON (`{\"key\": \"value\"}`)
  - And more — see [nlk/jsonfix documentation](https://github.com/nlink-jp/nlk) for the full list
- **Bypass mode**: `--bypass` passes the original input through to stdout if extraction fails, preventing pipeline interruptions

## Installation

Download the latest binary for your platform from the [releases page](https://github.com/nlink-jp/json-filter/releases).

Extract and place the binary in your `$PATH`:

```sh
unzip json-filter-<version>-<os>-<arch>.zip
mv json-filter /usr/local/bin/
```

## Usage

`json-filter` reads from stdin and writes processed JSON to stdout.

```sh
<command> | json-filter [flags]
```

### Flags

| Flag | Description |
|------|-------------|
| `--bypass` | Pass original input through if JSON extraction fails |
| `--version` | Print version information and exit |

### Examples

**Extract and prettify JSON from log output:**

```sh
echo 'INFO: data: {"id": 1, "name": "Alice"}' | json-filter
# {
#   "id": 1,
#   "name": "Alice"
# }
```

**Repair incomplete JSON:**

```sh
echo '{"data": {"item": "value"' | json-filter
# {
#   "data": {
#     "item": "value"
#   }
# }
```

**Extract JSON from markdown code fences:**

```sh
printf '```json\n{"key": "value"}\n```\n' | json-filter
# {
#   "key": "value"
# }
```

**Fix single-quoted keys, trailing commas, and unquoted keys:**

```sh
echo "{'name': 'Alice', 'age': 30,}" | json-filter
# {
#   "name": "Alice",
#   "age": 30
# }
```

**Use with curl:**

```sh
curl -s https://api.github.com/users/octocat | json-filter
```

**Use --bypass to avoid breaking a pipeline:**

```sh
some-command | json-filter --bypass | next-command
```

## Building

Requires Go 1.26 or later.

```sh
git clone https://github.com/nlink-jp/json-filter.git
cd json-filter
make build        # Build for the current platform → dist/json-filter
make build-all    # Cross-compile for all platforms → dist/json-filter-<os>-<arch>
make package      # Build and create .zip archives → dist/json-filter-<version>-<os>-<arch>.zip
make test         # Run the test suite
make clean        # Remove dist/
```

Target platforms: `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`.

## See Also

- [README.ja.md](README.ja.md) — 日本語ドキュメント
- [CHANGELOG.md](CHANGELOG.md)
