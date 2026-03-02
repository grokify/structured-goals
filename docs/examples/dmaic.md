# DMAIC Examples

Example DMAIC documents for various industries and use cases.

## Available Examples

The repository includes three comprehensive DMAIC examples:

| Example | Categories | Metrics | Industry |
|---------|------------|---------|----------|
| Manufacturing | 3 | 9 | Manufacturing/Quality |
| Software Development | 4 | 12 | Technology |
| Customer Service | 4 | 12 | Support/Service |

## Manufacturing Example

Quality metrics for manufacturing operations:

**Categories:**

- Production Quality (First Pass Yield, DPMO, Scrap Rate)
- Equipment Performance (OEE, MTBF, Unplanned Downtime)
- Process Control (Cpk, Cycle Time Variation, SPC Compliance)

**Key Features:**

- Process capability metrics (Cp, Cpk, Sigma level)
- Statistical process control limits
- Six Sigma initiative tracking

```json
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
  "processCapability": {
    "cp": 1.45,
    "cpk": 1.33,
    "sigmaLevel": 4.0,
    "dpmo": 6210
  }
}
```

## Software Development Example

DORA metrics and development quality tracking:

**Categories:**

- Code Quality (Coverage, Tech Debt, Review Coverage)
- Delivery Performance (Deployment Frequency, Lead Time, Change Failure Rate)
- System Reliability (Availability, MTTR, P95 Response Time)
- Defect Management (Bug Escape Rate, Fix Cycle Time, Open Critical Bugs)

**Key Features:**

- DORA metrics alignment
- Root cause analysis for MTTR improvements
- CI/CD pipeline initiatives

```json
{
  "id": "del-001",
  "name": "Deployment Frequency",
  "baseline": 2,
  "current": 12,
  "target": 20,
  "unit": "deploys/week",
  "phase": "Improve",
  "trendDirection": "higher_better",
  "thresholds": {
    "warning": 8,
    "critical": 4
  }
}
```

## Customer Service Example

Support operations and customer satisfaction metrics:

**Categories:**

- Customer Satisfaction (CSAT, NPS, CES)
- Response Performance (First Response Time, SLA Compliance, Handle Time)
- Resolution Quality (FCR, Reopen Rate, Escalation Rate)
- Channel Performance (Self-Service Rate, Bot Containment, Channel Switch)

**Key Features:**

- Multi-channel support tracking
- AI chatbot effectiveness metrics
- Customer effort optimization

```json
{
  "id": "sat-001",
  "name": "Customer Satisfaction Score",
  "baseline": 72.0,
  "current": 88.0,
  "target": 92.0,
  "unit": "%",
  "phase": "Control",
  "trendDirection": "higher_better",
  "controlLimits": {
    "ucl": 95.0,
    "lcl": 85.0,
    "centerLine": 90.0,
    "sigma": 2.5
  }
}
```

## Loading Examples

```go
import "github.com/grokify/structured-goals/dmaic"

// Load manufacturing example
doc, err := dmaic.ReadFile("examples/dmaic-manufacturing.json")
if err != nil {
    log.Fatal(err)
}

// Calculate health
health := doc.CalculateOverallHealth()
fmt.Printf("Overall Health: %.0f%%\n", health*100)

// Get metrics by phase
byPhase := doc.MetricsByPhase()
fmt.Printf("Control phase metrics: %d\n", len(byPhase[dmaic.PhaseControl]))

// Get metrics by status
byStatus := doc.MetricsByStatus()
fmt.Printf("Yellow status metrics: %d\n", len(byStatus[dmaic.StatusYellow]))
```

## Creating Your Own

Start with a template:

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/dmaic.schema.json",
  "metadata": {
    "id": "DMAIC-YOUR-001",
    "name": "Your Quality Metrics",
    "owner": "Your Name",
    "team": "Your Team",
    "period": "FY2025-Q1",
    "reviewCadence": "Weekly"
  },
  "categories": [
    {
      "id": "category-1",
      "name": "Category Name",
      "order": 1,
      "metrics": [
        {
          "id": "metric-001",
          "name": "Metric Name",
          "baseline": 0,
          "current": 0,
          "target": 100,
          "unit": "%",
          "phase": "Define",
          "trendDirection": "higher_better"
        }
      ]
    }
  ],
  "initiatives": []
}
```

## Related

- [DMAIC Framework](../frameworks/dmaic.md)
- [DMAIC Schema](../schema/dmaic.md)
- [Pandoc Renderer](../renderers/pandoc.md)
