package main

import (
	. "biology/cell"

	"fmt"
)

func main() {
    // Approximate ICF Composition Values (mEq/L)
    var icf = make(map[string]float64)
    icf["Na"] = 14.0
    icf["K"] = 120.0
    icf["Ca"] = 1e-4
    icf["Cl"] = 10.0
    icf["HCO3"] = 10.0
    icf["H"] = 40e-6

    // Approximate ECF Composition Values (mEq/L)
    var ecf = make(map[string]float64)
    ecf["Na"] = 140.0
    ecf["K"] = 4.0
    ecf["Ca"] = 2.5
    ecf["Cl"] = 105.0
    ecf["HCO3"] = 24.0
    ecf["H"] = 80e-6

    var ions = map[string]Ion{} 
    ions["Na"] = Ion{Z: 1, G: 0.05}
    ions["K"] = Ion{Z: 1, G: 1.0}
    ions["Ca"] = Ion{Z: 2, G: 0.05}
    ions["Cl"] = Ion{Z: -1, G: 0.45}
    ions["HCO3"] = Ion{Z: -1}
    ions["H"] = Ion{Z: +1}

    v := 90e-15 // Red blood cell (L)
	c := Cell{
        Membrane: Membrane{
            Area: 1,
            Thickness: 1,
            Radius: 1,
            Viscosity: 1,
            ICF: CF{
                Ions: icf,
                Volume: v,
            },
            ECF: CF{
                Ions: ecf,
                Volume: v,
            },
            Ions: ions,
        },
    }

	c.Init()

	fmt.Println(c)
}
