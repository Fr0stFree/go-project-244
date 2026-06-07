# Gendiff

[![Actions Status](https://github.com/Fr0stFree/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Fr0stFree/go-project-244/actions)
[![Tests and Lint](https://github.com/Fr0stFree/go-project-244/actions/workflows/test-and-lint.yml/badge.svg)](https://github.com/Fr0stFree/go-project-244/actions/workflows/test-and-lint.yml)

`gendiff` is a command-line utility that compares two configuration files and shows the difference between them.

The project supports JSON and YAML input files and can render the result in several output formats.

## Features

- Compares flat and nested configuration files.
- Supports `.json`, `.yaml`, and `.yml` files.
- Provides three output formats: `stylish`, `plain`, and `json`.
- Can be used as a CLI tool or through the public Go function `GenDiff`.

## Requirements

- Go `1.26.3`
- `make`

The linter target uses `golangci-lint` `v2.12.2`. If it is not installed, `make lint` and `make lint-fix` install it automatically into your Go binary directory.

## Installation

Clone the repository and build the binary:

```bash
git clone https://github.com/Fr0stFree/go-project-244.git
cd go-project-244
make build
```

The compiled program will be available at:

```bash
./bin/gendiff
```

## Usage

```bash
gendiff [global options] [arguments...]
```

Arguments:

- `file1` - path to the first configuration file.
- `file2` - path to the second configuration file.

Options:

| Flag | Description | Default |
| --- | --- | --- |
| `-f`, `--format` | Output format: `stylish`, `plain`, or `json` | `stylish` |
| `-h`, `--help` | Show help |  |

You can check the CLI help with:

```bash
./bin/gendiff --help
```

## Examples

All examples below use the compiled binary and fixtures from the `testdata` directory.

### Stylish Format

`stylish` is the default output format, so the `--format` flag can be omitted.

```bash
./bin/gendiff testdata/fixtures/json/file1.json testdata/fixtures/json/file2.json
```

Output:

```text
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
```

The same format works for YAML files:

```bash
./bin/gendiff testdata/fixtures/yaml/file1.yaml testdata/fixtures/yaml/file2.yaml --format stylish
```

Output:

```text
{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}
```

### Plain Format

`plain` is useful when you want a human-readable description of changed properties.

```bash
./bin/gendiff testdata/fixtures/json/file3.json testdata/fixtures/json/file4.json --format plain
```

Output:

```text
Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]
```

### JSON Format

`json` is useful when the diff result should be consumed by another program.

```bash
./bin/gendiff testdata/fixtures/json/file1.json testdata/fixtures/json/file2.json --format json
```

Output:

```json
{
  "follow": {
    "type": "removed",
    "value": false
  },
  "proxy": {
    "type": "removed",
    "value": "123.234.53.22"
  },
  "timeout": {
    "new": 20,
    "old": 50,
    "type": "changed"
  },
  "verbose": {
    "type": "added",
    "value": true
  }
}
```

## Makefile Commands

The project includes a `Makefile` with common development commands:

| Command | Description |
| --- | --- |
| `make build` | Builds the CLI binary into `bin/gendiff`. |
| `make test` | Runs all Go tests with verbose output. |
| `make lint` | Checks formatting with `gofmt` and runs `golangci-lint`. |
| `make lint-fix` | Formats code with `gofmt` and runs `golangci-lint --fix`. |
| `make install-lint` | Installs `golangci-lint` if it is missing. |

Typical local check before pushing changes:

```bash
make test
make lint
```
