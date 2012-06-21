package ppm

import (
	"fmt"
	"strings"
)

const IDENT = 4

func (c Color) String() (output string) {
	output = "Color("
	for i := 0; i < 2; i++ {
		output += fmt.Sprintf("%d, ", c[i])
	}
	output += fmt.Sprintf("%d)", c[2])
	return
}

func (r Row) String() (output string) {
	output = fmt.Sprintf("[%v", r[0])
	for i := 1; i < len(r); i++ {
		output += fmt.Sprintf(", %v", r[i])
	}
	output += "]"
	return
}

func (d Data) String() (output string) {
	output = "[\n"
	sep := ","
	for i, row := range d {
		// don't insert a comma after the last item
		if i == len(d)-1 {
			sep = ""
		}
		output += fmt.Sprintf("%s%v%s\n", strings.Repeat(" ", IDENT), row, sep)
	}
	output += "]"
	return
}
