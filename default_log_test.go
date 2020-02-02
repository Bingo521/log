package log

import (
	"testing"
)

func TestLog(t *testing.T) {
	Info("info")
	logs.Close()
}