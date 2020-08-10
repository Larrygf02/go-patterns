package main

import "fmt"

// Abstraccion interface
type DispositivoInterface interface {
	ConectarInternet() string
	SetConexion(Conexion)
}

// Abstraccion Abstracta
type Dispositivo struct {
	conexion Conexion
}

func (d *Dispositivo) SetConexion(conexion Conexion) {
	d.conexion = conexion
}

// Abstraccion refinada
type Telefono struct {
	numero string
	*Dispositivo
}

func (t *Telefono) ConectarInternet() string {
	return "Teléfono N° " + t.numero + " conectado a internet mediante " + t.conexion.Conectar()
}

// Abstraccion refinada
type Tablet struct {
	*Dispositivo
}

func (t *Tablet) ConectarInternet() string {
	return "Tablet conectada a internet mediante " + t.conexion.Conectar()
}

// Implementador interface
type Conexion interface {
	Conectar() string
}

type Red4G struct{}

func (r *Red4G) Conectar() string {
	return "4G"
}

type RedWifi struct{}

func (r *RedWifi) Conectar() string {
	return "Wifi"
}

func _() {
	telefonoA := &Telefono{"01230", &Dispositivo{}}
	telefonoA.SetConexion(&Red4G{})
	fmt.Printf("%s\n", telefonoA.ConectarInternet())

	telefonoB := &Telefono{"93929", &Dispositivo{}}
	telefonoB.SetConexion(&RedWifi{})
	fmt.Printf("%s\n", telefonoB.ConectarInternet())

	tablet := &Tablet{&Dispositivo{}}
	tablet.SetConexion(&RedWifi{})
	fmt.Printf("%s\n", tablet.ConectarInternet())
}
