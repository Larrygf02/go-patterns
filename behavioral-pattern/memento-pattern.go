//Memento es un patro de diseño que le permite guardar
// y restaurar el estado anterior de un objeto
package main

import "fmt"

//Interface
type Memento interface {
	SetContenido(string)
	GetContenido() string
}

//Memento
type EditorMemento struct {
	contenido string
}

func (em *EditorMemento) SetContenido(contenido string) {
	em.contenido = contenido
}

func (em *EditorMemento) GetContenido() string {
	return em.contenido
}

// Originador
type Editor struct {
	contenido string
}

func (e *Editor) VerContenido() string {
	return e.contenido
}

func (e *Editor) Escribir(texto string) {
	e.contenido = e.contenido + " " + texto
}

func (e *Editor) Guardar() Memento {
	editorMemento := &EditorMemento{}
	editorMemento.SetContenido(e.contenido)
	return editorMemento
}

func (e *Editor) Restaurar(memento Memento) {
	e.contenido = memento.GetContenido()
}

func main() {
	editor := &Editor{}
	editor.Escribir("Texto A")
	editor.Escribir("Texto B")
	fmt.Println(editor.VerContenido())

	memento := editor.Guardar()

	editor.Escribir("TextoC")
	fmt.Println(editor.VerContenido())

	memento = editor.Guardar()
	editor.Escribir("TextoD")
	fmt.Println(editor.VerContenido())

	editor.Restaurar(memento)
	fmt.Println(editor.VerContenido())
}
