package tconsole

import (
  "fmt"
  "bufio"
  "strings"
)

func padText(s string, n int) string {
  fstr := " %-" + fmt.Sprintf("%v", n) + "v"
  return fmt.Sprintf(fstr, s)
}

func splitLines(s string) []string {
    var lines []string
    sc := bufio.NewScanner(strings.NewReader(s))
    for sc.Scan() {
        lines = append(lines, sc.Text())
    }
    return lines
}
