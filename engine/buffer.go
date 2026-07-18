package engine

import (
	"strings"
)

type Buffer struct {
	Data [][]Cell
	Width, Height int
}

func NewBuffer(width, height int) Buffer {
	data := make([][]Cell, height)
	for i := range data {
		data[i] = make([]Cell, width)
		for j := range data[i] {
			data[i][j] = Cell{
				Char: ' ',
				Color: RGB{R: 255, G: 255, B: 255},
			}
		}
	}

	return Buffer {
		Data: data,
		Width: width,
		Height: height,
	}
}

func (buff Buffer) ToFrame() string {
	var sb strings.Builder
	const SCREEN_CLEARING_BYTES int = 7
	
	size := (buff.Width * buff.Height * 23) + buff.Height + SCREEN_CLEARING_BYTES
	sb.Grow(size)

	sb.WriteString("\x1b[H")
	
	for row := 0; row < buff.Height; row++ {
		for col := 0; col < buff.Width; col++ {
			cell := buff.Data[row][col]
			
			sb.WriteString("\x1b[38;2;")
			writeUint8(&sb, cell.Color.R)
			sb.WriteByte(';')
			writeUint8(&sb, cell.Color.G)
			sb.WriteByte(';')
			writeUint8(&sb, cell.Color.B)
			sb.WriteByte('m')
			sb.WriteRune(cell.Char)
		}
		if row != buff.Height - 1 {
			sb.WriteByte('\n')
		}
	}
	
	sb.WriteString("\x1b[0m")
	
	return sb.String()
}

func writeUint8(sb *strings.Builder, val uint8) {
	if val >= 100 {
		sb.WriteByte('0' + val/100)
		sb.WriteByte('0' + (val%100)/10)
		sb.WriteByte('0' + val%10)
	} else if val >= 10 {
		sb.WriteByte('0' + val/10)
		sb.WriteByte('0' + val%10)
	} else {
		sb.WriteByte('0' + val)
	}
}
