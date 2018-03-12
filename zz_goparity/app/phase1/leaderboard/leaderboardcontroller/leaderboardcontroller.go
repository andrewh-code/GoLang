package leaderboardcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"zz_goparity/app/phase1/leaderboard/leaderboardmodel"
	"zz_goparity/utilities/prettyprint"
)

func HelloWorldLeaderBoard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world from leader board")
}

func LeaderBoardTest(w http.ResponseWriter, r *http.Request) {

	var player leaderboardmodel.Player
	stats := leaderboardmodel.PlayerStats{Goals: 30, Assists: 30, SecondAssists: 7, Defence: 18, Throwaways: 14, Drops: 0}

	player.TucID = 47736
	player.ID = 1
	player.FirstName = "Bryan"
	player.LastName = "S"
	player.Salary = "$484,500"
	player.Wins = 4.5
	player.TimesTraded = 0
	player.PlayerStats = &stats

	json.NewEncoder(w).Encode(player)

}

func LeaderBoardTestArray(w http.ResponseWriter, r *http.Request) {

	var player []leaderboardmodel.Player
	// var stats []leaderboardmodel.PlayerStats
	var tempPlayer leaderboardmodel.Player
	var tempStats leaderboardmodel.PlayerStats

	tempStats.Goals = 30
	tempStats.Assists = 30
	tempStats.SecondAssists = 7
	tempStats.Defence = 18
	tempStats.Throwaways = 14
	tempStats.Drops = 0

	tempPlayer.TucID = 47736
	tempPlayer.ID = 1
	tempPlayer.FirstName = "Bryan"
	tempPlayer.LastName = "S"
	tempPlayer.Salary = "$484,500"
	tempPlayer.Wins = 4.5
	tempPlayer.TimesTraded = 0
	tempPlayer.PlayerStats = &tempStats

	player = append(player, tempPlayer)

	tempStats.Goals = 26
	tempStats.Assists = 24
	tempStats.SecondAssists = 6
	tempStats.Defence = 11
	tempStats.Throwaways = 17
	tempStats.Drops = 3

	tempPlayer.TucID = 49003
	tempPlayer.ID = 2
	tempPlayer.FirstName = "Shawn"
	tempPlayer.LastName = "H"
	tempPlayer.Salary = "$484,500"
	tempPlayer.Wins = 2
	tempPlayer.TimesTraded = 1
	tempPlayer.PlayerStats = &tempStats

	player = append(player, tempPlayer)

	out := prettyprint.PrettyPrintJSON(player)
	os.Stdout.Write(out)

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	//json.NewEncoder(w).Encode(player)

}
