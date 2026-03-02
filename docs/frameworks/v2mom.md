# V2MOM Framework

V2MOM (Vision, Values, Methods, Obstacles, Measures) is a strategic planning methodology developed at Salesforce by Marc Benioff.

## Overview

V2MOM provides a structured approach to strategic planning by defining:

| Component | Description | Example |
|-----------|-------------|---------|
| **Vision** | What you want to achieve | "Become the #1 customer success platform" |
| **Values** | Guiding principles | "Customer First", "Innovation", "Trust" |
| **Methods** | How you'll achieve the vision | "Expand enterprise sales", "Launch mobile app" |
| **Obstacles** | Challenges to overcome | "Market competition", "Technical debt" |
| **Measures** | How you'll track success | "Revenue growth", "Customer retention" |

## Document Structure

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/v2mom.schema.json",
  "metadata": {
    "id": "V2MOM-2025-001",
    "name": "Company Strategy 2025",
    "owner": "CEO",
    "period": "FY2025"
  },
  "vision": "Become the leading platform for strategic planning",
  "values": [
    {
      "name": "Simplicity",
      "description": "Make complex planning accessible"
    }
  ],
  "methods": [
    {
      "name": "Launch self-service platform",
      "description": "Enable customers to create plans without consulting"
    }
  ],
  "obstacles": [
    {
      "name": "Market awareness",
      "description": "Low brand recognition in target market"
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

## Go Usage

```go
import "github.com/grokify/structured-goals/v2mom"

// Create a new V2MOM document
doc := v2mom.New("V2MOM-2025-001", "Company Strategy", "CEO")
doc.Vision = "Become the leading platform for strategic planning"

// Add values
doc.Values = append(doc.Values, v2mom.Value{
    Name:        "Simplicity",
    Description: "Make complex planning accessible",
})

// Add methods
doc.Methods = append(doc.Methods, v2mom.Method{
    Name:        "Launch self-service platform",
    Description: "Enable customers to create plans without consulting",
})

// Validate
errors := doc.Validate(v2mom.DefaultValidationOptions())
```

## Structure Options

V2MOM supports three structural approaches:

### Flat Structure

All components at the same level:

```json
{
  "values": ["Value 1", "Value 2"],
  "methods": ["Method 1", "Method 2"]
}
```

### Nested Structure

Methods nested under values:

```json
{
  "values": [
    {
      "name": "Value 1",
      "methods": ["Method 1a", "Method 1b"]
    }
  ]
}
```

### Hybrid Structure

Mix of flat and nested:

```json
{
  "values": [
    {
      "name": "Value 1",
      "methods": ["Method 1a"]
    }
  ],
  "methods": ["Standalone Method"]
}
```

## Best Practices

1. **Keep Vision concise** - One sentence that inspires
2. **Limit Values to 3-5** - Focus on what truly matters
3. **Make Methods actionable** - Specific, time-bound actions
4. **Be honest about Obstacles** - Acknowledge real challenges
5. **Use quantifiable Measures** - Numbers you can track

## Related

- [V2MOM Schema](../schema/v2mom.md)
- [V2MOM Examples](../examples/v2mom.md)
- [V2MOM CLI](../cli/v2mom.md)
