package versionComparers

import "golang.org/x/mod/semver"

type SemanticVersionComparer struct {
}

func (c SemanticVersionComparer) Compare(left string, right string) int {
	return semver.Compare(left, right)
}
