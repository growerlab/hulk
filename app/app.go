package app

import (
	"fmt"
)

func Run(ctx *PushContext) error {
	fmt.Printf("%+v", ctx)
	return nil
}
