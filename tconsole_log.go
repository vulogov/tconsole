package tconsole

import (
  "github.com/pterm/pterm"
)

var CriticalPrefixStyle = pterm.Style{pterm.FgLightYellow, pterm.BgLightRed}
var InfoPrefixStyle = pterm.Style{pterm.FgBlack, pterm.BgCyan}
var WarningPrefixStyle = pterm.Style{pterm.FgBlack, pterm.BgYellow}
var ErrorPrefixStyle = pterm.Style{pterm.FgBlack, pterm.BgLightRed}


func (c *TConsole) Info(msg ...interface{}) {
  if c.isOutput("Info") {
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&InfoPrefixStyle).Sprint(padText("INFO", 10)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Warning(msg ...interface{}) {
  if c.isOutput("Warning") {
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&WarningPrefixStyle).Sprint(padText("WARNING", 10)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Error(msg ...interface{}) {
  if c.isOutput("Warning") {
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&ErrorPrefixStyle).Sprint(padText("ERROR", 10)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Critical(msg ...interface{}) {
  if c.isOutput("Critical") {
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&CriticalPrefixStyle).Sprint(padText("CRITICAL", 10)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}
