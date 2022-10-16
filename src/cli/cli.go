package cli

import "flag"

const RevisionNone = ""

type cli struct {
	repoDir      string
	fromRevision string
	toRevision   string
}

func NewCliAndRun() cli {
	c := cli{
		fromRevision: RevisionNone,
		toRevision:   RevisionNone,
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
	flag.StringVar(&c.fromRevision, "f", RevisionNone, "From revision. A git reference to get version from.")
	flag.StringVar(&c.toRevision, "t", RevisionNone, "From revision. A git reference to get version from.")
	flag.Parse()
}
