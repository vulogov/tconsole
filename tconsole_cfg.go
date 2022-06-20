package tconsole

import (
  "fmt"
)

func (c *TConsole) Set(key string, value interface{}) {
  c.cfg.Store(key, value)
}

func (c *TConsole) Get(key string, value interface{}) interface{} {
  res, _ := c.cfg.LoadOrStore(key, value)
  return res
}

func (c *TConsole) isOutput(key string) bool {
  if ! c.Get("console.Enable", true).(bool) {
    return false
  }
  if ! c.Get(fmt.Sprintf("console.%v", key), true).(bool) {
    return false
  }
  return true
}
