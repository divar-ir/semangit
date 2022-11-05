package plugins

import (
	"semangit/internal/plugins/base"
)

type Config struct {
	Name           string              `json:"name,omitempty"`
	ArgumentValues base.ArgumentValues `json:"argument_values,omitempty"`
}
