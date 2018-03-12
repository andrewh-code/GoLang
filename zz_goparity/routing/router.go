package routing

// base middleware design off of: https://github.com/mingrammer/go-todo-rest-api-example

import (
	// custom libraries
	"fmt"
	"net/http"
	"zz_goparity/app/phase1/leaderboard/leaderboardcontroller"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) InitializeRouter() {

	// create new gorilla mux router
	r.Router = mux.NewRouter().StrictSlash(true)

}

func (r *Router) SetRoutes() {

	//index
	r.Router.HandleFunc("/", HelloWorld).Methods("GET")

	// TODO: routing
	//r.Router.HandleFunc("/stats", statscontroller.HelloWorldStats).Methods("GET")
	r.Router.HandleFunc("/stats", leaderboardcontroller.LeaderBoardTestArray).Methods("GET")
	r.Router.HandleFunc("/stats/{playerid}", DummyFunc).Methods("GET")
	r.Router.HandleFunc("/stats/leaderboard", DummyFunc).Methods("GET")
	r.Router.HandleFunc("/stats/leaderboard/goals", DummyFunc).Methods("GET")
	r.Router.HandleFunc("/stats/leaderboard/assists", DummyFunc).Methods("GET")
	r.Router.HandleFunc("/stats/leaderboard/secondassists", DummyFunc).Methods("GET")
	r.Router.HandleFunc("/stats/leaderboard/defences", DummyFunc).Methods("GET")
	r.Router.HandleFunc("/stats/leaderboard/throaways", DummyFunc).Methods("GET")
	r.Router.HandleFunc("/stats/leaderboard/drops", DummyFunc).Methods("GET")
}

/*
	create wrapper functions (to make things look cleaner when setting routes)
	for the HTTP methods (get, put, patch, post, etc)
*/

// Get wrapper method for HTTP GET method
// func (r *Router) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	r.Router.HandleFunc(path, f).Methods("GET")
// }

// // Post wrapper method for HTTP POST method
// func (r *Router) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	r.Router.HandleFunc(path, f).Methods("POST")
// }

// // Put wrapper method for HTTP PUT method
// func (r *Router) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	r.Router.HandleFunc(path, f).Methods("PUT")
// }

// // Delete wrapper method for HTTP DELETE method
// func (r *Router) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	r.Router.HandleFunc(path, f).Methods("DELETE")
// }

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func DummyFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dummy api endpoint, replace")
}

func (r *Router) Run(port string) {
	err := http.ListenAndServe(port, r.Router)
	if err != nil {
		panic(err)
	}
}
