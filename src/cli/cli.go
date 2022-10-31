package cli

import (
	"flag"
	"semangit/src/gitrepo"
	"semangit/src/versionanalyzers"
)

type cli struct {
	repoDir                        string
	fromRevision                   string
	toRevision                     string
	versionAnalyzerName            string
	versionAnalyzersArgumentValues map[string]*versionanalyzers.ArgumentValues
}

func RunNewCli() cli {
	c := cli{
		fromRevision:                   gitrepo.RevisionNone,
		toRevision:                     gitrepo.RevisionNone,
		versionAnalyzersArgumentValues: make(map[string]*versionanalyzers.ArgumentValues),
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

func (c *cli) GetCurrentVersionAnalyzerArgumentValues() *versionanalyzers.ArgumentValues {
	return c.versionAnalyzersArgumentValues[c.versionAnalyzerName]
}

func (c *cli) parseFlags() {
	flag.StringVar(&c.repoDir, "d", ".", "Repo root path. Defaults to current directory.")
	flag.StringVar(&c.fromRevision, "f", gitrepo.RevisionNone, "From revision. A git reference to get version from.")
	flag.StringVar(&c.toRevision, "t", gitrepo.RevisionNone, "From revision. A git reference to get version from.")
	flag.StringVar(&c.versionAnalyzerName, "r", versionanalyzers.VersionAnalyzerNameHelm, "Version analyzer. Defaults to 'helm'.")
	c.defineVersionAnalyzersFlags()
	flag.Parse()

	if c.GetToRevision() == gitrepo.RevisionNone {
		panic("Provide TO revision (-t) to compare the version according to it.")
	}
}

func (c *cli) defineVersionAnalyzersFlags() {
	for _, versionAnalyzer := range versionanalyzers.GetAllAnalyzers() {
		argNamePrefix := versionAnalyzer.GetName() + "-"
		argumentValues := versionanalyzers.NewArgumentValues()
		for _, argumentDefinition := range versionAnalyzer.GetExtraArgumentDefinitions() {
			argumentValue := flag.String(argNamePrefix+argumentDefinition.Name, argumentDefinition.DefaultValue, argumentDefinition.Description)
			argumentValues[argumentDefinition.Name] = argumentValue
		}
		c.versionAnalyzersArgumentValues[versionAnalyzer.GetName()] = &argumentValues
	}
}
