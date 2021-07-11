package cell

import (
	"fmt"
	"time"
)

type Ion struct {
	Z  		int		// Charge (+/-)
	G		float64 // Conductance (1/ohm)
}

// Cellular Fluid
type CF struct {
    Ions 		map[string]float64	// Concentrations (mEq/L)
	Volume		float64
}

type Cell struct {
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
