package cell

import (
	"fmt"
	"time"
)

type Ion struct {
	Charge  		int		// (+/-)
	Concentration	float64 // mEq/L
}

// Cellular Fluid
type CF struct {
    Ions 		map[string]Ion
	Volume		float64
}

type Cell struct {
	Osmolarity	float64
	Charge		float64

	Membrane
}

func (c *Cell) Init() {
	go c.Membrane.Init()
	c.tick(1000 * time.Millisecond)
}

func (c *Cell) tick(d time.Duration) {
	for range time.Tick(d) {
		fmt.Printf("Cell tick \n")
	}
}
