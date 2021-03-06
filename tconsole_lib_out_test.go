package tconsole

import (
	"time"
	"testing"
)

func Test_tconsole_lib_out1(t *testing.T) {
	c, _ := New(nil)
	c.Println("Test with ident 0")
	c.Inc()
	c.Println("Test with ident 1")
	c.Inc()
	c.Println("Test with ident 2")
}

func Test_tconsole_lib_out2(t *testing.T) {
	c, _ := New(nil)
	c.Println("Test with spinner output")
	c.Message("To spinner #1")
	time.Sleep(time.Second)
	c.Stop("OK")
}
