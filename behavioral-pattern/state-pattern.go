// Interface estado
package main

import "fmt"

type Estado interface {
	Escribir(string) string
}

// Estado concreto
type EstadoNegrita struct{}

func (en *EstadoNegrita) Escribir(texto string) string {
	return "*" + texto + "*"
}

//Estado concreto
type EstadoCursiva struct{}

func (ec *EstadoCursiva) Escribir(texto string) string {
	return "_" + texto + "_"
}

// Contexto
type EditorMarkdown struct {
	estado Estado
}

func (em *EditorMarkdown) SetEstado(estado Estado) {
	em.estado = estado
}

func (em *EditorMarkdown) Redactar(texto string) string {
	if em.estado == nil {
		return texto
	}
	return em.estado.Escribir(texto)
}

func _() {
	editor := &EditorMarkdown{}
	fmt.Printf("Texto redactado sin estado: %s\n", editor.Redactar("Lorem ipsum"))

	editor.SetEstado(&EstadoNegrita{})
	fmt.Printf("Texto redactado en negrita: %s\n", editor.Redactar("Lorem ipsum"))

	editor.SetEstado(&EstadoCursiva{})
	fmt.Printf("Texto redactado en cursiva: %s\n", editor.Redactar("Lorem ipsum"))
}
