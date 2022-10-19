package versionReaders

var helmVersionReader = newHelmVersionReader()
var versionReaders = []VersionReader{
	&helmVersionReader,
}

func GetVersionReader(name string) VersionReader {
	for _, r := range versionReaders {
		if r.GetName() == name {
			return r
		}
	}
	panic("unknown version reader: " + name)
}

func GetAllVersionReaders() []VersionReader {
	return versionReaders
}
