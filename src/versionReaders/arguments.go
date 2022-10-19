package versionReaders

type argumentDefinition struct {
	Name         string
	DefaultValue string
	Help         string
}

type ArgumentValues map[string]*string

func NewArgumentValues() ArgumentValues {
	return make(ArgumentValues)
}
