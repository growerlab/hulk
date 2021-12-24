package app

// EventDispatcher backend响应这个事件
type EventDispatcher interface {
	// Dispatch 将event推送给redis的stream
	Dispatch(event *PushEvent) error
}

var _ EventDispatcher = (*EventDispatch)(nil)

type EventDispatch struct {
}

func (e *EventDispatch) Dispatch(event *PushEvent) error {
	panic("implement me")
}
