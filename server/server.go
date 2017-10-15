package server

import (
	"github.com/msantand/tarea_gg/database"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
)

type Server struct {
	database *database.Database
}

func New(database *database.Database) *Server {
	return &Server{
		database: database,
	}
}

func (server *Server) EntryList(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(server.database.EntryList(), "", "  ")
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (server *Server) EntryGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	entry, err := server.database.EntryGet(key)
	if err != nil {
		log.Printf("error: %s\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.MarshalIndent(entry, "", "  ")
	if err != nil {
		log.Printf("error: %s\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
