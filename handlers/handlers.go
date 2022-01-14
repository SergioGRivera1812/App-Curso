package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/SergioGRivera1812/App-Curso/middlew"
	"github.com/SergioGRivera1812/App-Curso/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores seteo mi puerto , handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter() // creando router

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) //Se pone a escuchar el puerto

}
