package chainofresponsibility

import (
	"fmt"
	"time"
)

type LOG_LEVEL int

const (
	DEBUG LOG_LEVEL = iota + 1
	INFO
	ERROR
)

func (l LOG_LEVEL) EnumIndex() int {
	return int(l)
}

type ILogProcessor interface {
	Log(LOG_LEVEL, string)
}
type InfoLogger struct{
	Base ILogProcessor
}

func (i *InfoLogger) Log(level LOG_LEVEL, msg string) {
	if level == INFO {
		fmt.Printf("INFO %s : Message : %s\n", time.Now().Format("2016-01-01 10:20:20"), msg)
	} else {
		i.Base.Log(level, msg)
	}
}

type DebugLogger struct{
	Base ILogProcessor
}

func (d *DebugLogger) Log(level LOG_LEVEL, msg string) {
	if level == DEBUG {
		fmt.Printf("DEBUG %s : Message : %s\n", time.Now().Format("2016-01-01 10:20:20"), msg)
	} else {
		d.Base.Log(level, msg)
	}
}

type ErrorLogger struct{
	Base ILogProcessor
}

func (e *ErrorLogger) Log(level LOG_LEVEL, msg string) {
	if level == ERROR {
		fmt.Printf("ERROR %s : Message : %s\n", time.Now().Format("2016-01-01 10:20:20"), msg)
	} else {
		e.Base.Log(level, msg)
	}
}

type NOLogger struct{
}

func (n *NOLogger) Log(level LOG_LEVEL, msg string) {
	fmt.Printf("Cant log as level %d not allowed \n", level)
}

func NewInfoLogger(base ILogProcessor) *InfoLogger {
	return &InfoLogger{
		Base: base,
	}
}
func NewDebugLogger(base ILogProcessor) *DebugLogger {
	return &DebugLogger{
		Base: base,
	}
}
func NewErrorLogger(base ILogProcessor) *ErrorLogger {
	return &ErrorLogger{
		Base: base,
	}
}
