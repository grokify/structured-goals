# Pandoc Renderer

The Pandoc renderer generates PDF-ready Markdown with YAML frontmatter for [Pandoc](https://pandoc.org/), the universal document converter.

## Overview

Pandoc output is ideal for:

- Printed reports
- Audit documentation
- Compliance artifacts
- Archive records

## Features

- **Sans-serif fonts** - Professional Helvetica/Arial typography
- **2cm margins** - Balanced page layout
- **LuaLaTeX support** - High-quality PDF rendering
- **Table of contents** - Auto-generated navigation
- **Section numbering** - Hierarchical document structure

## Usage

### CLI

```bash
genpandoc -i metrics.json -o report.md
```

### Go API

```go
import (
    "github.com/grokify/structured-goals/dmaic"
    "github.com/grokify/structured-goals/dmaic/render"
    "github.com/grokify/structured-goals/dmaic/render/pandoc"
)

doc, _ := dmaic.ReadFile("metrics.json")

renderer := pandoc.NewRenderer()
output, err := renderer.Render(doc, &render.Options{
    IncludeInitiatives: true,
    IncludeRootCauses:  true,
})

os.WriteFile("report.md", output, 0644)
```

## Output Example

```markdown
---
title: "Manufacturing Quality Metrics"
author: "Maria Rodriguez"
date: "FY2025-Q1"
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

# Executive Summary

| Metric | Value |
|--------|-------|
| Overall Health | 85% |
| Categories | 3 |
| Total Metrics | 9 |
| Green Status | 6 |
| Yellow Status | 3 |

# Production Quality

## First Pass Yield

| Property | Value |
|----------|-------|
| Current | 94.2% |
| Target | 98.0% |
| Baseline | 85.0% |
| Phase | Control |
| Status | Green |

### Control Limits

| Limit | Value |
|-------|-------|
| UCL | 99.0% |
| Center Line | 95.0% |
| LCL | 90.0% |
```

## Converting to PDF

Install Pandoc and LuaLaTeX:

```bash
# macOS
brew install pandoc
brew install --cask mactex

# Ubuntu
apt install pandoc texlive-full
```

Generate PDF:

```bash
pandoc report.md -o report.pdf --pdf-engine=lualatex
```

## YAML Frontmatter Options

The renderer generates these frontmatter settings:

| Setting | Value | Purpose |
|---------|-------|---------|
| `geometry` | `margin=2cm` | Page margins |
| `fontfamily` | `helvet` | Helvetica font |
| `fontsize` | `11pt` | Body text size |
| `documentclass` | `article` | Document type |
| `toc` | `true` | Table of contents |
| `numbersections` | `true` | Section numbers |

## Customization

### Custom Margins

```bash
pandoc report.md -o report.pdf \
  --pdf-engine=lualatex \
  -V geometry:margin=1in
```

### Different Font

```bash
pandoc report.md -o report.pdf \
  --pdf-engine=lualatex \
  -V fontfamily=libertinus
```

### A4 Paper

```bash
pandoc report.md -o report.pdf \
  --pdf-engine=lualatex \
  -V papersize=a4
```

## Report Sections

The renderer generates these sections:

1. **Executive Summary** - Health metrics, status counts
2. **Categories** - One section per category
3. **Metrics** - Subsection per metric with details
4. **Control Limits** - For Control phase metrics
5. **Root Causes** - When available
6. **Initiatives** - Improvement projects

## Related

- [Renderers Overview](overview.md)
- [Marp Renderer](marp.md)
- [genpandoc CLI](../cli/genpandoc.md)
