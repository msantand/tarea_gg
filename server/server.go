package server

import (
	"github.com/msantand/tarea_gg/database"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/msantand/tarea_gg"
	"io/ioutil"
)

type ServerCities struct {
	databaseCities *database.DbCities
}

func NewServerCities(database *database.DbCities) *ServerCities {
	return &ServerCities{
		databaseCities: database,
	}
}

func (server *ServerCities) PostCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	city := grb.City{}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}
	json.Unmarshal(b, &city)
	server.databaseCities.AddCity(grb.NewCity(city.Name))

	j, _ := json.Marshal(city)
	w.Write(j)
}


func (server *ServerCities) CityList(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(server.databaseCities.CityList(), "", "  ")
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//------------------Connections--------------------------------------

type ServerConnection struct {
	databaseConnection *database.DbConnections
}

func NewServerConnection (database *database.DbConnections) *ServerConnection {
	return &ServerConnection{
		databaseConnection: database,
	}
}

func (server *ServerConnection) PostConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	connection := grb.Connection{}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}
	json.Unmarshal(b, &connection)
	server.databaseConnection.AddConnection(grb.NewConnection(connection.From,connection.To,connection.Cost))

	j, _ := json.Marshal(connection)
	w.Write(j)
}


func (server *ServerConnection) ConnectionList(w http.ResponseWriter, r *http.Request) {
	data, err := json.MarshalIndent(server.databaseConnection.ConnectionList(), "", "  ")
	if err != nil {
		fmt.Errorf("error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
