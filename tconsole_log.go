package tconsole

import (
  "fmt"
  "github.com/pterm/pterm"
)

var CriticalPrefixStyle = pterm.Style{pterm.FgLightYellow, pterm.BgLightRed}
var DebugPrefixStyle = pterm.Style{pterm.FgBlack, pterm.BgGray}
var InfoPrefixStyle = pterm.Style{pterm.FgBlack, pterm.BgCyan}
var WarningPrefixStyle = pterm.Style{pterm.FgBlack, pterm.BgYellow}
var ErrorPrefixStyle = pterm.Style{pterm.FgBlack, pterm.BgLightRed}
var ErrorStyle = pterm.Style{pterm.FgLightRed}
var FunctionPrefixStyle = pterm.Style{pterm.FgLightWhite, pterm.BgGray}

const PSIZE=10

func (c *TConsole) Print(msg ...interface{}) {
  if c.isOutput("Print") {
    c.lastDbg  = ""
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.Sprint(padText(" ", PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Func(fname string, msg ...interface{}) {
  if c.isOutput("Function") {
    c.lastDbg  = ""
    c.Inc()
    c.fq.PushBack(fname)
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&FunctionPrefixStyle).Sprint(padText(fmt.Sprintf("⨍(%v) =", fname), PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Fclose(msg ...interface{}) {
  if c.isOutput("Function") && c.fq.Len() > 0 {
    c.lastDbg  = ""
    fname := c.fq.PopBack().(string)
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&FunctionPrefixStyle).Sprint(padText(fmt.Sprintf("⨍(%v) =", fname), PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
    c.Dec()
  }
}

func (c *TConsole) F(msg string, args ...interface{}) {
  if c.isOutput("Function") && c.fq.Len() > 0 {
    c.lastDbg  = ""
    fname := c.fq.Back().(string)
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&FunctionPrefixStyle).Sprint(padText(fmt.Sprintf("⨍(%v) =", fname), PSIZE)))
    fmsg = append(fmsg, fmt.Sprintf(msg, args...))
    c.Println(fmsg...)
  }
}

func (c *TConsole) Section(sname string, msg ...interface{}) {
  if c.isOutput("Debug") {
    c.lastDbg  = ""
    c.Inc()
    c.sq.PushBack(sname)
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&FunctionPrefixStyle).Sprint(padText(fmt.Sprintf("⊂(%v)", sname), PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) SectionEnd(msg ...interface{}) {
  if c.isOutput("Debug") && c.sq.Len() > 0 {
    c.lastDbg  = ""
    sname := c.sq.PopBack().(string)
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&FunctionPrefixStyle).Sprint(padText(fmt.Sprintf("⊄(%v)", sname), PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
    c.Dec()
  }
}

func (c *TConsole) Debug(msg ...interface{}) {
  var opfx string

  if c.isOutput("Debug") {
    fmsg := make([]interface{}, 0)
    if c.sq.Len() == 0 {
      pfx := padText("DEBUG", PSIZE)
      opfx = pfx
      if c.lastDbg == pfx {
        opfx = padText(" ", PSIZE)
      } else {
        c.lastDbg = pfx
      }
    } else {
      opfx = padText(" ", PSIZE)
    }
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&DebugPrefixStyle).Sprint(opfx))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Info(msg ...interface{}) {
  if c.isOutput("Info") {
    c.lastDbg  = ""
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&InfoPrefixStyle).Sprint(padText("INFO", PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Warning(msg ...interface{}) {
  if c.isOutput("Warning") {
    c.lastDbg  = ""
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&WarningPrefixStyle).Sprint(padText("WARNING", PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) IfError(err error, msg ...interface{}) {
  if c.isOutput("Error") && err != nil {
    c.lastDbg  = ""
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&ErrorPrefixStyle).Sprint(padText("ERROR", PSIZE)))
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&ErrorStyle).Sprint(fmt.Sprintf("\"%v\"", err)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Error(msg ...interface{}) {
  if c.isOutput("Error") {
    c.lastDbg  = ""
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&ErrorPrefixStyle).Sprint(padText("ERROR", PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}

func (c *TConsole) Critical(msg ...interface{}) {
  if c.isOutput("Critical") {
    c.lastDbg  = ""
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, pterm.DefaultBasicText.WithStyle(&CriticalPrefixStyle).Sprint(padText("CRITICAL", PSIZE)))
    fmsg = append(fmsg, msg...)
    c.Println(fmsg...)
  }
}
