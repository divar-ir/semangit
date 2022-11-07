package versionanalyzers

type Config struct {
	Name           string         `json:"name,omitempty"`
	ArgumentValues ArgumentValues `json:"argument_values,omitempty"`
}
