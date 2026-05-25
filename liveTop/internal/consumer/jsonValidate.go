package consumer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/jsonschema-go/jsonschema"
)

func validateJSON(jsonSchemaPath string, data interface{}) error {
	schemaBytes, err := os.ReadFile(jsonSchemaPath)
	if err != nil {
		return fmt.Errorf("Error reading JSON Schema file: %w", err)
	}

	var schema jsonschema.Schema
	if err := json.Unmarshal(schemaBytes, &schema); err != nil {
		return fmt.Errorf("Error parsing JSON Schema: %w", err)
	}

	resolved, err := schema.Resolve(&jsonschema.ResolveOptions{})
	if err != nil {
		return fmt.Errorf("Error resolving JSON Schema: %w", err)
	}

	return resolved.Validate(data)
}
