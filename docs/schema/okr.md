# OKR Schema

JSON Schema definition for OKR (Objectives and Key Results) documents.

## Schema URL

```
https://github.com/grokify/structured-goals/schema/okr.schema.json
```

## Document Structure

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/okr.schema.json",
  "metadata": { },
  "objectives": [ ]
}
```

## Properties

### metadata (required)

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| id | string | Yes | Unique identifier |
| name | string | Yes | Document name |
| owner | string | Yes | Document owner |
| period | string | No | Time period (e.g., "2025-Q1") |
| version | string | No | Document version |
| status | string | No | Document status |
| team | string | No | Team name |

### objectives (required)

Array of objective objects:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| id | string | No | Unique identifier |
| name | string | Yes | Objective name |
| description | string | No | Detailed description |
| order | integer | No | Display order |
| owner | string | No | Objective owner |
| keyResults | array | Yes | Array of key results |

### keyResults

Array of key result objects within each objective:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| id | string | No | Unique identifier |
| name | string | Yes | Key result name |
| description | string | No | Detailed description |
| baseline | number | Yes | Starting value |
| current | number | Yes | Current value |
| target | number | Yes | Target value |
| unit | string | No | Unit of measurement |
| order | integer | No | Display order |
| owner | string | No | Key result owner |
| confidence | number | No | Confidence level (0-1) |

## Example

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/okr.schema.json",
  "metadata": {
    "id": "OKR-2025-Q1",
    "name": "Q1 Engineering Objectives",
    "owner": "VP Engineering",
    "team": "Engineering",
    "period": "2025-Q1",
    "version": "1.0.0"
  },
  "objectives": [
    {
      "id": "obj-001",
      "name": "Improve Platform Reliability",
      "description": "Ensure our platform meets enterprise SLA requirements",
      "order": 1,
      "keyResults": [
        {
          "id": "kr-001",
          "name": "Achieve 99.95% uptime",
          "baseline": 99.5,
          "current": 99.8,
          "target": 99.95,
          "unit": "%",
          "order": 1
        },
        {
          "id": "kr-002",
          "name": "Reduce P1 incidents",
          "description": "Reduce P1 incidents to less than 2 per month",
          "baseline": 8,
          "current": 4,
          "target": 2,
          "unit": "incidents/month",
          "order": 2
        },
        {
          "id": "kr-003",
          "name": "Improve MTTR",
          "description": "Reduce mean time to recovery to under 30 minutes",
          "baseline": 120,
          "current": 45,
          "target": 30,
          "unit": "minutes",
          "order": 3
        }
      ]
    },
    {
      "id": "obj-002",
      "name": "Accelerate Development Velocity",
      "description": "Increase team productivity and delivery speed",
      "order": 2,
      "keyResults": [
        {
          "id": "kr-004",
          "name": "Increase deployment frequency",
          "baseline": 2,
          "current": 8,
          "target": 20,
          "unit": "deploys/week",
          "order": 1
        },
        {
          "id": "kr-005",
          "name": "Reduce lead time",
          "description": "Time from commit to production",
          "baseline": 168,
          "current": 24,
          "target": 8,
          "unit": "hours",
          "order": 2
        }
      ]
    }
  ]
}
```

## Progress Calculation

Key result progress is calculated as:

```
progress = (current - baseline) / (target - baseline)
```

Objective progress is the average of its key results.

## Related

- [OKR Framework](../frameworks/okr.md)
- [Schema Overview](overview.md)
