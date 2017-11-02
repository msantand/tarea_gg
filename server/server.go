package server

import (
	"github.com/msantand/tarea_gg/database"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/msantand/tarea_gg"
	"io/ioutil"
	_ "encoding/asn1"
)


type Server struct {
	ServerCities
	ServerConnections
	Matrix [10][10]int
}

type ServerCities struct {
	databaseCities 		*database.DbCities
}

type ServerConnections struct {
	databaseConnections *database.DbConnections
}

//----------------------------------

func NewServer(dbCities *database.DbCities, dbConnections *database.DbConnections) *Server {
	return &Server{
			ServerCities: ServerCities{
				databaseCities: dbCities,
			},
			ServerConnections: ServerConnections{
				databaseConnections: dbConnections,
			},
	}
}

func (server *Server) PostCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	city := grb.City{}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}
	json.Unmarshal(b, &city)

	if city.Name == ""{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	server.ServerCities.databaseCities.AddCity(grb.NewCity(city.Name))

	j, _ := json.Marshal(city)
	w.Write(j)
}


func (server *Server) CityList(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(server.ServerCities.databaseCities.CityList(), "", "  ")
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//------------------Connections--------------------------------------


func (server *Server) PostConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	connection := grb.Connection{}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}
	json.Unmarshal(b, &connection)


	cities := server.ServerCities.databaseCities.CityList()

	flagTo := false
	flagFrom := false

	for _, v := range cities {
		if v == connection.From  {
			flagFrom = true
		}
		if v == connection.To{
			flagTo = true
		}
	}
	if !(flagFrom && flagTo) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	server.ServerConnections.databaseConnections.AddConnection(grb.NewConnection(connection.From,connection.To,connection.Cost))

	j, _ := json.Marshal(connection)
	w.Write(j)
}


func (server *Server) ConnectionList(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(server.ServerConnections.databaseConnections.ConnectionList(), "", "  ")
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}