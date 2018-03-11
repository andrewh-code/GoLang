package model

type Player struct {
	ID          int16        `json:playerid`
	TucID       int16        `json:tucid`
	FirstName   string       `json:firstname`
	LastName    string       `json:lastname`
	CurrentTeam string       `json:currentteam`
	Salary      string       `json:salary`
	PlayerStats *PlayerStats `json:stats`
	Wins        int          `json:wins`
	TimesTraded int          `json:timestraded`
}

type PlayerStats struct {
	Goals         int `json:goals`
	Assists       int `json:assists`
	SecondAssists int `json:secondassists`
	Defence       int `json:defence`
	Throwaways    int `json:throwaways`
	Drops         int `json:drops`
}
