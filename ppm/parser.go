package ppm

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func parseInt(x string) (int, error) {
	num, err := strconv.ParseInt(x, 10, 0)
	return int(num), err
}

// NOTE: width and height values are undefined if err is not nil!
func getWidthAndHeight(line string) (width, height int, err error) {
	items := strings.Split(line, string(' '))
	// FIXME: len can be less than two -> adjust error message
	//l := len(items)
	switch l := len(items); {
	case l < 2:
		return 0, 0, errors.New("missing information: height")
	case l > 2:
		return 0, 0, errors.New("only two numbers allowed")
	default:
		w, h := items[0], items[1]
		width, err = parseInt(w)
		if err != nil {
			return width, 0, errors.New(fmt.Sprintf("invalid width: %s", w))
		}
		height, err = parseInt(h)
		if err != nil {
			return 0, height, errors.New(fmt.Sprintf("invalid height: %s", h))
		}
		return width, height, nil
	}
	panic("unreachable")
}

/*
func ParsePPM(source string) (*Image, error) {
	allLines := strings.Split(source, string('\n'))
	firstRow, secondRow, lines := allLines[0], allLines[1], allLines[2:]
	width, height, err := getWidthAndHeight(firstRow)
	if err != nil {
		return new(Image), err
	}
	maxValue, err := parseInt(lines[1])
	if err != nil {
		return new(Image), err
	}
	img := new(Image)
	img.Width = width
	img.Height = height
	img.MaxValue = maxValue
		data := [][]color
		for lineno, line := range lines {
			for _, value := range strings.Split(line, ' ') {
				if value == '' {
					continue
				} else {
					i, err := parseInt(value)
					if err != nil {
						errmsg := fmt.Sprintf(
							"invalid value %s (line %d)",
							value, lineno)
						return img, errors.New(errmsg)
					} else {
						// TODO: check if current color array
						// is "full". If full -> init new color
						// array and add the current i to this
						// color array. If non-full -> add the
						// current i to this color array.
					}
				}
			}
		}
	return img, nil
}
*/
