package tconsole

import (
	"testing"
)

func Test_tconsole_log1(t *testing.T) {
	c, _ := New(nil)
	c.Info("Info message")
}
