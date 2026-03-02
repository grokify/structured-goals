# v2mom CLI

Command-line tool for working with V2MOM documents.

## Installation

```bash
go install github.com/grokify/structured-goals/cmd/v2mom@latest
```

## Commands

### validate

Validate a V2MOM document against the JSON Schema:

```bash
v2mom validate document.json
```

Options:

| Flag | Description |
|------|-------------|
| `--strict` | Use strict validation rules |
| `--quiet` | Only output errors |

### render

Render a V2MOM document to various formats:

```bash
v2mom render document.json -o output.md
```

Options:

| Flag | Description |
|------|-------------|
| `-o, --output` | Output file path |
| `-f, --format` | Output format (marp, pandoc) |

## Examples

### Validate a Document

```bash
v2mom validate strategy.json
# ✓ Document is valid
```

### Strict Validation

```bash
v2mom validate strategy.json --strict
# ✓ Document is valid (strict mode)
```

### Generate Marp Slides

```bash
v2mom render strategy.json -f marp -o slides.md
```

### Generate PDF Report

```bash
v2mom render strategy.json -f pandoc -o report.md
pandoc report.md -o report.pdf --pdf-engine=lualatex
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Validation errors |
| 2 | File not found |
| 3 | Invalid JSON |
