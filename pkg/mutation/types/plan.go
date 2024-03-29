package types

import (
	v1 "github.com/DDnK-dev/apm-instrumentaion-operator/api/v1"
)

// Plan is a plan for mutation
// it should contain the list of containers to be mutated and the instrumentation to be applied
type Plan struct {
	Containers      []string
	Instrumentation *v1.Instrumentation
}
