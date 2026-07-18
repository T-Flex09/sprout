package engine

type RGB struct {
	R, G, B uint8
}

type Cell struct {
	Char rune
	Color RGB
}

func (c *Cell) Reset() {
	c.Char = ' '
	c.Color = RGB{R: 255, G: 255, B: 255}
}
