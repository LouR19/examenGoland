package main

import (
	"fmt"
)

type song struct {
	titulo    string
	artista   string
	siguiente *song
}
type linkedlist struct {
	t *song
}

func main() {
	opc := 0
	list := linkedlist{}
	fmt.Println("Bienvenido!")
	for {
		fmt.Println("---------------------------------------------")
		fmt.Println("------------Lista de Reproduccion------------\n 1-Agregar musica \n 2-Eliminar \n 3-Buscar \n 4-Imprimir \n 5-Salir")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			agregarmusica(&list)
		case 2:
			eliminar(&list)
		case 3:
			buscar(&list)
		case 4:
			imprimir(&list)
		case 5:
			fmt.Println("¡Hasta luego, Byebye!")
			return
		default:
			fmt.Println("---------------------------------------------")
			fmt.Println("Opcion invalida, intente nuevamente. Gracias!")
			continue
		}

	}
}
func agregarmusica(linkedlist *linkedlist) {
	var titulo, artista string
	fmt.Println("---------------------------------------------")
	fmt.Println("Ingrese el titulo de la musica: ")
	fmt.Scan(&titulo)
	fmt.Println("Ingrese el artista de la musica: ")
	fmt.Scan(&artista)

	aggnodo := &song{
		titulo:  titulo,
		artista: artista,
	}

	if linkedlist.t == nil {
		linkedlist.t = aggnodo
	} else {
		ultimonodo := linkedlist.t
		for ultimonodo.siguiente != nil {
			ultimonodo = ultimonodo.siguiente
		}
		ultimonodo.siguiente = aggnodo
	}
	fmt.Println("Musica agregada exitosamente!")
}

func eliminar(linkedlist *linkedlist) {
	var titulo string
	fmt.Println("---------------------------------------------")
	fmt.Println("Ingrese el titulo de la musica a eliminar: ")
	fmt.Scan(&titulo)

	if linkedlist.t == nil {
		fmt.Println("La lista está vacía.")
		return
	}

	if linkedlist.t.titulo == titulo {
		linkedlist.t = linkedlist.t.siguiente
		return
	}

	nodo := linkedlist.t
	for nodo.siguiente != nil && nodo.siguiente.titulo != titulo {
		nodo = nodo.siguiente
	}

	if nodo.siguiente == nil {
		fmt.Printf("La musica: %s, no se encontro en tu lista de reproduccion...\n", titulo)
		return
	}
	fmt.Println("Música eliminada correctamente!")
	nodo.siguiente = nodo.siguiente.siguiente
}

func imprimir(linkedlist *linkedlist) {
	if linkedlist.t == nil {
		fmt.Println("  vacia  ")
		return
	}
	fmt.Println("---------------------------------------------")
	fmt.Println("-------------Lista de Reproducción-----------")
	nodo := linkedlist.t
	for nodo != nil {
		fmt.Printf("Titulo: %s, Autor: %s\n", nodo.titulo, nodo.artista)
		nodo = nodo.siguiente
	}

}

func buscar(linkedlist *linkedlist) {
	var titulo string

	if linkedlist.t == nil {
		fmt.Println("Lista vacia")
		return
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("Ingrese titulo de la musica que desee buscar: ")
	fmt.Scan(&titulo)
	nodo := linkedlist.t
	for nodo != nil {
		if nodo.titulo == titulo {
			fmt.Printf("Título: %s, Autor: %s\n", nodo.titulo, nodo.artista)
			return
		}
		nodo = nodo.siguiente
	}

	fmt.Println("Música no encontrada en la lista...")

}
