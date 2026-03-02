# API Reference

Go library API documentation for Structured Goals.

## Packages

| Package | Description |
|---------|-------------|
| `github.com/grokify/structured-goals/v2mom` | V2MOM framework |
| `github.com/grokify/structured-goals/okr` | OKR framework |
| `github.com/grokify/structured-goals/dmaic` | DMAIC framework |
| `github.com/grokify/structured-goals/schema` | JSON Schema embedding |

## DMAIC Package

### Types

#### DMAICDocument

```go
type DMAICDocument struct {
    Metadata    Metadata     `json:"metadata"`
    Categories  []Category   `json:"categories"`
    Initiatives []Initiative `json:"initiatives,omitempty"`
}
```

#### Metadata

```go
type Metadata struct {
    ID            string    `json:"id"`
    Name          string    `json:"name"`
    Description   string    `json:"description,omitempty"`
    Owner         string    `json:"owner"`
    Team          string    `json:"team,omitempty"`
    Period        string    `json:"period,omitempty"`
    Version       string    `json:"version,omitempty"`
    Status        string    `json:"status,omitempty"`
    ReviewCadence string    `json:"reviewCadence,omitempty"`
    CreatedAt     time.Time `json:"createdAt,omitzero"`
    UpdatedAt     time.Time `json:"updatedAt,omitzero"`
}
```

#### Category

```go
type Category struct {
    ID          string   `json:"id,omitempty"`
    Name        string   `json:"name"`
    Description string   `json:"description,omitempty"`
    Order       int      `json:"order,omitempty"`
    Owner       string   `json:"owner,omitempty"`
    Metrics     []Metric `json:"metrics"`
}
```

#### Metric

```go
type Metric struct {
    ID                string             `json:"id,omitempty"`
    Name              string             `json:"name"`
    Description       string             `json:"description,omitempty"`
    Order             int                `json:"order,omitempty"`
    Baseline          float64            `json:"baseline"`
    Current           float64            `json:"current"`
    Target            float64            `json:"target"`
    Unit              string             `json:"unit,omitempty"`
    Phase             string             `json:"phase"`
    TrendDirection    string             `json:"trendDirection,omitempty"`
    ControlLimits     *ControlLimits     `json:"controlLimits,omitempty"`
    Thresholds        *Thresholds        `json:"thresholds,omitempty"`
    Frequency         string             `json:"frequency,omitempty"`
    DataSource        string             `json:"dataSource,omitempty"`
    Owner             string             `json:"owner,omitempty"`
    Status            string             `json:"status,omitempty"`
    ProcessCapability *ProcessCapability `json:"processCapability,omitempty"`
    DataPoints        []DataPoint        `json:"dataPoints,omitempty"`
    RootCauses        []RootCause        `json:"rootCauses,omitempty"`
    InitiativeIDs     []string           `json:"initiativeIds,omitempty"`
    ExternalLinks     map[string]string  `json:"externalLinks,omitempty"`
}
```

### Constants

```go
// DMAIC Phases
const (
    PhaseDefine  = "Define"
    PhaseMeasure = "Measure"
    PhaseAnalyze = "Analyze"
    PhaseImprove = "Improve"
    PhaseControl = "Control"
)

// Trend Directions
const (
    TrendHigherBetter = "higher_better"
    TrendLowerBetter  = "lower_better"
    TrendTargetValue  = "target_value"
)

// Status Values
const (
    StatusGreen  = "Green"
    StatusYellow = "Yellow"
    StatusRed    = "Red"
)
```

### Functions

#### Constructor

```go
func New(id, name, owner string) *DMAICDocument
```

Creates a new DMAIC document with the given metadata.

#### File Operations

```go
func ReadFile(filename string) (*DMAICDocument, error)
func Parse(data []byte) (*DMAICDocument, error)
func (doc *DMAICDocument) JSON() ([]byte, error)
func (doc *DMAICDocument) WriteFile(filename string) error
```

#### Analysis Methods

```go
func (doc *DMAICDocument) AllMetrics() []Metric
func (doc *DMAICDocument) MetricsByPhase() map[string][]Metric
func (doc *DMAICDocument) MetricsByStatus() map[string][]Metric
func (doc *DMAICDocument) CalculateOverallHealth() float64
func (cat *Category) CalculateCategoryHealth() float64
func (m *Metric) IsInControl() bool
func (m *Metric) CalculateStatus() string
```

#### Validation

```go
func (doc *DMAICDocument) Validate(opts ValidationOptions) []ValidationError
func DefaultValidationOptions() ValidationOptions
func StrictValidationOptions() ValidationOptions
func SixSigmaValidationOptions() ValidationOptions
```

## OKR Package

### Types

```go
type OKRDocument struct {
    Metadata   Metadata    `json:"metadata"`
    Objectives []Objective `json:"objectives"`
}

type Objective struct {
    ID          string      `json:"id,omitempty"`
    Name        string      `json:"name"`
    Description string      `json:"description,omitempty"`
    Order       int         `json:"order,omitempty"`
    Owner       string      `json:"owner,omitempty"`
    KeyResults  []KeyResult `json:"keyResults"`
}

type KeyResult struct {
    ID          string  `json:"id,omitempty"`
    Name        string  `json:"name"`
    Description string  `json:"description,omitempty"`
    Baseline    float64 `json:"baseline"`
    Current     float64 `json:"current"`
    Target      float64 `json:"target"`
    Unit        string  `json:"unit,omitempty"`
    Order       int     `json:"order,omitempty"`
    Owner       string  `json:"owner,omitempty"`
    Confidence  float64 `json:"confidence,omitempty"`
}
```

### Functions

```go
func New(id, name, owner string) *OKRDocument
func ReadFile(filename string) (*OKRDocument, error)
func (doc *OKRDocument) CalculateProgress() float64
func (obj *Objective) CalculateProgress() float64
func (kr *KeyResult) CalculateProgress() float64
func (doc *OKRDocument) Validate(opts ValidationOptions) []ValidationError
```

## Schema Package

### Embedded Schemas

```go
func V2MOMJSON() []byte
func V2MOMJSONString() string
func OKRJSON() []byte
func OKRJSONString() string
func DMAICJSON() []byte
func DMAICJSONString() string
```

### Schema Generation

```go
func GenerateV2MOMSchema() (*jsonschema.Schema, error)
func GenerateOKRSchema() (*jsonschema.Schema, error)
func GenerateDMAICSchema() (*jsonschema.Schema, error)
func WriteV2MOMSchema(filename string) error
func WriteOKRSchema(filename string) error
func WriteDMAICSchema(filename string) error
func GenerateAll(outputDir string) error
```

## Full Documentation

For complete API documentation, see [pkg.go.dev](https://pkg.go.dev/github.com/grokify/structured-goals).
