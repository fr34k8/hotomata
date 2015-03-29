package hotomata

import (
	"fmt"
)

type Logger interface {
	Print(...interface{})
	Printf(string, ...interface{})
}

func logLine(l Logger, filler rune, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	for len(msg) < 80 {
		msg = msg + string(filler)
	}
	l.Print(msg + "\n")
}
