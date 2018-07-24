package log

import (
	"testing"
)

func TestLog(t *testing.T) {
}

func TestPrint(t *testing.T) {
	var format = "%s %d %s %t"
	var result = mid(format, "qaz", 54, "wer", true)
	if result != "qaz 54 wer true" {
		t.Error("Log Format Error")
	}
}

func mid(format string, args ...interface{}) string {
	return Print(format, args)
}
