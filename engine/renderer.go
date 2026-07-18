package engine

import (
	"os"
	"golang.org/x/term"
)

func RenderFrame(buffer *Buffer) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return
	}
	if buffer.Width != width || buffer.Height != height {
		tempBuff := NewBuffer(width, height)
		buffer = &tempBuff
	}

	os.Stdout.WriteString(buffer.ToFrame())
}
