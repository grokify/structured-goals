// Package schema provides embedded JSON Schemas for V2MOM and OKR validation.
package schema

import (
	_ "embed"
)

// SchemaVersion is the version of the JSON Schemas.
const SchemaVersion = "1.0.0"

// V2MOM Schema constants.
const (
	// SchemaID is the canonical ID for the V2MOM JSON Schema.
	SchemaID = "https://github.com/grokify/structured-goals/schema/v2mom.schema.json"
)

// OKR Schema constants.
const (
	// OKRSchemaVersion is the version of the OKR JSON Schema.
	OKRSchemaVersion = "1.0.0"
)

//go:embed v2mom.schema.json
var schemaJSON []byte

//go:embed okr.schema.json
var okrSchemaJSON []byte

//go:embed dmaic.schema.json
var dmaicSchemaJSON []byte

// JSON returns the embedded V2MOM JSON Schema as bytes.
func JSON() []byte {
	return schemaJSON
}

// JSONString returns the embedded V2MOM JSON Schema as a string.
func JSONString() string {
	return string(schemaJSON)
}

// OKRJSON returns the embedded OKR JSON Schema as bytes.
func OKRJSON() []byte {
	return okrSchemaJSON
}

// OKRJSONString returns the embedded OKR JSON Schema as a string.
func OKRJSONString() string {
	return string(okrSchemaJSON)
}

// DMAICJSON returns the embedded DMAIC JSON Schema as bytes.
func DMAICJSON() []byte {
	return dmaicSchemaJSON
}

// DMAICJSONString returns the embedded DMAIC JSON Schema as a string.
func DMAICJSONString() string {
	return string(dmaicSchemaJSON)
}
