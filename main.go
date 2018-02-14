package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var leagueAPIKey = os.Args[1]

type Summoner struct {
	ID            int    `json:"id"`
	AccountID     int    `json:"accountId"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}

func GetSummoner(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	username := params["username"]
	URL := "https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/" + username + "?api_key=" + leagueAPIKey
	jsonData := ""

	//request data from riotgames api
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		jsonData = string(data)
	}
	//print out json received from riotgames
	fmt.Fprint (w, jsonData) 

	var userInfo Summoner //create Summoner object to hold passed in info

	//Unmarshal into a Summoner object
	err = json.Unmarshal([]byte(jsonData), &userInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "\nusername: " + userInfo.Name)

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/summonername/{username}", GetSummoner).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
