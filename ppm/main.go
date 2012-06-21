package ppm

import (
	"fmt"
	"io/ioutil"
)

type Color [3]int
type Row []Color
type Data []Row

var black = Color{0, 0, 0}

type Image struct {
	Width    int
	Height   int
	MaxValue int
	Data     Data
}

func (i *Image) IsValid() bool {
	if len(i.Data) != i.Height {
		return false
	}
	for _, row := range i.Data {
		if len(row) != i.Width {
			return false
		}
		for _, color := range row {
			for _, value := range color {
				if value > i.MaxValue {
					return false
				}
			}
		}
	}
	return true
}

func (d1 Data) Equals(d2 Data) bool {
	if len(d1) != len(d2) {
		return false
	}
	var width1, width2 int
	for i, row := range d1 {
		r1, r2 := row, d2[i]
		width1, width2 = len(r1), len(r2)
		if width1 != width2 {
			return false
		}
		for j, color := range r1 {
			if color != r2[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	content, err := ioutil.ReadFile("test.ppm")
	if err != nil {
		fmt.Printf("Error with reading file: %s", err)
	}
	println(string(content))
}
