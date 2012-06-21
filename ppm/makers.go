package ppm

const maxValue = 255

func MakeImage(w, h, realCols, realRows int, f (func(i, j int) Color)) *Image {
	img := new(Image)
	img.Width = w
	img.Height = h
	img.MaxValue = maxValue
	var data Data
	var row Row
	for i := 0; i < realRows; i++ {
		row = Row{}
		for j := 0; j < realCols; j++ {
			color := f(i, j)
			row = append(row, color)
		}
		data = append(data, row)
	}
	img.Data = data
	return img
}

func MakeValidImage(w, h int, f (func(i, j int) Color)) *Image {
	return MakeImage(w, h, w, h, f)
}

func MakeSolidImage(w, h, realCols, realRows int, color Color) *Image {
	fun := func(i, j int) Color { return color }
	return MakeImage(w, h, realCols, realRows, fun)
}

func MakeValidSolidImage(w, h int, color Color) *Image {
	return MakeSolidImage(w, h, w, h, color)
}

func MakeBlackImage(w, h, realCols, realRows int) *Image {
	return MakeSolidImage(w, h, realCols, realRows, black)
}

func MakeValidBlackImage(w, h int) *Image {
	return MakeBlackImage(w, h, w, h)
}
