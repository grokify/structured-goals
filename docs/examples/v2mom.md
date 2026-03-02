# V2MOM Examples

Example V2MOM documents for various use cases.

## Company Strategy Example

A comprehensive company-level V2MOM:

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/v2mom.schema.json",
  "metadata": {
    "id": "V2MOM-COMPANY-2025",
    "name": "Company Strategy 2025",
    "owner": "CEO",
    "period": "FY2025",
    "version": "1.0.0",
    "status": "Active"
  },
  "vision": "Become the most trusted platform for strategic planning, helping 10,000 organizations achieve their goals",
  "values": [
    {
      "name": "Customer Success",
      "description": "Our customers' success is our success. Every decision starts with the customer.",
      "order": 1
    },
    {
      "name": "Simplicity",
      "description": "Make complex strategic planning accessible to everyone.",
      "order": 2
    },
    {
      "name": "Transparency",
      "description": "Open communication internally and with customers.",
      "order": 3
    },
    {
      "name": "Continuous Improvement",
      "description": "Always be learning, iterating, and getting better.",
      "order": 4
    }
  ],
  "methods": [
    {
      "name": "Launch Enterprise Self-Service Platform",
      "description": "Enable enterprise customers to create, manage, and track strategic plans independently",
      "order": 1,
      "owner": "VP Product"
    },
    {
      "name": "Expand to European Markets",
      "description": "Establish sales presence in UK, Germany, and France",
      "order": 2,
      "owner": "VP Sales"
    },
    {
      "name": "Build Partner Ecosystem",
      "description": "Develop partnerships with consulting firms and system integrators",
      "order": 3,
      "owner": "VP Partnerships"
    },
    {
      "name": "Achieve SOC 2 Type II Certification",
      "description": "Complete security audit and certification process",
      "order": 4,
      "owner": "VP Engineering"
    }
  ],
  "obstacles": [
    {
      "name": "Market Awareness",
      "description": "Low brand recognition in enterprise segment",
      "severity": "High",
      "mitigation": "Increase marketing spend, sponsor industry events, publish thought leadership content"
    },
    {
      "name": "Enterprise Sales Cycle",
      "description": "Long sales cycles (6-12 months) for enterprise deals",
      "severity": "Medium",
      "mitigation": "Offer free trials, develop ROI calculator, build customer references"
    },
    {
      "name": "Technical Debt",
      "description": "Legacy codebase slowing feature development",
      "severity": "Medium",
      "mitigation": "Allocate 20% of engineering time to debt reduction"
    }
  ],
  "measures": [
    {
      "name": "Annual Recurring Revenue",
      "baseline": 5000000,
      "current": 7500000,
      "target": 15000000,
      "unit": "USD"
    },
    {
      "name": "Enterprise Customers",
      "baseline": 50,
      "current": 85,
      "target": 200,
      "unit": "customers"
    },
    {
      "name": "Net Promoter Score",
      "baseline": 45,
      "current": 52,
      "target": 65,
      "unit": "NPS"
    },
    {
      "name": "Employee Satisfaction",
      "baseline": 72,
      "current": 78,
      "target": 85,
      "unit": "%"
    }
  ]
}
```

## Team V2MOM Example

A product team V2MOM aligned with company strategy:

```json
{
  "$schema": "https://github.com/grokify/structured-goals/schema/v2mom.schema.json",
  "metadata": {
    "id": "V2MOM-PRODUCT-2025",
    "name": "Product Team V2MOM 2025",
    "owner": "VP Product",
    "period": "FY2025"
  },
  "vision": "Deliver a self-service platform that customers love, reducing time-to-value by 80%",
  "values": [
    {
      "name": "User-Centric Design",
      "description": "Every feature starts with user research"
    },
    {
      "name": "Data-Driven Decisions",
      "description": "Measure impact, iterate quickly"
    },
    {
      "name": "Quality First",
      "description": "Ship features that work flawlessly"
    }
  ],
  "methods": [
    {
      "name": "Launch Template Library",
      "description": "Pre-built templates for common strategic frameworks",
      "owner": "Product Manager"
    },
    {
      "name": "Implement AI Assistant",
      "description": "AI-powered suggestions for goals and metrics",
      "owner": "Tech Lead"
    },
    {
      "name": "Build Integration Hub",
      "description": "Connectors for popular business tools",
      "owner": "Integration Team"
    }
  ],
  "obstacles": [
    {
      "name": "Feature Scope Creep",
      "description": "Risk of trying to do too much",
      "mitigation": "Strict prioritization using RICE scoring"
    },
    {
      "name": "Resource Constraints",
      "description": "Limited engineering bandwidth",
      "mitigation": "Focus on high-impact features only"
    }
  ],
  "measures": [
    {
      "name": "Time to First Plan",
      "baseline": 120,
      "target": 15,
      "unit": "minutes"
    },
    {
      "name": "Feature Adoption Rate",
      "baseline": 40,
      "target": 75,
      "unit": "%"
    },
    {
      "name": "Customer Support Tickets",
      "baseline": 500,
      "target": 200,
      "unit": "tickets/month"
    }
  ]
}
```

## Using Examples

Load and work with examples:

```go
import "github.com/grokify/structured-goals/v2mom"

doc, err := v2mom.ReadFile("company-v2mom.json")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Vision: %s\n", doc.Vision)
fmt.Printf("Values: %d\n", len(doc.Values))
fmt.Printf("Methods: %d\n", len(doc.Methods))
```

## Related

- [V2MOM Framework](../frameworks/v2mom.md)
- [V2MOM Schema](../schema/v2mom.md)
