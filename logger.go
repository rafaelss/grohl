package scrolls

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Logger struct {
	stream  io.Writer
	context map[string]interface{}
}

type LogData map[string]interface{}

func NewLogger(stream io.Writer) *Logger {
	return NewLoggerWithContext(stream, make(map[string]interface{}))
}

func NewLoggerWithContext(stream io.Writer, context map[string]interface{}) *Logger {
	if stream == nil {
		stream = os.Stdout
	}
	return &Logger{stream, context}
}

func (l *Logger) Log(data map[string]interface{}) {
	l.stream.Write([]byte(l.buildLine(data)))
}

func (l *Logger) NewContext(data map[string]interface{}) *Logger {
	return NewLoggerWithContext(l.stream, dupeMaps(l.context, data))
}

func (l *Logger) AddContext(key string, value interface{}) {
	l.context[key] = value
}

func (l *Logger) DeleteContext(key string) {
	delete(l.context, key)
}

func (l *Logger) buildLine(data map[string]interface{}) string {
	merged := dupeMaps(l.context, data)
	pieces := make([]string, len(merged))

	index := 0
	for key, value := range merged {
		pieces[index] = fmt.Sprintf("%s=%s", key, value)
		index = index + 1
	}

	return strings.Join(pieces, space)
}

func dupeMaps(maps ...map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, orig := range maps {
		for key, value := range orig {
			merged[key] = value
		}
	}
	return merged
}

const (
	space = " "
	empty = ""
)