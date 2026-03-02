# Quick Start

This guide walks you through creating and working with strategic planning documents.

## Choose Your Framework

| Framework | Best For | Example |
|-----------|----------|---------|
| **V2MOM** | Strategic planning, company vision | Annual company goals |
| **OKR** | Goal setting, measurable outcomes | Quarterly team objectives |
| **DMAIC** | Continuous improvement, Six Sigma | Process optimization metrics |

## DMAIC Example

Create a quality metrics document:

```go
package main

import (
    "fmt"
    "github.com/grokify/structured-goals/dmaic"
)

func main() {
    // Create a new DMAIC document
    doc := dmaic.New("DMAIC-2025-001", "Quality Metrics", "Jane Smith")

    // Add a category with metrics
    doc.Categories = append(doc.Categories, dmaic.Category{
        Name: "Production Quality",
        Metrics: []dmaic.Metric{
            {
                Name:           "First Pass Yield",
                Baseline:       85.0,
                Current:        94.2,
                Target:         98.0,
                Unit:           "%",
                Phase:          dmaic.PhaseControl,
                TrendDirection: dmaic.TrendHigherBetter,
            },
        },
    })

    // Calculate overall health
    health := doc.CalculateOverallHealth()
    fmt.Printf("Overall Health: %.0f%%\n", health*100)
}
```

## OKR Example

Create an objectives document:

```go
package main

import (
    "fmt"
    "github.com/grokify/structured-goals/okr"
)

func main() {
    doc := okr.New("OKR-2025-Q1", "Q1 Objectives", "John Doe")

    doc.Objectives = append(doc.Objectives, okr.Objective{
        Name: "Improve Customer Satisfaction",
        KeyResults: []okr.KeyResult{
            {
                Name:     "Increase NPS from 40 to 60",
                Baseline: 40,
                Current:  52,
                Target:   60,
            },
        },
    })

    fmt.Printf("Progress: %.0f%%\n", doc.CalculateProgress()*100)
}
```

## Loading from JSON

Load an existing document:

```go
doc, err := dmaic.ReadFile("metrics.json")
if err != nil {
    log.Fatal(err)
}

// Validate the document
errors := doc.Validate(dmaic.DefaultValidationOptions())
for _, e := range errors {
    fmt.Printf("Validation: %s\n", e.Message)
}
```

## Generating Output

### Marp Slides

```go
import "github.com/grokify/structured-goals/dmaic/render/marp"

renderer := marp.NewRenderer()
output, err := renderer.Render(doc, &render.Options{
    IncludeInitiatives: true,
    ShowControlCharts:  true,
})
```

### Pandoc PDF-Ready Markdown

```go
import "github.com/grokify/structured-goals/dmaic/render/pandoc"

renderer := pandoc.NewRenderer()
output, err := renderer.Render(doc, &render.Options{
    IncludeInitiatives: true,
})
```

## Next Steps

- [Framework Documentation](../frameworks/dmaic.md) - Detailed framework guides
- [CLI Reference](../cli/genschema.md) - Command-line tool usage
- [JSON Schema](../schema/overview.md) - Schema validation
