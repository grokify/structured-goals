# genpandoc CLI

Generate Pandoc-ready Markdown from DMAIC documents.

## Installation

```bash
go install github.com/grokify/structured-goals/cmd/genpandoc@latest
```

## Usage

```bash
genpandoc -i input.json -o output.md [options]
```

## Options

| Flag | Description | Default |
|------|-------------|---------|
| `-i, --input` | Input JSON file (required) | - |
| `-o, --output` | Output Markdown file (required) | - |
| `--initiatives` | Include initiatives section | `true` |
| `--root-causes` | Include root causes | `true` |
| `--control-limits` | Show control limits | `true` |

## Examples

### Basic Usage

```bash
genpandoc -i metrics.json -o report.md
```

### Minimal Report

```bash
genpandoc -i metrics.json -o report.md \
  --initiatives=false \
  --root-causes=false
```

### Generate PDF

```bash
genpandoc -i metrics.json -o report.md
pandoc report.md -o report.pdf --pdf-engine=lualatex
```

## Output Format

The generated Markdown includes YAML frontmatter:

```yaml
---
title: "Document Title"
author: "Owner Name"
date: "Period"
geometry:
  - margin=2cm
fontfamily: helvet
fontsize: 11pt
documentclass: article
header-includes:
  - \renewcommand{\familydefault}{\sfdefault}
  - \usepackage{booktabs}
  - \usepackage{longtable}
toc: true
numbersections: true
---
```

## Pandoc Conversion

### PDF with LuaLaTeX

```bash
pandoc report.md -o report.pdf --pdf-engine=lualatex
```

### DOCX

```bash
pandoc report.md -o report.docx
```

### HTML

```bash
pandoc report.md -o report.html --standalone
```

## Requirements

For PDF generation:

- [Pandoc](https://pandoc.org/installing.html)
- [TeX Live](https://www.latex-project.org/get/) or [MiKTeX](https://miktex.org/)

Install on macOS:

```bash
brew install pandoc
brew install --cask mactex
```

Install on Ubuntu:

```bash
sudo apt install pandoc texlive-full
```

## Related

- [Pandoc Renderer](../renderers/pandoc.md)
- [DMAIC Framework](../frameworks/dmaic.md)
