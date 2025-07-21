package corr

import (
	"fmt"
	"testing"
)

// https://en.wikipedia.org/wiki/ANSI_escape_code
func TestColoring(t *testing.T) {
	msg := "an old falcon"

	reset := "\033[0m"
	bold := "\033[1m"
	underline := "\033[4m"
	strike := "\033[9m"
	italic := "\033[3m"

	cRed := "\033[31m"
	cGreen := "\033[32m"
	cYellow := "\033[33m"
	cBlue := "\033[34m"
	cPurple := "\033[35m"
	cCyan := "\033[36m"
	cWhite := "\033[37m"

	fmt.Println(msg)

	fmt.Println(cRed + msg)
	fmt.Println(cGreen + msg)
	fmt.Println(cYellow + msg)
	fmt.Println(cBlue + msg)
	fmt.Println(cPurple + msg)
	fmt.Println(cWhite + msg)
	fmt.Println(cCyan + msg + reset)

	fmt.Println(bold + msg)
	fmt.Println(italic + msg + reset)
	fmt.Println(strike + msg + reset)
	fmt.Println(underline + msg + reset)
	fmt.Println(msg)
}
