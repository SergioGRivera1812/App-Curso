package main

import (
	"log"

	"github.com/SergioGRivera1812/App-Curso/bd"
	"github.com/SergioGRivera1812/App-Curso/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la Base ")
		return
	}
	handlers.Manejadores()
}
