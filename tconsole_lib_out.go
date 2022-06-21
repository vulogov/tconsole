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
  DefaultSection.WithLevel(c.offset).Println(msg...)
}

func (c *TConsole) Message(msg string) {
  if len(c.lastMsg) > 0 {
    c.Print(c.lastMsg)
  }
  c.lastMsg = msg
  c.spinner.UpdateText(msg)
}
