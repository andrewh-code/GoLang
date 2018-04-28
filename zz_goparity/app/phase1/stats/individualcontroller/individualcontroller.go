package individualcontroller

import (
	"net/http"
	"os"
	"time"
	"zz_goparity/app/phase1/stats/playermodel"
	"zz_goparity/utilities/prettyprint"

	"log"

	"strconv"

	"github.com/gorilla/mux"
)

func IndividualStats(w http.ResponseWriter, r *http.Request) {

	// url parameters are string hashmap String[String]
	// convert variable to int64
	playerID, err := strconv.ParseInt(mux.Vars(r)["playerid"], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	//convert playerId to int64

	player := mockPlayer()
	var out []byte

	if playerID == player.TucID {
		out = prettyprint.PrettyPrintJSON(player)
		os.Stdout.Write(out)

	} else {
		out = []byte("Error unable to find player...")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

}

func IndividualTest(w http.ResponseWriter, r *http.Request) {

	player := mockPlayer()

	out := prettyprint.PrettyPrintJSON(player)
	os.Stdout.Write(out)

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	//json.NewEncoder(w).Encode(player)

}

func mockPlayer() playermodel.Player {

	var player playermodel.Player
	stats := playermodel.PlayerStats{Goals: 30, Assists: 30, SecondAssists: 7, Defence: 18, Throwaways: 14, Drops: 0}

	player.LastUpdated = time.Now().Unix()
	player.TucID = 47736
	player.ID = 1
	player.FirstName = "Bryan"
	player.LastName = "S"
	player.Salary = "$484,500"
	player.Wins = 4.5
	player.TimesTraded = 0
	player.PlayerStats = &stats

	return player
}
