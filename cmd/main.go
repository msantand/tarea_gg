package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/msantand/tarea_gg/database"
	"github.com/msantand/tarea_gg/server"
)

func main() {
	dbCities := database.NewDbCities()
	dbConnection := database.NewDbConnection()


	myServer := server.NewServer(dbCities, dbConnection)

	router := mux.NewRouter()
	router.NewRoute().
		Path("/cities").
		HandlerFunc(myServer.CityList).
		Methods("GET").
		Name("CitiesList")

	router.NewRoute().
		Path("/cities").
		HandlerFunc(myServer.PostCity).
		Methods("POST").
		Name("AddingCity")

	router.NewRoute().
		Path("/connections").
		HandlerFunc(myServer.PostConnection).
		Methods("POST").
		Name("AddingConnection")

	router.NewRoute().
		Path("/connections").
		HandlerFunc(myServer.ConnectionList).
		Methods("GET").
		Name("ConnectionList")
/*
	router.NewRoute().
		Path("/solve/?from={name}&to={name}").
		HandlerFunc(myServerCities.CityList).
		Methods("GET").
		Name("CitiesList")
*/
	http.ListenAndServe(":8000", router)
}
