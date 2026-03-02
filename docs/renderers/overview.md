# Renderers Overview

Structured Goals provides multiple renderers for converting JSON documents into presentation and documentation formats.

## Available Renderers

| Renderer | Output Format | Use Case |
|----------|---------------|----------|
| **Marp** | Markdown slides | Presentations, stakeholder reviews |
| **Pandoc** | PDF-ready Markdown | Reports, documentation, printing |

## Renderer Interface

All DMAIC renderers implement a common interface:

```go
type Renderer interface {
    Format() string
    FileExtension() string
    Render(doc *dmaic.DMAICDocument, opts *Options) ([]byte, error)
}
```

## Render Options

Common options available for all renderers:

```go
type Options struct {
    // Content inclusion
    IncludeDataPoints      bool
    IncludeRootCauses      bool
    IncludeInitiatives     bool

    // Display options
    ShowControlCharts      bool
    ShowCapabilityMetrics  bool
    GroupByPhase           bool
    GroupByStatus          bool

    // Limits
    MaxDataPoints          int

    // Styling
    Theme                  string
    CustomCSS              string

    // Metadata
    Metadata               map[string]string
}
```

## Usage Example

```go
import (
    "github.com/grokify/structured-goals/dmaic"
    "github.com/grokify/structured-goals/dmaic/render"
    "github.com/grokify/structured-goals/dmaic/render/marp"
    "github.com/grokify/structured-goals/dmaic/render/pandoc"
)

// Load document
doc, err := dmaic.ReadFile("metrics.json")
if err != nil {
    log.Fatal(err)
}

// Render options
opts := &render.Options{
    IncludeInitiatives: true,
    ShowControlCharts:  true,
}

// Marp slides
marpRenderer := marp.NewRenderer()
slides, err := marpRenderer.Render(doc, opts)

// Pandoc PDF-ready
pandocRenderer := pandoc.NewRenderer()
report, err := pandocRenderer.Render(doc, opts)
```

## Output Workflow

### Marp Slides

```bash
# Generate Marp markdown
genpandoc --format marp -i metrics.json -o slides.md

# Convert to PDF with Marp CLI
marp slides.md --pdf
```

### Pandoc PDF

```bash
# Generate Pandoc markdown
genpandoc -i metrics.json -o report.md

# Convert to PDF with Pandoc
pandoc report.md -o report.pdf --pdf-engine=lualatex
```

## Choosing a Renderer

| Need | Recommended Renderer |
|------|---------------------|
| Executive presentation | Marp |
| Printed report | Pandoc |
| Stakeholder review | Marp |
| Audit documentation | Pandoc |
| Team meeting slides | Marp |
| Archive/compliance | Pandoc |
