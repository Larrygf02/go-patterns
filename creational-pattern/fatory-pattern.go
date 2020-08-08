package main
import "fmt"
// Producto abstracto
type Puerta interface {
	VerMaterial() string
}

// producto concreto
type PuertaMadera struct{}

func (pm *PuertaMadera) VerMaterial() string {
	return "Madera"
}

// producto concreto
type PuertaMetal struct{}

func (pm *PuertaMetal) VerMaterial() string {
	return "Metal"
}

// producto concreto
type PuertaAdobe struct{}

func (pa *PuertaAdobe) VerMaterial() string {
	return "Adobe"
}

// Fabrica Abstracta Interface
type FabricaPuerta interface {
	ConstruirPuerta() Puerta
}

// Fabrica Concreta
type FabricaPuertaMadera struct{}

func (fpm *FabricaPuertaMadera) ConstruirPuerta() Puerta {
	return &PuertaMadera{}
}

// Fabrica Concreta
type FabricaPuertaMetal struct{}

func (fpm *FabricaPuertaMetal) ConstruirPuerta() Puerta {
	return &PuertaMetal{}
}

// Fabrica Concreta
type FabricaPuertaAdobe struct{}

func (fpa *FabricaPuertaAdobe) ConstruirPuerta() Puerta {
	return &PuertaAdobe{}
}

func main() {
	fabricaPuertaMadera := &FabricaPuertaMadera{}
	puertaMadera := fabricaPuertaMadera.ConstruirPuerta()
	fmt.Printf("Se construyo una puerta de: %s\n", puertaMadera.VerMaterial())

	fabricaPuertaMetal := &FabricaPuertaMetal{}
	puertaMetal := fabricaPuertaMetal.ConstruirPuerta()
	fmt.Printf("Se construyo una puerta de: %s\n", puertaMetal.VerMaterial())

	fabricaPuertaAdobe := &FabricaPuertaAdobe{}
	puertaAdobe := fabricaPuertaAdobe.ConstruirPuerta()
	fmt.Printf("Se construyo una puerta de: %s\n", puertaAdobe.VerMaterial())
}