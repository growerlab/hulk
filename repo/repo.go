package repo

import (
	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
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
