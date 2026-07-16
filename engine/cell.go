package engine

type Cell struct {
	Char rune
	Color string
}

func (c *Cell) Reset() {
	c.Char = ' '
	c.Color = ""
}
