# JSON Schema Overview

Structured Goals provides JSON Schema definitions for validating strategic planning documents.

## Available Schemas

| Schema | Description | File |
|--------|-------------|------|
| V2MOM | Vision, Values, Methods, Obstacles, Measures | `v2mom.schema.json` |
| OKR | Objectives and Key Results | `okr.schema.json` |
| DMAIC | Define, Measure, Analyze, Improve, Control | `dmaic.schema.json` |

## Schema Locations

Schemas are available at:

```
https://github.com/grokify/structured-goals/schema/{framework}.schema.json
```

## Using Schemas

### In JSON Documents

Reference the schema in your JSON file:

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/dmaic.schema.json",
  "metadata": {
    "id": "DMAIC-2025-001",
    "name": "Quality Metrics"
  }
}
```

### VS Code Integration

VS Code automatically validates JSON files with `$schema` references. For additional features:

1. Install the "JSON Schema Validator" extension
2. Add to `.vscode/settings.json`:

```json
{
  "json.schemas": [
    {
      "fileMatch": ["**/dmaic*.json"],
      "url": "./schema/dmaic.schema.json"
    }
  ]
}
```

### Go Validation

Use embedded schemas for programmatic validation:

```go
import (
    "github.com/grokify/structured-goals/schema"
    "github.com/xeipuuv/gojsonschema"
)

func validateDMAIC(jsonData []byte) error {
    schemaLoader := gojsonschema.NewStringLoader(schema.DMAICJSONString())
    docLoader := gojsonschema.NewBytesLoader(jsonData)

    result, err := gojsonschema.Validate(schemaLoader, docLoader)
    if err != nil {
        return err
    }

    if !result.Valid() {
        for _, err := range result.Errors() {
            fmt.Printf("- %s\n", err)
        }
        return fmt.Errorf("validation failed")
    }

    return nil
}
```

## Schema Generation

Schemas are generated from Go types using `invopop/jsonschema`:

```bash
go run cmd/genschema/main.go
```

This generates:

- `schema/v2mom.schema.json`
- `schema/okr.schema.json`
- `schema/dmaic.schema.json`

## Embedded Schemas

Schemas are embedded in the Go binary for runtime access:

```go
import "github.com/grokify/structured-goals/schema"

// Get schema as []byte
jsonBytes := schema.DMAICJSON()

// Get schema as string
jsonString := schema.DMAICJSONString()
```

## Schema Versioning

Schemas follow semantic versioning aligned with the library:

- Major version changes indicate breaking schema changes
- Minor version changes add optional fields
- Patch version changes are documentation/clarification only

## Related

- [V2MOM Schema](v2mom.md)
- [OKR Schema](okr.md)
- [DMAIC Schema](dmaic.md)
- [genschema CLI](../cli/genschema.md)
