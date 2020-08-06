package main

import "fmt"

// Interface
type Receptor interface {
	ProcesarMensaje(int, string) string
}

// Receptor de alta prioridad
type ReceptorAltaPrioridad struct {
	siguiente Receptor
}

func (rap ReceptorAltaPrioridad) ProcesarMensaje(prioridad int, mensaje string) string {
	if prioridad >= 5 {
		return "Procesando mensaje de alta prioridad " + mensaje
	}
	if rap.siguiente != nil {
		return rap.siguiente.ProcesarMensaje(prioridad, mensaje)
	}
	return ""
}

// Receptor de baja prioridad
type ReceptorBajaPrioridad struct {
	siguiente Receptor
}

func (rbp ReceptorBajaPrioridad) ProcesarMensaje(prioridad int, mensaje string) string {
	if prioridad < 5 {
		return "Procesando mensaje de baja prioridad: " + mensaje
	}
	if rbp.siguiente != nil {
		return rbp.siguiente.ProcesarMensaje(prioridad, mensaje)
	}
	return ""
}

func mainx() {
	manejadores := ReceptorBajaPrioridad{
		siguiente: ReceptorAltaPrioridad{},
	}
	fmt.Println(manejadores.ProcesarMensaje(4, "Mensaje publico "))
	fmt.Println(manejadores.ProcesarMensaje(3, "Otro mensaje publico "))
	fmt.Println(manejadores.ProcesarMensaje(8, "Envio de las llaves privadas "))
}
