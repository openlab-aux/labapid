package logging

import "fmt"
import "io"

const (
	INFO     int = 1
	DEBUG    int = 2
	WARNING  int = 3
	ERROR    int = 4
	CRITICAL int = 5
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

func (l *Logger) Info(s string) {
	if l.loglevel <= INFO {
		fmt.Fprintf(l.writer, l.template, "INFO", s)
	}
}

func (l *Logger) Debug(s string) {
	if l.loglevel <= DEBUG {
		fmt.Fprintf(l.writer, l.template, "DEBUG", s)
	}
}

func (l *Logger) Warning(s string) {
	if l.loglevel <= WARNING {
		fmt.Fprintf(l.writer, l.template, "WARNING", s)
	}
}

func (l *Logger) Error(s string) {
	if l.loglevel <= ERROR {
		fmt.Fprintf(l.writer, l.template, "ERROR", s)
	}
}

func (l *Logger) Critical(s string) {
	if l.loglevel <= CRITICAL {
		fmt.Fprintf(l.writer, l.template, "CRITICAL", s)
	}
}
