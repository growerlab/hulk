package app

import (
	"fmt"
	"os"

	"github.com/growerlab/hulk/repo"
	"github.com/pkg/errors"
)

type Action string

const (
	ActionCreated Action = "created" // create branch or tag
	ActionRemoved Action = "removed" // remove branch or tag
	ActionPushed  Action = "pushed"  // push commit
)

type RefType string

const (
	RefTypeBranch RefType = "branch"
	RefTypeTag    RefType = "tag"
)

type PushSession struct {
	After  string // old
	Before string // new
	Ref    string // branch, tag

	RepoDir string

	RepoOwner string // namespace.path
	RepoPath  string // repository name

	Action  Action
	RefType RefType
}

func (r *PushSession) IsNullOldCommit() bool {
	return r.After == repo.NullRef
}

func (r *PushSession) IsNullNewCommit() bool {
	return r.Before == repo.NullRef
}

func (r *PushSession) IsNewBranch() bool {
	return repo.IsBranch(r.Ref) && r.IsNullOldCommit() && !r.IsNullNewCommit()
}

func (r *PushSession) IsNewTag() bool {
	return repo.IsTag(r.Ref) && r.IsNullOldCommit() && !r.IsNullNewCommit()
}

func (r *PushSession) IsCommitPush() bool {
	return !r.IsNullOldCommit() && !r.IsNullNewCommit()
}

func (r *PushSession) prepare() error {
	r.RepoOwner = RepoOwner
	r.RepoPath = RepoPath

	if r.IsCommitPush() {
		r.Action = ActionPushed
	} else {
		if repo.IsBranch(r.Ref) {
			r.RefType = RefTypeBranch
			if r.IsNewBranch() {
				r.Action = ActionCreated
			} else {
				r.Action = ActionRemoved
			}
		} else if repo.IsTag(r.Ref) {
			r.RefType = RefTypeTag
			if r.IsNewTag() {
				r.Action = ActionCreated
			} else {
				r.Action = ActionRemoved
			}
		} else {
			return errors.Errorf("invalid ref '%s'", r.Ref)
		}
	}
	return nil
}

func Context() *PushSession {
	pwd, err := os.Getwd()
	if err != nil {
		ErrPanic(err)
	}

	ctx := &PushSession{
		RepoDir: pwd,
	}
	_, err = fmt.Scan(&ctx.After)
	ErrPanic(err)
	_, err = fmt.Scan(&ctx.Before)
	ErrPanic(err)
	_, err = fmt.Scan(&ctx.Ref)
	ErrPanic(err)

	if err := ctx.prepare(); err != nil {
		ErrPanic(err)
	}
	return ctx
}

func ErrPanic(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}
