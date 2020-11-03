package repo

import (
	"github.com/pkg/errors"
	"gopkg.in/src-d/go-git.v4"
)

const (
	NullRef = "0000000000000000000000000000000000000000"
)

type Repository struct {
	repoPath string
	repo     *git.Repository
}

func NewRepository(repoPath string) *Repository {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		panic(errors.WithStack(err))
	}

	return &Repository{
		repoPath: repoPath,
		repo:     repo,
	}
}
