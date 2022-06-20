package tconsole

import (
  "github.com/pterm/pterm"
)

type TConsole struct {
  lastMsg     string
  offset      int
  spinner     *pterm.SpinnerPrinter
}

func New() (*TConsole, error) {
  var err error
  res := new(TConsole)
  res.Enable()
  res.EnableColor()
  res.lastMsg  = ""
  res.spinner, err = pterm.DefaultSpinner.WithRemoveWhenDone(true).Start()
  if err != nil {
    return nil, err
  }
  return res, nil
}

func (c *TConsole) Enable() {
  pterm.EnableOutput()
  pterm.RecalculateTerminalSize()
}

func (c *TConsole) Disable() {
  pterm.DisableOutput()
}

func (c *TConsole) EnableColor() {
  pterm.EnableColor()
  pterm.EnableStyling()
}

func (c *TConsole) DisableColor() {
  pterm.DisableColor()
  pterm.DisableStyling()
}

func (c *TConsole) ResetOffset() {
  c.offset = 0
}

func (c *TConsole) Inc() {
  c.offset += 1
}

func (c *TConsole) Dec() {
  if c.offset == 0 {
    return
  }
  c.offset -= 1
}

func (c *TConsole) Stop(msg ...interface{}) {
  if len(c.lastMsg) > 0 {
    c.Println(c.lastMsg)
  }
  c.spinner.Stop()
}
