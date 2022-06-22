package tconsole

import (
  "time"
  "github.com/lrita/cmap"
  "github.com/pterm/pterm"
  "github.com/gammazero/deque"
)

type TConsole struct {
  cfg        *cmap.Cmap
  lastMsg     string
  lastDbg     string
  offset      int
  spinner     *pterm.SpinnerPrinter
  fq          *deque.Deque[interface{}]
  sq          *deque.Deque[interface{}]
}

var SpinnerStyle = pterm.Style{pterm.FgLightRed}
var SpinnerTextStyle =  pterm.Style{pterm.FgDefault, pterm.BgDefault}

var DefaultSpinner = pterm.SpinnerPrinter{
	Sequence:            []string{"▀ ", " ▀", " ▄", "▄ "},
	Style:               &SpinnerStyle,
	Delay:               time.Millisecond * 200,
	ShowTimer:           true,
	TimerRoundingFactor: time.Second,
	TimerStyle:          &pterm.ThemeDefault.TimerStyle,
	MessageStyle:        &SpinnerTextStyle,
	SuccessPrinter:      &pterm.Success,
	FailPrinter:         &pterm.Error,
	WarningPrinter:      &pterm.Warning,
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
  res.lastDbg  = ""
  if res.Get("console.Msg", true).(bool) {
    res.spinner, err = DefaultSpinner.WithRemoveWhenDone(true).Start()
  }
  if err != nil {
    return nil, err
  }
  res.fq = deque.New[interface{}]()
  res.sq = deque.New[interface{}]()
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
  c.fq.Clear()
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
  if ! c.Get("console.Msg", true).(bool) {
    return
  }
  if len(c.lastMsg) > 0 {
    c.Print(c.lastMsg)
  }
  c.spinner.Stop()
}
