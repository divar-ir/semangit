package versionanalyzers

type ArgumentName = string
type ArgumentValue = string

type ArgumentDefinition struct {
	Name         ArgumentName
	DefaultValue ArgumentValue
	Description  string
}

type ArgumentValues map[ArgumentName]*ArgumentValue

func NewArgumentValues() ArgumentValues {
	return make(ArgumentValues)
}
