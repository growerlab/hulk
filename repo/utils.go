package repo

import (
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func IsBranch(ref string) bool {
	return plumbing.ReferenceName(ref).IsBranch()
}

func IsTag(ref string) bool {
	return plumbing.ReferenceName(ref).IsTag()
}
