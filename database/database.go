package database

import (
	"github.com/msantand/tarea_gg"
	"errors"
)

type DbCities struct {
	cities map[string]*grb.City
}

func NewDbCities() *DbCities {
	return &DbCities{
		cities: make(map[string]*grb.City),
	}
}

func (db *DbCities) AddCity(city *grb.City) {
	db.cities[city.Name] = city
}

func (db *DbCities) CityList() []string {
	cities := make([]string, len(db.cities))
	var i = 0
	for _, city := range db.cities {
		cities[i] = city.Name
		i += 1
	}

	return cities
}

func (db *DbCities) CityGet(key string) (*grb.City, error) {
	value, ok := db.cities[key]
	if !ok {
		return nil, errors.New("city not found")
	}
	return value, nil
}

//----------------------------------------------------------------------------------------------------

type DbConnections struct {
	connections map[string]*grb.Connection
}

func NewDbConnection() *DbConnections {
	return &DbConnections{
		connections: make(map[string]*grb.Connection),
	}
}

func (db *DbConnections) AddConnection(connection *grb.Connection) {
	db.connections[connection.From] = connection
}


func (db *DbConnections) ConnectionList() []*grb.Connection {
	cities := make([]*grb.Connection, len(db.connections))

	var i = 0
	for _, entry := range db.connections {
		cities[i] = entry
		i += 1
	}

	return cities
}