# Structured Goals

[![Build Status](https://github.com/grokify/structured-goals/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/grokify/structured-goals/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/grokify/structured-goals)](https://goreportcard.com/report/github.com/grokify/structured-goals)
[![Go Reference](https://pkg.go.dev/badge/github.com/grokify/structured-goals.svg)](https://pkg.go.dev/github.com/grokify/structured-goals)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/grokify/structured-goals/blob/master/LICENSE)

A Go library and CLI toolkit for managing strategic planning documents with JSON Schema validation and multiple output formats.

## Supported Frameworks

| Framework | Description | Use Case |
|-----------|-------------|----------|
| **V2MOM** | Vision, Values, Methods, Obstacles, Measures | Strategic planning (Salesforce methodology) |
| **OKR** | Objectives and Key Results | Goal setting (Google/Intel methodology) |
| **DMAIC** | Define, Measure, Analyze, Improve, Control | Continuous improvement (Six Sigma methodology) |

## Key Features

- **JSON Schema Validation** - Validate documents against comprehensive schemas
- **Multiple Renderers** - Generate Marp slides, Pandoc PDF-ready markdown
- **Statistical Process Control** - DMAIC supports UCL/LCL/CenterLine for Six Sigma
- **Flexible Structure** - V2MOM supports flat, nested, or hybrid structures
- **CLI Tools** - Command-line tools for validation and rendering

## Quick Example

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

## Installation

```bash
go get github.com/grokify/structured-goals
```

## Documentation

- [Getting Started](getting-started/installation.md) - Installation and quick start guide
- [Frameworks](frameworks/v2mom.md) - Detailed framework documentation
- [CLI Reference](cli/v2mom.md) - Command-line tool documentation
- [API Reference](api.md) - Go library API documentation

## License

MIT License - see [LICENSE](https://github.com/grokify/structured-goals/blob/master/LICENSE) for details.
