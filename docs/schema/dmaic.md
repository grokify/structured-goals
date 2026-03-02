# DMAIC Schema

JSON Schema definition for DMAIC (Define, Measure, Analyze, Improve, Control) documents.

## Schema URL

```
https://github.com/grokify/structured-goals/schema/dmaic.schema.json
```

## Document Structure

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/dmaic.schema.json",
  "metadata": { },
  "categories": [ ],
  "initiatives": [ ]
}
```

## Properties

### metadata (required)

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| id | string | Yes | Unique identifier |
| name | string | Yes | Document name |
| owner | string | Yes | Document owner |
| team | string | No | Team name |
| period | string | No | Time period |
| version | string | No | Document version |
| status | string | No | Document status |
| reviewCadence | string | No | Review frequency |

### categories (required)

Array of category objects:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| id | string | No | Unique identifier |
| name | string | Yes | Category name |
| description | string | No | Detailed description |
| order | integer | No | Display order |
| owner | string | No | Category owner |
| metrics | array | Yes | Array of metrics |

### metrics

Array of metric objects within each category:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| id | string | No | Unique identifier |
| name | string | Yes | Metric name |
| description | string | No | Detailed description |
| order | integer | No | Display order |
| baseline | number | Yes | Starting value |
| current | number | Yes | Current value |
| target | number | Yes | Target value |
| unit | string | No | Unit of measurement |
| phase | string | Yes | DMAIC phase |
| trendDirection | string | No | Trend interpretation |
| controlLimits | object | No | SPC control limits |
| thresholds | object | No | Warning/critical thresholds |
| frequency | string | No | Measurement frequency |
| dataSource | string | No | Data source |
| owner | string | No | Metric owner |
| status | string | No | Current status |
| processCapability | object | No | Six Sigma capability |
| dataPoints | array | No | Historical data |
| rootCauses | array | No | Root cause analysis |

### phase (enum)

| Value | Description |
|-------|-------------|
| Define | Problem identification |
| Measure | Data collection |
| Analyze | Root cause analysis |
| Improve | Solution implementation |
| Control | Process monitoring |

### trendDirection (enum)

| Value | Description |
|-------|-------------|
| higher_better | Higher values indicate improvement |
| lower_better | Lower values indicate improvement |
| target_value | Specific target is optimal |

### controlLimits

| Property | Type | Description |
|----------|------|-------------|
| ucl | number | Upper Control Limit |
| lcl | number | Lower Control Limit |
| centerLine | number | Target/mean value |
| sigma | number | Standard deviation |

### thresholds

| Property | Type | Description |
|----------|------|-------------|
| warning | number | Warning threshold |
| critical | number | Critical threshold |

### processCapability

| Property | Type | Description |
|----------|------|-------------|
| cp | number | Process Capability |
| cpk | number | Process Capability Index |
| sigmaLevel | number | Sigma level (1-6) |
| dpmo | integer | Defects Per Million Opportunities |

### initiatives

Array of initiative objects:

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| id | string | No | Unique identifier |
| name | string | Yes | Initiative name |
| description | string | No | Detailed description |
| owner | string | No | Initiative owner |
| status | string | No | Current status |
| startDate | string | No | Start date (ISO 8601) |
| endDate | string | No | End date (ISO 8601) |
| metricIds | array | No | Linked metric IDs |
| expectedImpact | string | No | Expected outcome |
| actualImpact | string | No | Actual outcome |

## Example

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/dmaic.schema.json",
  "metadata": {
    "id": "DMAIC-2025-001",
    "name": "Manufacturing Quality Metrics",
    "owner": "Maria Rodriguez",
    "team": "Quality Engineering",
    "period": "FY2025-Q1",
    "reviewCadence": "Weekly"
  },
  "categories": [
    {
      "id": "production",
      "name": "Production Quality",
      "order": 1,
      "metrics": [
        {
          "id": "prod-001",
          "name": "First Pass Yield",
          "baseline": 85.0,
          "current": 94.2,
          "target": 98.0,
          "unit": "%",
          "phase": "Control",
          "trendDirection": "higher_better",
          "controlLimits": {
            "ucl": 99.0,
            "lcl": 90.0,
            "centerLine": 95.0,
            "sigma": 1.5
          },
          "thresholds": {
            "warning": 92.0,
            "critical": 88.0
          },
          "status": "Green"
        }
      ]
    }
  ],
  "initiatives": [
    {
      "id": "init-001",
      "name": "Six Sigma Black Belt Project",
      "owner": "Tom Chen",
      "status": "In Progress",
      "metricIds": ["prod-001"],
      "expectedImpact": "Achieve 4.5 sigma level"
    }
  ]
}
```

## Related

- [DMAIC Framework](../frameworks/dmaic.md)
- [DMAIC Examples](../examples/dmaic.md)
- [Schema Overview](overview.md)
