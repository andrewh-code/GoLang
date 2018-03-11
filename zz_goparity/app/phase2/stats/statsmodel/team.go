package model

type Team struct {
	ID        int16  `json:playerid`
	FirstName string `json:gmfirstname`
	LastName  string `json:gmlastname`
	TeamName  string `json:teamname`
}
