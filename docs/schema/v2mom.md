# V2MOM Schema

JSON Schema definition for V2MOM (Vision, Values, Methods, Obstacles, Measures) documents.

## Schema URL

```
https://github.com/grokify/structured-goals/schema/v2mom.schema.json
```

## Document Structure

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/v2mom.schema.json",
  "metadata": { },
  "vision": "",
  "values": [ ],
  "methods": [ ],
  "obstacles": [ ],
  "measures": [ ]
}
```

## Properties

### metadata (required)

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| id | string | Yes | Unique identifier |
| name | string | Yes | Document name |
| owner | string | Yes | Document owner |
| period | string | No | Time period (e.g., "FY2025") |
| version | string | No | Document version |
| status | string | No | Document status |

### vision (required)

A single string describing the vision statement.

### values (required)

Array of value objects:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| name | string | Yes | Value name |
| description | string | No | Detailed description |
| order | integer | No | Display order |

### methods (required)

Array of method objects:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| name | string | Yes | Method name |
| description | string | No | Detailed description |
| order | integer | No | Display order |
| owner | string | No | Method owner |

### obstacles (required)

Array of obstacle objects:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| name | string | Yes | Obstacle name |
| description | string | No | Detailed description |
| severity | string | No | Impact severity |
| mitigation | string | No | Mitigation strategy |

### measures (required)

Array of measure objects:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| name | string | Yes | Measure name |
| description | string | No | Detailed description |
| baseline | number | No | Starting value |
| current | number | No | Current value |
| target | number | No | Target value |
| unit | string | No | Unit of measurement |

## Example

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/v2mom.schema.json",
  "metadata": {
    "id": "V2MOM-2025-001",
    "name": "Company Strategy 2025",
    "owner": "CEO",
    "period": "FY2025",
    "version": "1.0.0"
  },
  "vision": "Become the leading platform for strategic planning",
  "values": [
    {
      "name": "Simplicity",
      "description": "Make complex planning accessible to everyone",
      "order": 1
    },
    {
      "name": "Transparency",
      "description": "Open communication at all levels",
      "order": 2
    }
  ],
  "methods": [
    {
      "name": "Launch self-service platform",
      "description": "Enable customers to create plans independently",
      "order": 1,
      "owner": "Product Team"
    }
  ],
  "obstacles": [
    {
      "name": "Market awareness",
      "description": "Low brand recognition in target market",
      "severity": "Medium",
      "mitigation": "Increase marketing spend and PR efforts"
    }
  ],
  "measures": [
    {
      "name": "Monthly Active Users",
      "baseline": 1000,
      "current": 2500,
      "target": 10000,
      "unit": "users"
    }
  ]
}
```

## Related

- [V2MOM Framework](../frameworks/v2mom.md)
- [V2MOM Examples](../examples/v2mom.md)
- [Schema Overview](overview.md)
