// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package schema

import "encoding/json"
import "fmt"

type ModuleRef struct {
	// Id corresponds to the JSON schema field "id".
	Id string `json:"id"`

	// Path corresponds to the JSON schema field "path".
	Path string `json:"path"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ModuleRef) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in ModuleRef: required")
	}
	if v, ok := raw["path"]; !ok || v == nil {
		return fmt.Errorf("field path in ModuleRef: required")
	}
	type Plain ModuleRef
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ModuleRef(plain)
	return nil
}

type PluginRef struct {
	// Id corresponds to the JSON schema field "id".
	Id string `json:"id"`

	// Path corresponds to the JSON schema field "path".
	Path string `json:"path"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PluginRef) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in PluginRef: required")
	}
	if v, ok := raw["path"]; !ok || v == nil {
		return fmt.Errorf("field path in PluginRef: required")
	}
	type Plain PluginRef
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PluginRef(plain)
	return nil
}

type RegistryJson struct {
	// Modules corresponds to the JSON schema field "modules".
	Modules []ModuleRef `json:"modules,omitempty"`

	// Plugins corresponds to the JSON schema field "plugins".
	Plugins []PluginRef `json:"plugins,omitempty"`

	// Repos corresponds to the JSON schema field "repos".
	Repos []string `json:"repos,omitempty"`

	// SchemaVersion corresponds to the JSON schema field "schemaVersion".
	SchemaVersion string `json:"schemaVersion,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RegistryJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain RegistryJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["modules"]; !ok || v == nil {
		plain.Modules = []ModuleRef{}
	}
	if v, ok := raw["plugins"]; !ok || v == nil {
		plain.Plugins = []PluginRef{}
	}
	if v, ok := raw["repos"]; !ok || v == nil {
		plain.Repos = []string{}
	}
	if v, ok := raw["schemaVersion"]; !ok || v == nil {
		plain.SchemaVersion = "1.0.0"
	}
	*j = RegistryJson(plain)
	return nil
}
