package app

import (
	"github.com/chanxuehong/util/math"
)

var _ Hook = (*HookEvent)(nil)

type PushEvent struct {
	*PushSession
	CommitCount   int    `json:"commit_count"`
	RefCount      int    `json:"ref_count"`
	CommitMessage string `json:"commit_message"` // commit message
}

// 创建推送事件
type HookEvent struct {
}

func (h *HookEvent) Label() string {
	return "event"
}

func (h *HookEvent) Priority() int {
	return math.MaxInt
}

func (h *HookEvent) Process(sess *PushSession) error {
	panic("implement me")
}
