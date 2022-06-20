package main

import "time"
import "tconsole"

func main() {
  c, _ := tconsole.New(nil)
  c.Println("Hello!")
  c.Message("Doing #1")
  time.Sleep(time.Second)
  c.Inc()
  c.Println("Hello again")
  c.Message("Doing #2")
  time.Sleep(time.Second)
  c.Inc()
  c.Stop()
  time.Sleep(time.Second)
}
