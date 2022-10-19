package versionReaders

import "fmt"

type VersionReader interface {
	GetName() string
	GetArgumentDefinitions() []argumentDefinition
	ReadVersion(*ArgumentValues) (string, error)
}

type versionReader struct {
}

func (r *versionReader) GetArgumentDefinitions() []argumentDefinition {
	return []argumentDefinition{}
}

func (r *versionReader) ReadVersion(arguments *ArgumentValues) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func (r *versionReader) GetName() string {
	panic("not implemented")
}
