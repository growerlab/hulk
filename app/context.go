package app

import (
	"fmt"
	"os"

	"github.com/growerlab/hulk/repo"
	"github.com/pkg/errors"
)

type PushContext struct {
	After  string // old
	Before string // new
	Ref    string // branch, tag

	RepoDir string
}

func (r *PushContext) IsNullOldCommit() bool {
	return r.After == repo.NullRef
}

func (r *PushContext) IsNewBranch() bool {
	return repo.IsBranch(r.Ref) && r.IsNullOldCommit()
}

func (r *PushContext) IsNewTag() bool {
	return repo.IsTag(r.Ref) && r.IsNullOldCommit()
}

func Context() *PushContext {
	pwd, err := os.Getwd()
	if err != nil {
		ErrPanic(err)
	}

	ctx := &PushContext{
		RepoDir: pwd,
	}
	_, err = fmt.Scan(&ctx.After)
	ErrPanic(err)
	_, err = fmt.Scan(&ctx.Before)
	ErrPanic(err)
	_, err = fmt.Scan(&ctx.Ref)
	ErrPanic(err)

	return ctx
}

func ErrPanic(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}
