package app

import (
	"github.com/pkg/errors"
)

var app *App

type Hook interface {
	Label() string
	Process(sess *PushSession) error
}

type App struct {
	hooks map[string]Hook
}

func (a *App) RegisterHook(h Hook) {
	a.hooks[h.Label()] = h
}

func (a *App) Run(sess *PushSession) error {
	for _, hook := range a.hooks {
		if err := hook.Process(sess); err != nil {
			return err
		}
	}
	return nil
}

func Run(sess *PushSession) error {
	if app == nil {
		return errors.Errorf("must init App")
	}
	return app.Run(sess)
}
