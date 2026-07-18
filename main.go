package main

import (
	"sprout/engine"
)

func main() {
	buffer := engine.NewBuffer(0, 0)
	for {
		engine.RenderFrame(&buffer)
	}
}
