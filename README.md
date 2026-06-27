# Gendiff

`gendiff` is a command-line utility that compares two configuration files and shows the difference between them.

The project supports JSON and YAML input files and can render the result in several output formats.

[A live demo](https://asciinema.org/a/1203308) is available on Asciinema.

## Features

- Compares flat and nested configuration files.
- Supports `.json`, `.yaml`, and `.yml` files.
- Provides three output formats: `stylish`, `plain`, and `json`.
- Can be used as a CLI tool or through the public Go function `GenDiff`.

## Pipeline Status

[![Actions Status](https://github.com/Fr0stFree/go-project-244/actions/workflows/hexlet-check.yml/badge.svg)](https://github.com/Fr0stFree/go-project-244/actions)
[![Tests and Lint](https://github.com/Fr0stFree/go-project-244/actions/workflows/test-and-lint.yml/badge.svg)](https://github.com/Fr0stFree/go-project-244/actions/workflows/test-and-lint.yml)

## Quality

[![Test Coverage](https://img.shields.io/endpoint?url=https://raw.githubusercontent.com/Fr0stFree/go-project-244/master/.github/badges/coverage-badge.json)](https://github.com/Fr0stFree/go-project-244/actions/workflows/test-and-lint.yml)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=Fr0stFree_go-project-244&metric=bugs)](https://sonarcloud.io/summary/new_code?id=Fr0stFree_go-project-244)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=Fr0stFree_go-project-244&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=Fr0stFree_go-project-244)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=Fr0stFree_go-project-244&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=Fr0stFree_go-project-244)

[![SonarQube Cloud](https://sonarcloud.io/images/project_badges/sonarcloud-light.svg)](https://sonarcloud.io/summary/new_code?id=Fr0stFree_go-project-244)

## Requirements

- Go `1.26.3`
- `make`

The linter target uses `golangci-lint` `v2.12.2`. Run `make install-lint` once before `make lint` or `make lint-fix`.

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
./bin/gendiff [options] <first-file> <second-file>
```

Arguments:

- `first-file` - path to the first configuration file.
- `second-file` - path to the second configuration file.

The command expects exactly two positional file arguments.

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

All examples below use the compiled binary and fixtures from the `internal/testdata/fixtures` directory.

### Stylish Format

`stylish` is the default output format, so the `--format` flag can be omitted.

```bash
./bin/gendiff internal/testdata/fixtures/json/file1.json internal/testdata/fixtures/json/file2.json
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
./bin/gendiff internal/testdata/fixtures/yaml/file1.yaml internal/testdata/fixtures/yaml/file2.yaml --format stylish
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
./bin/gendiff internal/testdata/fixtures/json/file3.json internal/testdata/fixtures/json/file4.json --format plain
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

`json` is useful when the diff result should be consumed by another program. It serializes the full diff tree, including unchanged nodes.

```bash
./bin/gendiff internal/testdata/fixtures/json/file1.json internal/testdata/fixtures/json/file2.json --format json
```

Output:

```json
[
  {
    "key": "follow",
    "type": "removed",
    "old": false,
    "new": null
  },
  {
    "key": "host",
    "type": "unchanged",
    "old": "hexlet.io",
    "new": "hexlet.io"
  },
  {
    "key": "proxy",
    "type": "removed",
    "old": "123.234.53.22",
    "new": null
  },
  {
    "key": "timeout",
    "type": "changed",
    "old": 50,
    "new": 20
  },
  {
    "key": "verbose",
    "type": "added",
    "old": null,
    "new": true
  }
]
```

## Makefile Commands

The project includes a `Makefile` with common development commands:

| Command | Description |
| --- | --- |
| `make build` | Builds the CLI binary into `bin/gendiff`. |
| `make test` | Runs all Go tests with verbose output. |
| `make test-coverage` | Runs tests, writes `coverage.out`, and prints coverage by function. |
| `make lint` | Checks formatting with `gofmt` and runs the installed `golangci-lint`. |
| `make lint-fix` | Formats code with `gofmt` and runs the installed `golangci-lint --fix`. |
| `make install-lint` | Installs the configured `golangci-lint` version. |

To update the coverage badge manually:

```bash
make test-coverage
.github/scripts/generate-coverage-badge.sh
```

Typical local check before pushing changes:

```bash
make install-lint
make test
make test-coverage
make lint
```

## Project Structure

```text
.
|-- cmd/gendiff        # CLI entry point
|-- .github/badges     # Generated badge data
|-- .github/scripts    # Maintenance scripts
|-- .github/workflows  # GitHub Actions workflows
|-- internal/diff      # Builds the internal diff representation
|-- internal/formatter # Renders diff in stylish, plain, and JSON formats
|-- internal/parser    # Parses JSON and YAML files
|-- internal/testdata  # Fixtures and expected outputs used by tests and examples
|-- gendiff.go         # Public package API
`-- Makefile           # Development commands
```

## Public API

The root package exposes `GenDiff`:

```go
result, err := code.GenDiff("file1.json", "file2.json", "stylish")
```

It accepts two file paths and an output format, then returns the formatted diff string.
