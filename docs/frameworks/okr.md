# OKR Framework

OKR (Objectives and Key Results) is a goal-setting framework popularized by Intel and Google for defining and tracking objectives and their outcomes.

## Overview

OKRs consist of two components:

| Component | Description | Characteristics |
|-----------|-------------|-----------------|
| **Objective** | What you want to achieve | Qualitative, inspirational, time-bound |
| **Key Results** | How you'll measure success | Quantitative, measurable, specific |

## Document Structure

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/okr.schema.json",
  "metadata": {
    "id": "OKR-2025-Q1",
    "name": "Q1 Engineering Objectives",
    "owner": "VP Engineering",
    "period": "2025-Q1"
  },
  "objectives": [
    {
      "id": "obj-001",
      "name": "Improve Platform Reliability",
      "description": "Ensure our platform meets enterprise SLA requirements",
      "keyResults": [
        {
          "id": "kr-001",
          "name": "Achieve 99.95% uptime",
          "baseline": 99.5,
          "current": 99.8,
          "target": 99.95,
          "unit": "%"
        },
        {
          "id": "kr-002",
          "name": "Reduce P1 incidents to less than 2 per month",
          "baseline": 8,
          "current": 4,
          "target": 2,
          "unit": "incidents/month"
        }
      ]
    }
  ]
}
```

## Go Usage

```go
import "github.com/grokify/structured-goals/okr"

// Create a new OKR document
doc := okr.New("OKR-2025-Q1", "Q1 Engineering Objectives", "VP Engineering")

// Add an objective with key results
doc.Objectives = append(doc.Objectives, okr.Objective{
    ID:          "obj-001",
    Name:        "Improve Platform Reliability",
    Description: "Ensure our platform meets enterprise SLA requirements",
    KeyResults: []okr.KeyResult{
        {
            ID:       "kr-001",
            Name:     "Achieve 99.95% uptime",
            Baseline: 99.5,
            Current:  99.8,
            Target:   99.95,
            Unit:     "%",
        },
    },
})

// Calculate progress
progress := doc.CalculateProgress()
fmt.Printf("Overall Progress: %.0f%%\n", progress*100)

// Validate
errors := doc.Validate(okr.DefaultValidationOptions())
```

## Progress Calculation

Key Result progress is calculated as:

```
progress = (current - baseline) / (target - baseline)
```

For example:

- Baseline: 99.5%
- Current: 99.8%
- Target: 99.95%
- Progress: (99.8 - 99.5) / (99.95 - 99.5) = 0.3 / 0.45 = 66.7%

## Grading Scale

OKRs typically use a 0.0 to 1.0 grading scale:

| Score | Meaning |
|-------|---------|
| 0.0 - 0.3 | Failed to make progress |
| 0.4 - 0.6 | Made progress but fell short |
| 0.7 - 1.0 | Achieved or exceeded target |

!!! note "Stretch Goals"
    OKRs are often set as stretch goals. Achieving 70% is considered success.

## Best Practices

1. **3-5 Objectives per period** - Focus on what matters most
2. **3-5 Key Results per Objective** - Enough to measure, not overwhelming
3. **Ambitious but achievable** - Stretch goals that inspire
4. **Regular check-ins** - Weekly or bi-weekly progress reviews
5. **Separate from compensation** - OKRs are for alignment, not performance reviews

## Cadence

| Level | Typical Cadence |
|-------|-----------------|
| Company | Annual |
| Team | Quarterly |
| Individual | Quarterly |

## Related

- [OKR Schema](../schema/okr.md)
- [OKR API Reference](../api.md#okr-package)
