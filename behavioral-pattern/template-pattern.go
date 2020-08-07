package main

import "fmt"

type DeployInterface interface {
	Testear()
	Compilar()
	Publicar()
}

type Deploy struct{}

func (d Deploy) Construir(di DeployInterface) {
	fmt.Println("Ejecutando las siguientes acciones :")
	di.Testear()
	di.Compilar()
	di.Publicar()
}

//Clase concreta - Android
type DeployAndroid struct {
	Deploy
}

func (d DeployAndroid) Testear() {
	fmt.Println("Android: Testeando")
}

func (d DeployAndroid) Compilar() {
	fmt.Println("Android: compilando")
}

func (d DeployAndroid) Publicar() {
	fmt.Println("Android: Publicando")
}

//Clase Concreata -ios
type DeployiOs struct {
	Deploy
}

func (d DeployiOs) Testear() {
	fmt.Println("iOs: Testeando")
}

func (d DeployiOs) Compilar() {
	fmt.Println("iOs: compilando")
}

func (d DeployiOs) Publicar() {
	fmt.Println("iOs: Publicando")
}

func main() {
	deployAndroid := DeployAndroid{Deploy{}}
	deployAndroid.Construir(&deployAndroid)

	deployiOs := DeployiOs{Deploy{}}
	deployiOs.Construir(&deployiOs)
}
