package log

import (
	"testing"
)

func TestSimple(t *testing.T) {
	l := New()
	l.Warn("Hello world")
	l.Warnf("Hello %s", "srid")
	l.Print("xyz")
	// Fatal("oops")
}
