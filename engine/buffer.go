package engine

type Buffer struct {
	Data [][]Cell
	Width, Height int
}

func newBuffer(width, height int) Buffer {
	data := make([][]Cell, height)
	for i := range data {
		data[i] = make([]Cell, width)
		for j := range data[i] {
			data[i][j] = Cell{
				Char: ' ',
				Color: "",
			}
		}
	}
	return Buffer {
		Data: data,
		Width: width,
		Height: height,
	}
}
