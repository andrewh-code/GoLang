package leaderboardmodel

// LeaderBoardPlayer - model for the tpl spreadsheet
type Player struct {
	ID          int64        `json:playerid`
	TucID       int64        `json:tucid`
	LastUpdated int64        `json:lastupdated`
	FirstName   string       `json:firstname`
	LastName    string       `json:lastname`
	Salary      string       `json:salary`
	Wins        float32      `json:wins`
	TimesTraded int          `json:timestraded`
	PlayerStats *PlayerStats `json:stats`
}

// LeaderBoardPlayerStats - subnode structure for leaderboard stats
type PlayerStats struct {
	Goals         int `json:goals`
	Assists       int `json:assists`
	SecondAssists int `json:secondassists`
	Defence       int `json:defence`
	Throwaways    int `json:throwaways`
	Drops         int `json:drops`
}
