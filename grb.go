package grb

type City struct {
	Name string `json:"name"`
}

type Connection struct {
	From    string    `json:"from"`
	To      string    `json:"to"`
	Cost    int    	  `json:"cost"`
}


type Solve struct {
    Cost    string
    Path    []string

}

func NewCity(newCity string) *City {
	return &City{
		Name: newCity,
	}
}

func NewConnection(c1 string, c2 string, cost int) *Connection {
	return &Connection{
		From: c1,
		To:   c2,
		Cost: cost,
	}
}
