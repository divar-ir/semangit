package cli

import (
	"flag"
	"semangit/src/gitRepoManager"
	"semangit/src/versionReaders"
)

type cli struct {
	repoDir                  string
	fromRevision             string
	toRevision               string
	versionReaderName        string
	versionReadersFlagValues map[string]*versionReaders.ArgumentValues
}

func RunNewCli() cli {
	c := cli{
		fromRevision:             gitRepoManager.RevisionNone,
		toRevision:               gitRepoManager.RevisionNone,
		versionReadersFlagValues: make(map[string]*versionReaders.ArgumentValues),
	}
	c.parseFlags()

	return c
}

func (c *cli) GetRepoDir() string {
	return c.repoDir
}

func (c *cli) GetFromRevision() string {
	return c.fromRevision
}

func (c *cli) GetToRevision() string {
	return c.toRevision
}

func (c *cli) GetVersionReaderName() string {
	return c.versionReaderName
}

func (c *cli) GetCurrentVersionReaderArguments() *versionReaders.ArgumentValues {
	return c.versionReadersFlagValues[c.versionReaderName]
}

func (c *cli) parseFlags() {
	flag.StringVar(&c.repoDir, "d", ".", "Repo root path. Defaults to current directory.")
	flag.StringVar(&c.fromRevision, "f", gitRepoManager.RevisionNone, "From revision. A git reference to get version from.")
	flag.StringVar(&c.toRevision, "t", gitRepoManager.RevisionNone, "From revision. A git reference to get version from.")
	flag.StringVar(&c.versionReaderName, "r", versionReaders.VersionReaderNameHelm, "Version reader. Defaults to 'helm'.")
	c.defineVersionReadersFlags()
	flag.Parse()

	if c.GetToRevision() == gitRepoManager.RevisionNone {
		panic("Provide TO revision (-t) to compare the version according to it.")
	}
}

func (c *cli) defineVersionReadersFlags() {
	for _, versionReader := range versionReaders.GetAllVersionReaders() {
		argNamePrefix := versionReader.GetName() + "-"
		argumentValues := versionReaders.NewArgumentValues()
		for _, arg := range versionReader.GetArgumentDefinitions() {
			value := flag.String(argNamePrefix+arg.Name, arg.DefaultValue, arg.Help)
			argumentValues[arg.Name] = value
		}
		c.versionReadersFlagValues[versionReader.GetName()] = &argumentValues
	}
}
