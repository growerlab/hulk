package app

import (
	"sort"

	"github.com/pkg/errors"
)

var app *App

type Hook interface {
	Label() string                   // 钩子名称
	Priority() int                   // 钩子优先级,数字越小越先执行
	Process(sess *PushSession) error // 执行钩子
}

type App struct {
	hooks []Hook
}

func (a *App) RegisterHook(hooks ...Hook) {
	a.hooks = append(a.hooks, hooks...)
}

func (a *App) sortHooksByPriority() {
	sort.Slice(a.hooks, func(i, j int) bool {
		return a.hooks[i].Priority() < a.hooks[j].Priority()
	})
}

func (a *App) Run(sess *PushSession) error {
	a.sortHooksByPriority()

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
