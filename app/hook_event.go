package app

var _ Hook = (*HookEvent)(nil)

// 创建推送事件
type HookEvent struct {
}

func (h HookEvent) Label() string {
	return "event"
}

func (h HookEvent) Process(sess *PushSession) error {
	panic("implement me")
}
