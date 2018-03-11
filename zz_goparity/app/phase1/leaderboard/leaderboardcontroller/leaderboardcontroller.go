package leaderboardcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zz_goparity/app/phase1/leaderboard/leaderboardmodel"
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
