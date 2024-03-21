package instrument

import (
	"errors"
)

var (
	ErrAlreadyInstrumented     = errors.New("pod is already instrumented")
	ErrInvalidAnnotationCont   = errors.New("duplicated target container annotation setting exists")
	ErrInvalidAnnotationValue  = errors.New("invalid annotation value")
	ErrInstrumentationNotFound = errors.New("instrumentation not found")
)
