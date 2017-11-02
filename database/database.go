package database

import (
	"github.com/msantand/tarea_gg"
	_ "errors"
	"fmt"
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

//----------------------------------------------------------------------------------------------------

type DbConnections struct {
	connections map[grb.Destination]*grb.Connection
}

func NewDbConnection() *DbConnections {
	return &DbConnections{
		connections: make(map[grb.Destination]*grb.Connection),
	}
}

func (db *DbConnections) AddConnection(connection *grb.Connection) {
	destination := grb.Destination{
		Dest1: connection.From,
		Dest2: connection.To,
	}
	fmt.Println(destination)
	flag, key := db.SameConnection(destination)
	fmt.Println(key)
	if  flag{
		fmt.Println("entre aki")
		db.connections[key].Cost = connection.Cost
		return
	}
	db.connections[destination] = connection
}

func (db *DbConnections) SameConnection(destination grb.Destination) (bool, grb.Destination){
	for key := range db.connections{
		if CompareDestination(destination, key){
			return true, key
		}
	}
	return false, grb.Destination{}
}

func CompareDestination(dest1 grb.Destination, dest2 grb.Destination) bool{
	if dest1 == dest2{
		return true
	}
	if (dest1.Dest1 == dest2.Dest2) && (dest1.Dest2 == dest2.Dest1){
		return true
	}
	return false
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