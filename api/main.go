package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

//I decided to use global variables instead of a DB just for saving time in the dev phase

var byteValue []byte
var result root
var toRet toReturn

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("task.recording.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened json file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &result)

	for i, v := range result.Records {
		if v.Setup.NodeName != "" {
			toRet.Records = append(toRet.Records, item{i, v.Event.T, v.Time, v.Setup.NodeName})
		} else {
			toRet.Records = append(toRet.Records, item{i, v.Event.T, v.Time, v.Setup.URL})
		}
	}

	/*
		*********ALTERNATIVE WITH DB INTERACTION**********

		I would have split the input file into different n chunks and assigned them to a pool of x goroutines worker (worker pool).
		Each goroutine would have taken a chunk from the remaining ones and inserted it into the DB.

		chunks =   | c0 | c1 | c2 | ..... | cn |
		inserted = | 0  | 0  | 1  | ..... | 0  |


	*/

	r := mux.NewRouter()
	r.HandleFunc("/api/test", Test).Methods("GET")
	r.HandleFunc("/api/data", GetFullData).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/deleteItemWithID", DeleteItemWithID).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/api/stats", EvaluateStats).Methods("GET", "OPTIONS")

	server := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:3000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("Server is listening on " + server.Addr)
	server.ListenAndServe()
}

//Test function description here
func Test(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//GetFullData function description here
func GetFullData(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

	w.Header().Set("content-type", "application/json")
	toReturn, _ := json.Marshal(toRet)
	w.Write(toReturn)
}

//DeleteItemWithID function description here
func DeleteItemWithID(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}
	//get id parameter
	if req.Body != nil {
		defer req.Body.Close()
	}

	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var id idStruct
	jsonErr := json.Unmarshal(body, &id)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	toRet.Records = removeIt(id.ID, toRet.Records)
	/*
		*********ALTERNATIVE WITH DB INTERACTION**********

		DELETE FROM public.records
		WHERE id = 'given_id';


	*/

	w.Header().Set("content-type", "application/json")
	toReturnJSON, _ := json.Marshal(toRet)
	w.Write(toReturnJSON)
}

//EvaluateStats function description here
func EvaluateStats(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}
	log.Println("------------- STARTING EVALUATIONS -------------")
	var stats statistics

	/*****************************
		STATS EVALUATION
	******************************/
	c := make(chan bool)
	go countDifferentInteractions(c, &stats)
	go calculateTotalTimeOfInteractions(c, &stats)
	go calculateLongestSequences(c, &stats)
	go calculateMinMaxMeanTimeBetweenInteractions(c, &stats)
	<-c
	<-c
	<-c
	<-c
	close(c)
	log.Println("-------------- ENDING EVALUATIONS --------------")
	w.Header().Set("content-type", "application/json")
	statsJSON, _ := json.Marshal(stats)
	w.Write(statsJSON)
}
