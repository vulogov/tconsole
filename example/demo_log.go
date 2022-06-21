package main

import "tconsole"
import "errors"

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
  c.Print("Hello, this is indentation-aware Print")
  c.Debug("Debug message with more and more and more  increment")
  c.Func("abc", "Output about function call")
  c.Debug("Debug message #1")
  c.Debug("Debug message #2")
  c.Debug("Debug message #3")
  c.F("More info about function call")
  c.Func("cde", "Output about function call")
  c.Fclose("No more about function")
  c.IfError(errors.New("some error"), "yes, there are errors")
  c.Fclose("No more about function")
  c.ResetOffset()
  c.Critical("Critical message")
  c.Debug("Debug message #4")
  c.Stop()
}
