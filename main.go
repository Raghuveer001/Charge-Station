// Go package
package main

import (
	"encoding/json"

	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Create a structure  for Add Charging
type AddChargingStation struct {
	StationID    int
	EnergyOutput string
	Type         string
}

// Create a Structure for Start Charging
type StartCharging struct {
	StationID              int
	VehicleBatteryCapacity string
	CurrentVehicleCharge   string
	ChargingStartTime      string
}

var chargeStation []AddChargingStation

var startcharging []StartCharging

func RequestHandler() {
	router := mux.NewRouter()
	router.HandleFunc("/chst", chargingstation).Methods("POST")
	router.HandleFunc("/allchst", ChargingStations)
	router.HandleFunc("/stch", startchargings).Methods("POST")
	router.HandleFunc("/allstch", Startcharging)

	http.ListenAndServe("127.0.0.1:8000", router)

}

// Making a Post Request for Add Charge Station
func chargingstation(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var chst AddChargingStation
	json.Unmarshal(reqBody, &chst)
	chargeStation = append(chargeStation, chst)
	json.NewEncoder(w).Encode(chst)

}

// Making a Get Request for Add Charge Station
func ChargingStations(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chargeStation)
}

var i int = 0

//Making a Post Request for Start Charinig
func startchargings(w http.ResponseWriter, r *http.Request) {

	if i < len(chargeStation) {
		i++
		reqBody, _ := ioutil.ReadAll(r.Body)
		var stch StartCharging
		json.Unmarshal(reqBody, &stch)
		startcharging = append(startcharging, stch)
		json.NewEncoder(w).Encode(stch)

	} else {

	}
}

//Making a Get Request for Start Charginig

func Startcharging(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startcharging)

}

func main() {
	chargeStation = []AddChargingStation{}
	startcharging = []StartCharging{}

	RequestHandler()

}
