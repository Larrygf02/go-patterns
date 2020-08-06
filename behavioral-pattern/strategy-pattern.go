// Interface
package main

import "fmt"

type Estrategia interface {
	RealizarOperacion(int, int) int
}

// Estrategia Suma
type EstrategiaSuma struct{}

func (e EstrategiaSuma) RealizarOperacion(num1 int, num2 int) int {
	return num1 + num2
}

// Estrategia Resta
type EstrategiaResta struct{}

func (e EstrategiaResta) RealizarOperacion(num1 int, num2 int) int {
	return num1 - num2
}

// Estrategia Multiplicacion
type EstrategiaMultiplica struct{}

func (e EstrategiaMultiplica) RealizarOperacion(num1 int, num2 int) int {
	return num1 * num2
}

// Contexto
type Contexto struct {
	estrategia Estrategia
}

func (c *Contexto) EjecutarOperacion(num1 int, num2 int) int {
	return c.estrategia.RealizarOperacion(num1, num2)
}

var contexto Contexto

func _() {
	num1 := 10
	num2 := 5
	// suma
	contexto = Contexto{EstrategiaSuma{}}
	suma := contexto.EjecutarOperacion(num1, num2)
	fmt.Println(suma)
	// resta
	contexto = Contexto{EstrategiaResta{}}
	resta := contexto.EjecutarOperacion(num1, num2)
	fmt.Println(resta)
	// multiplicacion
	contexto = Contexto{EstrategiaMultiplica{}}
	producto := contexto.EjecutarOperacion(num1, num2)
	fmt.Println(producto)
}
