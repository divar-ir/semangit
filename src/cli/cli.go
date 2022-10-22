package cli

import (
	"flag"
	"semangit/src/gitRepoManager"
	"semangit/src/versionAnalyzers"
)

type cli struct {
	repoDir                        string
	fromRevision                   string
	toRevision                     string
	versionAnalyzerName            string
	versionAnalyzersArgumentValues map[string]*versionAnalyzers.ArgumentValues
}

func RunNewCli() cli {
	c := cli{
		fromRevision:                   gitRepoManager.RevisionNone,
		toRevision:                     gitRepoManager.RevisionNone,
		versionAnalyzersArgumentValues: make(map[string]*versionAnalyzers.ArgumentValues),
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

func (c *cli) GetVersionAnalyzerName() string {
	return c.versionAnalyzerName
}

func (c *cli) GetCurrentVersionAnalyzerArgumentValues() *versionAnalyzers.ArgumentValues {
	return c.versionAnalyzersArgumentValues[c.versionAnalyzerName]
}

func (c *cli) parseFlags() {
	flag.StringVar(&c.repoDir, "d", ".", "Repo root path. Defaults to current directory.")
	flag.StringVar(&c.fromRevision, "f", gitRepoManager.RevisionNone, "From revision. A git reference to get version from.")
	flag.StringVar(&c.toRevision, "t", gitRepoManager.RevisionNone, "From revision. A git reference to get version from.")
	flag.StringVar(&c.versionAnalyzerName, "r", versionAnalyzers.VersionAnalyzerNameHelm, "Version analyzer. Defaults to 'helm'.")
	c.defineVersionAnalyzersFlags()
	flag.Parse()

	if c.GetToRevision() == gitRepoManager.RevisionNone {
		panic("Provide TO revision (-t) to compare the version according to it.")
	}
}

func (c *cli) defineVersionAnalyzersFlags() {
	for _, versionAnalyzer := range versionAnalyzers.GetAllAnalyzers() {
		argNamePrefix := versionAnalyzer.GetName() + "-"
		argumentValues := versionAnalyzers.NewArgumentValues()
		for _, argumentDefinition := range versionAnalyzer.GetExtraArgumentDefinitions() {
			argumentValue := flag.String(argNamePrefix+argumentDefinition.Name, argumentDefinition.DefaultValue, argumentDefinition.Description)
			argumentValues[argumentDefinition.Name] = argumentValue
		}
		c.versionAnalyzersArgumentValues[versionAnalyzer.GetName()] = &argumentValues
	}
}
