// Command genpandoc generates Pandoc Markdown from DMAIC JSON files.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/grokify/structured-goals/dmaic"
	"github.com/grokify/structured-goals/dmaic/render"
	"github.com/grokify/structured-goals/dmaic/render/pandoc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: genpandoc <input.json> [output.md]")
		fmt.Println("       genpandoc examples/identity-security-saas/dmaic.json")
		os.Exit(1)
	}

	inputPath := filepath.Clean(os.Args[1])
	outputPath := ""
	if len(os.Args) >= 3 {
		outputPath = filepath.Clean(os.Args[2])
	} else {
		// Default: replace .json with .md
		ext := filepath.Ext(inputPath)
		outputPath = strings.TrimSuffix(inputPath, ext) + "-report.md"
	}

	// Read DMAIC document
	doc, err := dmaic.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Create renderer and render
	renderer := pandoc.New()
	opts := render.SixSigmaOptions()

	output, err := renderer.Render(doc, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering: %v\n", err)
		os.Exit(1)
	}

	// Write output (user-controlled path is expected for CLI tool)
	if err := os.WriteFile(outputPath, output, 0600); err != nil { //nolint:gosec // G703: user-controlled path is intentional for CLI
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s\n", outputPath)
	fmt.Println("\nTo convert to PDF with LuaLaTeX:")
	fmt.Printf("  pandoc %s -o %s --pdf-engine=lualatex\n", outputPath, strings.TrimSuffix(outputPath, ".md")+".pdf")
}
