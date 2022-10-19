package cli

import (
	"flag"
	"semangit/src/gitRepoManager"
)

type cli struct {
	repoDir      string
	fromRevision string
	toRevision   string
}

func RunNewCli() cli {
	c := cli{
		fromRevision: gitRepoManager.RevisionNone,
		toRevision:   gitRepoManager.RevisionNone,
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

func (c *cli) parseFlags() {
	flag.StringVar(&c.repoDir, "d", ".", "Repo root path. Defaults to current directory.")
	flag.StringVar(&c.fromRevision, "f", gitRepoManager.RevisionNone, "From revision. A git reference to get version from.")
	flag.StringVar(&c.toRevision, "t", gitRepoManager.RevisionNone, "From revision. A git reference to get version from.")
	flag.Parse()

	if c.GetToRevision() == gitRepoManager.RevisionNone {
		panic("Provide TO revision (-t) to compare the version according to it.")
	}
}
