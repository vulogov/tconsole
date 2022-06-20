package tconsole

import (
  "github.com/pterm/pterm"
)

var DefaultStyle = pterm.Style{pterm.FgDefault}

var DefaultSection = pterm.SectionPrinter{
	Style:           &DefaultStyle,
	Level:           0,
	TopPadding:      0,
	BottomPadding:   0,
	IndentCharacter: " ",
}

func (c *TConsole) Println(msg ...interface{}) {
  if c.offset == 0 {
    DefaultSection.WithLevel(0).Println(msg...)
  } else {
    fmsg := make([]interface{}, 0)
    fmsg = append(fmsg, "└")
    fmsg = append(fmsg, msg...)
    DefaultSection.WithLevel(c.offset-1).Println(fmsg...)
  }
}

func (c *TConsole) Message(msg string) {
  if len(c.lastMsg) > 0 {
    c.Println(c.lastMsg)
  }
  c.lastMsg = msg
  c.spinner.UpdateText(msg)
}
