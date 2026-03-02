# Marp Renderer

The Marp renderer generates Markdown slides compatible with [Marp](https://marp.app/), the Markdown Presentation Ecosystem.

## Overview

Marp slides are ideal for:

- Executive presentations
- Stakeholder reviews
- Team meetings
- Progress updates

## Slide Structure

The Marp renderer generates the following slide sequence:

1. **Title Slide** - Document name, owner, period
2. **Executive Summary** - Overall health, phase distribution
3. **Category Overview** - One slide per category with metrics summary
4. **Metric Details** - Detailed slides for Yellow/Red status metrics
5. **DMAIC Phase Summary** - Metrics grouped by phase
6. **Control Charts** - SPC visualizations (optional)
7. **Root Cause Analysis** - Analysis findings (optional)
8. **Initiatives Summary** - Improvement projects (optional)
9. **Questions** - Closing slide

## Usage

### Go API

```go
import (
    "github.com/grokify/structured-goals/dmaic"
    "github.com/grokify/structured-goals/dmaic/render"
    "github.com/grokify/structured-goals/dmaic/render/marp"
)

doc, _ := dmaic.ReadFile("metrics.json")

renderer := marp.NewRenderer()
output, err := renderer.Render(doc, &render.Options{
    IncludeInitiatives:    true,
    ShowControlCharts:     true,
    ShowCapabilityMetrics: true,
})

os.WriteFile("slides.md", output, 0644)
```

## Output Example

```markdown
---
marp: true
theme: default
paginate: true
---

# Manufacturing Quality Metrics

**Owner:** Maria Rodriguez
**Period:** FY2025-Q1
**Status:** Active

---

## Executive Summary

| Metric | Status |
|--------|--------|
| Overall Health | 85% |
| Green Metrics | 6 |
| Yellow Metrics | 3 |
| Red Metrics | 0 |

---

## Production Quality

| Metric | Current | Target | Status |
|--------|---------|--------|--------|
| First Pass Yield | 94.2% | 98.0% | 🟢 |
| DPMO | 3,400 | 2,000 | 🟢 |
| Scrap Rate | 1.8% | 1.0% | 🟡 |
```

## Converting to PDF

Install Marp CLI:

```bash
npm install -g @marp-team/marp-cli
```

Generate PDF:

```bash
marp slides.md --pdf
```

Generate HTML:

```bash
marp slides.md --html
```

Generate PowerPoint:

```bash
marp slides.md --pptx
```

## Themes

Marp supports several built-in themes:

| Theme | Description |
|-------|-------------|
| `default` | Clean, minimal design |
| `gaia` | Yellow accent, modern |
| `uncover` | Dark background |

Custom themes can be specified in options:

```go
opts := &render.Options{
    Theme: "gaia",
}
```

## Status Icons

The renderer uses emoji icons for status:

| Status | Icon |
|--------|------|
| Green | 🟢 |
| Yellow | 🟡 |
| Red | 🔴 |

## Related

- [Renderers Overview](overview.md)
- [Pandoc Renderer](pandoc.md)
- [DMAIC Framework](../frameworks/dmaic.md)
