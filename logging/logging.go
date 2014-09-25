package logging

import "fmt"
import "io"

const (
	DEBUG    int = iota
	INFO     int = iota
	WARNING  int = iota
	ERROR    int = iota
	CRITICAL int = iota
)

type Logger struct {
	loglevel int
	template string
	writer   io.Writer
}

func New(loglevel int, template string, writer io.Writer) *Logger {

	l := Logger{
		loglevel: loglevel,
		template: template,
		writer:   writer,
	}

	return &l
}

func (l *Logger) Debug(s string) {
	if l.loglevel >= DEBUG {
		fmt.Fprintf(l.writer, l.template, "DEBUG", s)
	}
}

func (l *Logger) Info(s string) {
	if l.loglevel >= INFO {
		fmt.Fprintf(l.writer, l.template, "INFO", s)
	}
}

func (l *Logger) Warning(s string) {
	if l.loglevel >= WARNING {
		fmt.Fprintf(l.writer, l.template, "WARNING", s)
	}
}

func (l *Logger) Error(s string) {
	if l.loglevel >= ERROR {
		fmt.Fprintf(l.writer, l.template, "ERROR", s)
	}
}

func (l *Logger) Critical(s string) {
	if l.loglevel >= CRITICAL {
		fmt.Fprintf(l.writer, l.template, "CRITICAL", s)
	}
}
