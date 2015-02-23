package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/ballot/{id:[0-9]+}", getBallot).Methods("GET")
	rtr.HandleFunc("/ballot/{id:[0-9]+}", postBallot).Methods("POST")
	rtr.HandleFunc("/make", getMake).Methods("GET")
	rtr.HandleFunc("/make", postMake).Methods("POST")
	rtr.PathPrefix("/").Handler(http.FileServer(http.Dir("./website")))

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

func getBallot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Println("GET request on ballot id: " + id)
	w.Write([]byte("Hello TWAT " + id))
}

func postBallot(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("Fuck you " + name))
}

func getMake(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(http.Dir("index.html")))
}

func postMake(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(http.Dir("index.html")))
}
