package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/msantand/tarea_gg"
	"github.com/msantand/tarea_gg/database"
	"github.com/msantand/tarea_gg/server"
)

func main() {
	db := database.New()
	db.AddEntry(grb.NewFakeEntry())
	db.AddEntry(grb.NewFakeEntry())
	db.AddEntry(grb.NewFakeEntry())

	myServer := server.New(db)

	router := mux.NewRouter()
	router.NewRoute().
		Path("/").
		HandlerFunc(myServer.EntryList).
		Methods("GET").
		Name("EntryList")

	router.NewRoute().
		Path("/{key}").
		HandlerFunc(myServer.EntryGet).
		Methods("GET").
		Name("EntryGet")

	http.ListenAndServe(":8000", router)
}
