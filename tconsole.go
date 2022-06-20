package tconsole

import (
  "github.com/lrita/cmap"
  "github.com/pterm/pterm"
)

type TConsole struct {
  cfg        *cmap.Cmap
  lastMsg     string
  offset      int
  hadAngle    bool
  spinner     *pterm.SpinnerPrinter
}

func New(cfg *cmap.Cmap) (*TConsole, error) {
  var err error

  res := new(TConsole)
  if cfg != nil {
    res.cfg = cfg
  } else {
    res.cfg = new(cmap.Cmap)
  }
  res.Enable()
  res.EnableColor()
  res.lastMsg  = ""
  if res.Get("console.Msg", true).(bool) {
    res.spinner, err = pterm.DefaultSpinner.WithRemoveWhenDone(true).Start()
  }
  if err != nil {
    return nil, err
  }
  res.hadAngle = false
  return res, nil
}

func (c *TConsole) Enable() {
  pterm.EnableOutput()
  pterm.RecalculateTerminalSize()
  c.Set("console.Enable", true)
}

func (c *TConsole) Disable() {
  pterm.DisableOutput()
  c.Set("console.Enable", false)
}

func (c *TConsole) EnableColor() {
  pterm.EnableColor()
  pterm.EnableStyling()
  c.Set("console.Color", true)
}

func (c *TConsole) DisableColor() {
  pterm.DisableColor()
  pterm.DisableStyling()
  c.Set("console.Color", true)
}

func (c *TConsole) ResetOffset() {
  c.offset = 0
  c.hadAngle = false
}

func (c *TConsole) Inc() {
  c.offset += 1
  c.hadAngle = false
}

func (c *TConsole) Dec() {
  c.hadAngle = false
  if c.offset == 0 {
    return
  }
  c.offset -= 1
}

func (c *TConsole) Stop(msg ...interface{}) {
  if ! c.Get("console.Msg", true).(bool) {
    return
  }
  if len(c.lastMsg) > 0 {
    c.Println(c.lastMsg)
  }
  c.spinner.Stop()
}
