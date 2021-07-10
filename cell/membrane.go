package cell

import (
	"math"
	"time"
	"fmt"
)

const (
	kB = 1.380649e-23 // Boltzmann constant (J⋅K^−1)
	T  = 255+37          // Body temperature 37 C (K)
	NA = 6.0221409e+23 // Avogadro's constant (mol^-1)
	R = 8.135 // Gas constant (J K^-1 mol^-1)
	F = 9.684e+4 // Faraday's constant (C mol^-1)
)

type Membrane struct {
	Area      float64 // Surface area for diffusion (cm2)
	Radius    float64 // Molecular radius
	Viscosity float64 // Viscosity of the medium
	Thickness float64 // Membrane thickness

	Ions 	[]string // Permeable Ions

	ICF		CF 	// Intracellular Fluid
	ECF		CF	// Extracellular Fluid
}

func (m *Membrane) Init() {
	m.tick(1000 * time.Millisecond)
}

func (cf *CF) Get_Mmol(ion string) float64 {
	// Convert mEq/L to mmol
	return cf.Volume * (cf.Ions[ion].Concentration * math.Abs(float64(cf.Ions[ion].Charge)))
}

func (cf *CF) Set_Concentration(ion string, mmol float64) {
	cf.Ions[ion] = Ion{Charge: cf.Ions[ion].Charge, Concentration: mmol/(cf.Volume * math.Abs(float64(cf.Ions[ion].Charge)))}
}

func (m *Membrane) Calculate_Diffusion_Coefficient() float64 {
	return (kB * T) / (6 * math.Pi * m.Radius * m.Viscosity)
}

func (m *Membrane) Calculate_Diffusion() float64 {
	D := m.Calculate_Diffusion_Coefficient()
	return (D) / m.Thickness
}

func (m *Membrane) Calculate_Flux(ion string, cA float64, bA float64) float64 {
	// TODO: Calculation for Partition Coefficient

	K := 1.0 // Partition Coefficient
	P := K * m.Calculate_Diffusion()
	return P * m.Area * (cA - bA) // Net rate of diffusion (mmol/s)
}

func (m *Membrane) Calculate_Equilibrium_Potential(ion string) float64 {
	return (-(1.0e+3*(2.303*R*T)/F)/float64(m.ICF.Ions[ion].Charge))*math.Log10((m.ICF.Ions[ion].Concentration * math.Abs(float64(m.ICF.Ions[ion].Charge)))/(m.ECF.Ions[ion].Concentration*math.Abs(float64(m.ECF.Ions[ion].Charge))))
}

func (m *Membrane) Ca2_ATPase() {
	// TODO: Determine actual rate
	// TODO: Add ATP calculation
	
	rate := 500.0 // Assume 400-500 ions per second
	// E1 + E2 State (Ca2+ ICF -> ECF)
	m.Transfer_Ions("Ca", (1.0*rate*1e3)/NA)
}


func (m *Membrane) Na_K_ATPase() {
	// TODO: Determine actual rate
	// TODO: Add ATP calculation

	rate := 500.0 // Assume 400-500 ions per second
	// E1 State (3Na+ ICF -> ECF)
	m.Transfer_Ions("Na", (3.0*rate*1e3)/NA)
	// E2 State (2K+ ECF -> ICF)
	m.Transfer_Ions("K", -(2.0*rate*1e3)/NA)
}

// Move Ions between A & B
func (m *Membrane) Transfer_Ions(ion string, mmol float64) {
	icfMmol := m.ICF.Get_Mmol(ion)
	ecfMmol := m.ECF.Get_Mmol(ion)

	if icfMmol - mmol < 0 {
		mmol = icfMmol
	}
	if ecfMmol + mmol < 0 {
		mmol = ecfMmol
	}

	m.ICF.Set_Concentration(ion, icfMmol - mmol)
	m.ECF.Set_Concentration(ion, ecfMmol + mmol)
}

func (m *Membrane) Simple_Diffusion() {
	// Simple Diffusion
	for _, ion := range m.Ions {
		flux := m.Calculate_Flux(ion, m.ICF.Ions[ion].Concentration, m.ECF.Ions[ion].Concentration)
		m.Transfer_Ions(ion, flux)
	}
}

func (m *Membrane) Membrane_Potential() {
	Meq := 0.0
	for _, ion := range m.Ions {
		Meq = Meq + m.Calculate_Equilibrium_Potential(ion)
	}
	fmt.Printf("Membrane Potential : %v mV \n", Meq)
}


func (m *Membrane) tick(d time.Duration) {
	for range time.Tick(d) {
		m.Simple_Diffusion();
		m.Na_K_ATPase();
		m.Membrane_Potential();
		//fmt.Printf("mEq/L [%v] [ICF]: %v \n", "K", m.ICF.Ions["K"].Concentration)
		//fmt.Printf("mEq/L [%v] [ECF]: %v \n", "K", m.ECF.Ions["K"].Concentration)
	}
}
