package individualcontroller

import (
	"net/http"
	"os"
	"time"
	"zz_goparity/app/phase1/stats/playermodel"
	"zz_goparity/utilities/prettyprint"
)

func IndividualTest(w http.ResponseWriter, r *http.Request) {

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

	out := prettyprint.PrettyPrintJSON(player)
	os.Stdout.Write(out)

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	//json.NewEncoder(w).Encode(player)

}
