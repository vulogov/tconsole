package tconsole

import (
	"fmt"
	"testing"
)

func Test_tconsole_lib1(t *testing.T) {
	fmt.Println("|",padText("WARNING", 10),"|")
	fmt.Println("|",padText("INFO", 10), "|")

}
