package app

import "os"

var (
	RepoOwner = os.Getenv("GROWERLAB_REPO_OWNER")
	RepoPath  = os.Getenv("GROWERLAB_REPO_NAME")
)

func init() {
	app = &App{
		hooks: map[string]Hook{},
	}
	app.RegisterHook(&HookEvent{})
}
