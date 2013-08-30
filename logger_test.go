package scrolls

import (
	"bytes"
	"testing"
)

func TestLogData(t *testing.T) {
	logger, buf := loggerWithBuffer()
	logger.Log(LogData{
		"a": "1", "b": "2",
	})

	if result := buf.String(); result != "a=1 b=2" {
		t.Errorf("Bad log output: %s", result)
	}
}

func TestContextObject(t *testing.T) {
	logger, buf := loggerWithBuffer()
	logger.AddContext("a", "1")
	logger.AddContext("b", "1")
	context := logger.NewContext(LogData{
		"b": "2", "c": "2",
	})

	context.Log(LogData{
		"c": "3", "d": "4",
	})

	if result := buf.String(); result != "a=1 b=2 c=3 d=4" {
		t.Errorf("Bad log output: %s", result)
	}
}

func TestLogDataWithContext(t *testing.T) {
	logger, buf := loggerWithBuffer()
	logger.AddContext("a", "1")
	logger.AddContext("b", "1")

	logger.Log(LogData{
		"b": "2", "c": "3",
	})

	if result := buf.String(); result != "a=1 b=2 c=3" {
		t.Errorf("Bad log output: %s", result)
	}
}

func TestContextDelete(t *testing.T) {
	logger, buf := loggerWithBuffer()
	logger.AddContext("a", "1")
	logger.AddContext("b", "1")
	logger.DeleteContext("b")

	logger.Log(nil)

	if result := buf.String(); result != "a=1" {
		t.Errorf("Bad log output: %s", result)
	}
}

func TestLogEmptyData(t *testing.T) {
	logger, buf := loggerWithBuffer()
	logger.Log(nil)

	if result := buf.String(); result != "" {
		t.Errorf("Bad log output: %s", result)
	}
}

func loggerWithBuffer() (*Logger, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	return NewLogger(buf), buf
}