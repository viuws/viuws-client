// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"
import "reflect"

type IOCardinality string

const IOCardinalityMultiple IOCardinality = "multiple"
const IOCardinalitySingle IOCardinality = "single"

type InputSpec struct {
	// Cardinality corresponds to the JSON schema field "cardinality".
	Cardinality interface{} `json:"cardinality,omitempty"`

	// Description corresponds to the JSON schema field "description".
	Description string `json:"description"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name"`

	// Required corresponds to the JSON schema field "required".
	Required bool `json:"required,omitempty"`

	// SupportedFilePatterns corresponds to the JSON schema field
	// "supportedFilePatterns".
	SupportedFilePatterns interface{} `json:"supportedFilePatterns,omitempty"`
}

type ModuleJson struct {
	// ArgsSchema corresponds to the JSON schema field "argsSchema".
	ArgsSchema interface{} `json:"argsSchema,omitempty"`

	// ArgsUISchema corresponds to the JSON schema field "argsUISchema".
	ArgsUISchema interface{} `json:"argsUISchema,omitempty"`

	// Container corresponds to the JSON schema field "container".
	Container OCIRuntimeConfig `json:"container"`

	// Description corresponds to the JSON schema field "description".
	Description string `json:"description"`

	// EnvSchema corresponds to the JSON schema field "envSchema".
	EnvSchema interface{} `json:"envSchema,omitempty"`

	// EnvUISchema corresponds to the JSON schema field "envUISchema".
	EnvUISchema interface{} `json:"envUISchema,omitempty"`

	// Inputs corresponds to the JSON schema field "inputs".
	Inputs ModuleJsonInputs `json:"inputs,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name"`

	// Outputs corresponds to the JSON schema field "outputs".
	Outputs ModuleJsonOutputs `json:"outputs,omitempty"`

	// SchemaVersion corresponds to the JSON schema field "schemaVersion".
	SchemaVersion string `json:"schemaVersion,omitempty"`
}

type ModuleJsonInputs map[string]InputSpec

// UnmarshalJSON implements json.Unmarshaler.
func (j *InputSpec) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["description"]; !ok || v == nil {
		return fmt.Errorf("field description in InputSpec: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in InputSpec: required")
	}
	type Plain InputSpec
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["cardinality"]; !ok || v == nil {
		plain.Cardinality = "multiple"
	}
	if v, ok := raw["required"]; !ok || v == nil {
		plain.Required = false
	}
	*j = InputSpec(plain)
	return nil
}

type OCIRuntimeConfig struct {
	// Args corresponds to the JSON schema field "args".
	Args []string `json:"args"`

	// Cwd corresponds to the JSON schema field "cwd".
	Cwd interface{} `json:"cwd,omitempty"`

	// Env corresponds to the JSON schema field "env".
	Env OCIRuntimeConfigEnv `json:"env,omitempty"`

	// Image corresponds to the JSON schema field "image".
	Image string `json:"image"`

	// Tag corresponds to the JSON schema field "tag".
	Tag string `json:"tag,omitempty"`
}

type OCIRuntimeConfigEnv map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCIRuntimeConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["args"]; !ok || v == nil {
		return fmt.Errorf("field args in OCIRuntimeConfig: required")
	}
	if v, ok := raw["image"]; !ok || v == nil {
		return fmt.Errorf("field image in OCIRuntimeConfig: required")
	}
	type Plain OCIRuntimeConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["env"]; !ok || v == nil {
		plain.Env = map[string]interface{}{}
	}
	if v, ok := raw["tag"]; !ok || v == nil {
		plain.Tag = "latest"
	}
	*j = OCIRuntimeConfig(plain)
	return nil
}

type OutputSpec struct {
	// Cardinality corresponds to the JSON schema field "cardinality".
	Cardinality interface{} `json:"cardinality,omitempty"`

	// Description corresponds to the JSON schema field "description".
	Description string `json:"description"`

	// GeneratedFilePattern corresponds to the JSON schema field
	// "generatedFilePattern".
	GeneratedFilePattern interface{} `json:"generatedFilePattern,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OutputSpec) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["description"]; !ok || v == nil {
		return fmt.Errorf("field description in OutputSpec: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in OutputSpec: required")
	}
	type Plain OutputSpec
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["cardinality"]; !ok || v == nil {
		plain.Cardinality = "multiple"
	}
	*j = OutputSpec(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IOCardinality) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IOCardinality {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IOCardinality, v)
	}
	*j = IOCardinality(v)
	return nil
}

type ModuleJsonOutputs map[string]OutputSpec

var enumValues_IOCardinality = []interface{}{
	"single",
	"multiple",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModuleJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["container"]; !ok || v == nil {
		return fmt.Errorf("field container in ModuleJson: required")
	}
	if v, ok := raw["description"]; !ok || v == nil {
		return fmt.Errorf("field description in ModuleJson: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in ModuleJson: required")
	}
	type Plain ModuleJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["inputs"]; !ok || v == nil {
		plain.Inputs = ModuleJsonInputs{}
	}
	if v, ok := raw["outputs"]; !ok || v == nil {
		plain.Outputs = ModuleJsonOutputs{}
	}
	if v, ok := raw["schemaVersion"]; !ok || v == nil {
		plain.SchemaVersion = "1.0.0"
	}
	*j = ModuleJson(plain)
	return nil
}
