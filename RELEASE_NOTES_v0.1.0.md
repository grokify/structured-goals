# Release Notes - v0.1.0

**Release Date:** January 25, 2026

## Overview

This is the initial release of **structured-goals**, a Go library and CLI tool for managing strategic planning documents. It supports both V2MOM (Vision, Values, Methods, Obstacles, Measures) and OKR (Objectives and Key Results) frameworks with JSON Schema validation and Marp slide generation.

structured-goals integrates with [structureddocs](https://github.com/grokify/structureddocs) for consistent Marp presentation rendering across the structured-* ecosystem.

## Highlights

- **V2MOM Support** - Complete V2MOM document management with validation and rendering
- **OKR Support** - Full OKR framework with objectives, key results, and progress tracking
- **Marp Integration** - Generate presentation slides using structureddocs shared themes
- **Flexible Structure** - Support for flat, nested, and hybrid document structures
- **CLI Tools** - Command-line interface for document generation, validation, and rendering

## Installation

```bash
go install github.com/grokify/structured-goals/cmd/v2mom@v0.1.0
```

Or as a library:

```bash
go get github.com/grokify/structured-goals
```

## Features

### V2MOM Framework

The `v2mom` package provides:

- Complete V2MOM types: Vision, Values, Methods, Obstacles, Measures
- Structure modes: flat (traditional), nested (OKR-aligned), hybrid
- Terminology options: v2mom, okr, or hybrid naming
- Validation with structure enforcement and path-aware error reporting

### OKR Framework

The `okr` package provides:

- OKR types: Objectives, Key Results, Risks, Alignment
- Progress tracking with 0.0-1.0 scoring (0.7 = success threshold)
- Confidence levels (Low/Medium/High)
- Validation with configurable rules
- Score grading (A-F) and descriptions

### Marp Slide Generation

Generate professional presentation slides with:

- Three built-in themes: `default`, `corporate`, `minimal`
- Consistent styling via structureddocs integration
- Progress dashboards and metrics visualization
- Terminology-aware labels based on configuration

### CLI Commands

| Command | Description |
|---------|-------------|
| `v2mom init` | Generate template V2MOM JSON |
| `v2mom validate` | Validate V2MOM JSON with structure detection |
| `v2mom generate marp` | Generate Marp slides from V2MOM |

### CLI Flags

**init**:

- `--name` - V2MOM name (default: "My V2MOM")
- `--output`, `-o` - Output file path (default: v2mom.json)
- `--structure` - Template structure: flat, nested, hybrid (default: nested)

**validate**:

- `--structure` - Enforce specific structure mode

**generate marp**:

- `--output`, `-o` - Output file path (default: stdout)
- `--theme` - Slide theme: default, corporate, minimal
- `--terminology` - Display terminology: v2mom, okr, hybrid

## Quick Start

```bash
# Create a new V2MOM
v2mom init --name "Q1 2026 Product Strategy" -o strategy.json

# Edit the JSON file with your vision, values, methods, etc.

# Validate the document
v2mom validate strategy.json

# Generate presentation slides
v2mom generate marp strategy.json -o strategy.md

# Convert to PDF/HTML using Marp CLI
marp strategy.md -o strategy.pdf
```

## Library Usage

```go
package main

import (
    "fmt"
    "github.com/grokify/structured-goals/v2mom"
    "github.com/grokify/structured-goals/okr"
    "github.com/grokify/structured-goals/render/marp"
)

func main() {
    // Load V2MOM from file
    doc, err := v2mom.ReadFile("my-v2mom.json")
    if err != nil {
        panic(err)
    }

    // Validate
    errs := doc.Validate(v2mom.DefaultValidationOptions())
    if errs.HasErrors() {
        for _, e := range errs.Errors() {
            fmt.Printf("Error at %s: %s\n", e.Path, e.Message)
        }
    }

    // Generate Marp slides
    renderer := marp.NewRenderer()
    output, err := renderer.Render(doc, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(output))

    // Load and work with OKR documents
    okrDoc, err := okr.ReadFile("okr.json")
    if err != nil {
        panic(err)
    }
    progress := okrDoc.CalculateOverallProgress()
    fmt.Printf("Overall OKR Progress: %.0f%%\n", progress*100)
}
```

## Examples

The `examples/` directory contains sample documents:

- `product-v2mom.json` - Full-featured product strategy example
- `minimal-v2mom.json` - Minimum valid V2MOM
- `agentplexus-v2mom.json` - OKR-aligned nested structure example

## Dependencies

- Go 1.24+
- github.com/grokify/structureddocs v0.1.0
- github.com/spf13/cobra v1.10.2

## Contributors

- John Wang (@grokify)
- Claude Opus 4.5 (Co-Author)

## Links

- [GitHub Repository](https://github.com/grokify/structured-goals)
- [Go Package Documentation](https://pkg.go.dev/github.com/grokify/structured-goals)
- [Changelog](CHANGELOG.md)
- [structureddocs](https://github.com/grokify/structureddocs) - Shared rendering utilities
