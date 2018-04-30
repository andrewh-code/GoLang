package mockdata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// PLayer - model for the tpl spreadsheet
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

// PlayerStats - subnode structure for leaderboard stats
type PlayerStats struct {
	Goals         int `json:goals`
	Assists       int `json:assists`
	SecondAssists int `json:secondassists`
	Defence       int `json:defence`
	Throwaways    int `json:throwaways`
	Drops         int `json:drops`
}

func (p Player) toString() string {
	return toJSON(p)
}

func toJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func main() {

	players := getPlayers()
	// for _, p := range pages {
	// 	fmt.Println(p.toString())
	// }

	fmt.Println(toJSON(players))
}

func getPlayers() []Player {
	raw, err := ioutil.ReadFile("./test_json.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var p []Player
	json.Unmarshal(raw, &p)
	return p
}
