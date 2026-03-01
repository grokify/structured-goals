// Command genschema generates JSON Schema files from Go types.
package main

import (
	"fmt"
	"os"

	"github.com/grokify/structured-goals/schema"
)

func main() {
	gen := schema.NewGenerator()

	// Generate all schemas
	if err := gen.GenerateAll("schema"); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating schemas: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Generated schema/okr.schema.json")
	fmt.Println("Generated schema/v2mom.schema.json")
	fmt.Println("Generated schema/dmaic.schema.json")
}
