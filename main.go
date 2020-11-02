package main

import (
	"log"
	"os"

	"github.com/growerlab/hulk/app"
)

func main() {
	l := app.NewLogger()
	log.SetOutput(l)

	defer func() {
		l.Flush()
		if e := recover(); e != nil {
			log.Println(e)
			os.Exit(1)
		}
	}()

	ctx := app.Context()
	if err := app.Run(ctx); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
