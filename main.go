package main

import (
	"time"
	"sprout/engine"
)
const fps int = 20
func main() {
	buffer := engine.NewBuffer(0, 0)
	for {
		engine.RenderFrame(&buffer)
		var delayMS time.Duration = time.Duration(1000 / fps)
		time.Sleep(delayMS * time.Millisecond)
	}
}
