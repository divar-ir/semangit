package versionReaders

import "fmt"

type VersionReader interface {
	GetName() string
	GetArgumentDefinitions() []argumentDefinition
	ReadVersion(*ArgumentValues) (string, error)
}

type baseReader struct {
}

func (r *baseReader) GetArgumentDefinitions() []argumentDefinition {
	return []argumentDefinition{}
}

func (r *baseReader) ReadVersion(arguments *ArgumentValues) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (r *baseReader) GetName() string {
	panic("not implemented")
}
