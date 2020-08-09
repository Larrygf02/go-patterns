// objetivo
package main

import "fmt"

type Guerrero interface {
	UsarArma() string
}

type Elfo struct{}

func (e *Elfo) UsarArma() string {
	return "Atacando con arco y flecha"
}

//Adaptable
type GuerreroMagico interface {
	UsarMagia() string
}

type Mago struct{}

func (m *Mago) UsarMagia() string {
	return "Atacando con Magia"
}

//Adaptador
type MagoAdaptador struct {
	guerrero GuerreroMagico
}

func (ma *MagoAdaptador) UsarArma() string {
	return ma.guerrero.UsarMagia()
}

// cliente
type Jugador struct {
	guerrero Guerrero
}

func (j *Jugador) Atacar() string {
	return j.guerrero.UsarArma()
}

func _() {
	jugadorA := &Jugador{&Elfo{}}
	fmt.Printf("Jugador A: %s\n", jugadorA.Atacar())

	jugadorB := &Jugador{&MagoAdaptador{&Mago{}}}
	fmt.Printf("Jugador B: %s\n", jugadorB.Atacar())
}
