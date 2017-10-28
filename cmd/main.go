package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/msantand/tarea_gg/database"
	"github.com/msantand/tarea_gg/server"
)

func main() {
	dbCities := database.NewDbCities()
	myServerCities := server.NewServerCities(dbCities)

	router := mux.NewRouter()
	router.NewRoute().
		Path("/cities").
		HandlerFunc(myServerCities.CityList).
		Methods("GET").
		Name("CitiesList")

	router.NewRoute().
		Path("/cities").
		HandlerFunc(myServerCities.PostCity).
		Methods("POST").
		Name("AddingCity")



	dbConnection := database.NewDbConnection()
	myServerConnection := server.NewServerConnection(dbConnection)

	router.NewRoute().
		Path("/connections").
		HandlerFunc(myServerConnection.PostConnection).
		Methods("POST").
		Name("AddingConnection")

	router.NewRoute().
		Path("/connections").
		HandlerFunc(myServerConnection.ConnectionList).
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
