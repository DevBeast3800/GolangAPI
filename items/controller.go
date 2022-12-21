package items

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GenericResponse returns timestamp and status to item in JSON
type GenericResponse struct {
	StatusCode int16
	Status     string
	TimeStamp  time.Time
}

// ItemIndex => returns a slice of all items in the database
func ItemIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usrs, err := AllItems()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usrs)
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)

}

// ItemCreate creates a item instance in the database
func ItemCreate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usr := Item{}
	json.NewDecoder(req.Body).Decode(&usr)

	err := usr.CreateItem()
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "Item Created Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}

// ItemFind finds a item instance by ID in the database
func ItemFind(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	u := Item{ID: id}

	usr, err := u.FindItem()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usr)
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)
}

// ItemUpdate Updates a item instance in the database
func ItemUpdate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	usr := Item{ID: id}
	json.NewDecoder(req.Body).Decode(&usr)

	err = usr.UpdateItem()
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 202,
		Status:     "Item Updated Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(res, "%s\n", rj)
}

// ItemDelete deletes a item instance from the database
func ItemDelete(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	u := Item{ID: id}

	err = u.DeleteItem()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "Item succesfully DELETED",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - ITEMS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}
