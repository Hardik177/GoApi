package main

import (
    "fmt"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Player struct {
	ID     string  `json:"id"`
	Team string `json: "Team"`
	JerseyNo   int  `json:"JerseyNo"`
	Name  *Name  `json:"Name"`
	Position string `json:"Position"`
}

type Name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var players []Player

func getPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
    fmt.Print(params)
	for _, item := range players {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Player{})
}

func getPlayerByTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
    fmt.Print(params)
	for _, item := range players {
		if item.Team == params["Team"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func getPlayerByTeamPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
    fmt.Print(params)
	for _, item := range players {
		if item.Team == params["Team"] && item.Position == params["Position"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var player Player
	_ = json.NewDecoder(r.Body).Decode(&player)
	player.ID = strconv.Itoa(rand.Intn(100000000))
	players = append(players, player)
	json.NewEncoder(w).Encode(player)
}

func updatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range players {
		if item.ID == params["id"] {
			players = append(players[:index], players[index+1:]...)
			var player Player
			_ = json.NewDecoder(r.Body).Decode(&player)
			player.ID = params["id"]
			players = append(players, player)
			json.NewEncoder(w).Encode(player)
			return
		}
	}
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range players {
		if item.ID == params["id"] {
			players = append(players[:index], players[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(players)
}

func main() {
	r := mux.NewRouter()

	players = append(players, Player{ID: "1", Team:"Real Madrid", JerseyNo: 9, Name: &Name{Firstname: "Karim", Lastname: "Benzema"}, Position: "CF"})
	players = append(players, Player{ID: "2", JerseyNo: 10, Team:"Barcelona", Name: &Name{Firstname: "Lionel", Lastname: "Messi"}, Position: "RWF"})
	players = append(players, Player{ID: "3", JerseyNo: 1, Team:"Barcelona", Name: &Name{Firstname: "Marc-andre", Lastname: "TerStegen"}, Position: "GK"})
	players = append(players, Player{ID: "4", JerseyNo: 3, Team:"Barcelona", Name: &Name{Firstname: "Gerrad", Lastname: "Pique"}, Position: "CB"})
	players = append(players, Player{ID: "5", JerseyNo: 21, Team:"Barcelona", Name: &Name{Firstname: "Frankie", Lastname: "De Jong"}, Position: "CMF"})
	players = append(players, Player{ID: "6", JerseyNo: 18, Team:"Barcelona", Name: &Name{Firstname: "Jordi", Lastname: "Alaba"}, Position: "LB"})
	players = append(players, Player{ID: "7", Team:"Real Madrid", JerseyNo: 4, Name: &Name{Firstname: "Sergio", Lastname: "Ramos"}, Position: "CB"})
	players = append(players, Player{ID: "8", Team:"Real Madrid", JerseyNo: 5, Name: &Name{Firstname: "Rafael", Lastname: "Varane"}, Position: "CB"})
	players = append(players, Player{ID: "9", Team:"Real Madrid", JerseyNo: 8, Name: &Name{Firstname: "Toni", Lastname: "Kroos"}, Position: "CMF"})
	players = append(players, Player{ID: "10", Team:"Real Madrid", JerseyNo: 10, Name: &Name{Firstname: "Luka", Lastname: "Modric"}, Position: "CMF"})


	r.HandleFunc("/", getPlayers).Methods("GET")
	r.HandleFunc("/{id}", getPlayer).Methods("GET")
	r.HandleFunc("/filter/{Team}/{Position}", getPlayerByTeamPosition).Methods("GET")
	r.HandleFunc("/filter/{Team}", getPlayerByTeam).Methods("GET")
	r.HandleFunc("/", createPlayer).Methods("POST")
	r.HandleFunc("/{id}", updatePlayer).Methods("PUT")
	r.HandleFunc("/{id}", deletePlayer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
