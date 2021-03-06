package main

import "fmt"

type NavegadorInterface interface {
	Direccion(string) string
}

type Navegador struct{}

func (n *Navegador) Direccion(url string) string {
	return "Respuesta de la url " + url
}

// proxy
type NavegadorProxy struct {
	navegador NavegadorInterface
}

func (n *NavegadorProxy) Direccion(url string) string {
	if url == "http://twitter.com" || url == "http://facebook.com" {
		return "Acceso restringido a " + url
	}
	return n.navegador.Direccion(url)
}

func _() {
	navegadorProxy := &NavegadorProxy{&Navegador{}}
	fmt.Printf("%s\n", navegadorProxy.Direccion("http://google.com"))
	fmt.Printf("%s\n", navegadorProxy.Direccion("http://twitter.com"))
	fmt.Printf("%s\n", navegadorProxy.Direccion("http://facebook.com"))
}
