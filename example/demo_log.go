package main

import "tconsole"

func main() {
  c, _ := tconsole.New(nil)
  c.Info("Info message")
  c.Inc()
  c.Info("Info message with increment")
  c.Warning("Warning message")
  c.Error("Error message")
  c.Inc()
  c.Error("Error message with more increments")
  c.Warning("Warning message with more increments")
  c.Info("Info message with another increment")
  c.Inc()
  c.Info("Info message with more and more  increment")
  c.Inc()
  c.Info("Info message with more and more and more  increment")
  c.ResetOffset()
  c.Critical("Critical message")
  c.Stop()
}
