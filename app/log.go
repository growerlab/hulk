package app

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	"github.com/pkg/errors"
)

var (
	calldepth   = 2
	projectName = "hulk"
)

func NewLogger() *Logger {
	f, err := os.OpenFile(fmt.Sprintf("%s.log", projectName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic(errors.WithStack(err))
	}
	return &Logger{out: f}
}

type Logger struct {
	out io.Writer
}

func (l *Logger) Flush() {
	if o, ok := l.out.(io.Closer); ok {
		_ = o.Close()
	}
}

func (l *Logger) Write(b []byte) (n int, err error) {
	var sb bytes.Buffer

	_, file, line, ok := runtime.Caller(calldepth)

	sb.WriteString(fmt.Sprintf("[%s] ", projectName))
	sb.WriteString(time.Now().Format(time.RFC3339))
	sb.WriteString(" ")
	if ok {
		sb.WriteString(fmt.Sprintf("%s:%d ", file, line))
	}
	sb.Write(b)

	nb, err := sb.WriteTo(l.out)
	return int(nb), err
}
