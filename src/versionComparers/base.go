package versionComparers

type VersionComparer interface {
	// Compare compares two versions as left and right, and returns:
	// 0, if left equals right
	// a negative value, if left is smaller than right
	// a positive value, if left is greater than right
	Compare(left string, right string) int
}
