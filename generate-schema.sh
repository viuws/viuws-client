#!/usr/bin/env bash

# requires go-jsonschema v0.13.0 from https://github.com/omissis/go-jsonschema 

rm -rf internal/schema
mkdir -p internal/schema
for file in "$@"; do
    go-jsonschema -o "internal/schema/$(basename "$file" .json).go" -p schema --tags json "$file"
done

sed -i 's/plain\.Inputs = map\[string\]interface{}{}/plain\.Inputs = ModuleJsonInputs{}/g' internal/schema/module.go
sed -i 's/plain\.Outputs = map\[string\]interface{}{}/plain\.Outputs = ModuleJsonOutputs{}/g' internal/schema/module.go

sed -i 's/plain\.Environments = map\[string\]interface{}{}/plain\.Environments = WorkflowJsonEnvironments{}/g' internal/schema/workflow.go
sed -i 's/plain\.Tasks = map\[string\]interface{}{}/plain\.Tasks = WorkflowJsonTasks{}/g' internal/schema/workflow.go
