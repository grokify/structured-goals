# genschema CLI

Generate JSON Schema files from Go type definitions.

## Installation

```bash
go install github.com/grokify/structured-goals/cmd/genschema@latest
```

## Usage

```bash
genschema [options]
```

## Options

| Flag | Description | Default |
|------|-------------|---------|
| `-o, --output` | Output directory | `./schema` |
| `-f, --format` | Schema format (json, yaml) | `json` |
| `--all` | Generate all schemas | `true` |
| `--v2mom` | Generate V2MOM schema only | `false` |
| `--okr` | Generate OKR schema only | `false` |
| `--dmaic` | Generate DMAIC schema only | `false` |

## Examples

### Generate All Schemas

```bash
genschema
# Generated schema/v2mom.schema.json
# Generated schema/okr.schema.json
# Generated schema/dmaic.schema.json
```

### Generate Specific Schema

```bash
genschema --dmaic
# Generated schema/dmaic.schema.json
```

### Custom Output Directory

```bash
genschema -o ./my-schemas
# Generated my-schemas/v2mom.schema.json
# Generated my-schemas/okr.schema.json
# Generated my-schemas/dmaic.schema.json
```

## Generated Files

| File | Description |
|------|-------------|
| `v2mom.schema.json` | V2MOM document schema |
| `okr.schema.json` | OKR document schema |
| `dmaic.schema.json` | DMAIC document schema |

## Schema Features

Generated schemas include:

- Type definitions for all document structures
- Required field validation
- Enum constraints for known values (phases, status, etc.)
- Pattern validation for IDs
- Format validation for dates

## Using Generated Schemas

### VS Code Validation

Add to your JSON file:

```json
{
  "$schema": "./schema/dmaic.schema.json",
  "metadata": { ... }
}
```

### Programmatic Validation

```go
import (
    "github.com/grokify/structured-goals/schema"
    "github.com/xeipuuv/gojsonschema"
)

schemaLoader := gojsonschema.NewStringLoader(schema.DMAICJSONString())
docLoader := gojsonschema.NewReferenceLoader("file:///path/to/doc.json")

result, err := gojsonschema.Validate(schemaLoader, docLoader)
```

## Related

- [Schema Overview](../schema/overview.md)
- [genpandoc CLI](genpandoc.md)
