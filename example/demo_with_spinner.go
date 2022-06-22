package main

import "time"
import "tconsole"

func main() {
  c, _ := tconsole.New(nil)
  c.Print("Hello!")
  c.Message("Doing #1")
  time.Sleep(time.Second)
  c.Inc()
  c.Warning("Warning message")
  c.Print("Hello again")
  c.Message("Doing #2")
  time.Sleep(time.Second)
  c.Debug("Properly indented debug message")
  c.Stop()
  time.Sleep(time.Second)
}
