package itemsorders

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GenericResponse returns timestamp and status to itemsorder in JSON
type GenericResponse struct {
	StatusCode int16
	Status     string
	TimeStamp  time.Time
}

// ItemsorderIndex => returns a slice of all itemsorders in the database
func ItemsorderIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usrs, err := AllItemsorders()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usrs)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)

}

// ItemsorderCreate creates a itemsorder instance in the database
func ItemsorderCreate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usr := Itemsorder{}
	json.NewDecoder(req.Body).Decode(&usr)

	err := usr.CreateItemsorder()
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "Itemsorder Created Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}

// ItemsorderFind finds a itemsorder instance by ID in the database
func ItemsorderFind(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	u := Itemsorder{ID: id}

	usr, err := u.FindItemsorder()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usr)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)
}

// ItemsorderUpdate Updates a itemsorder instance in the database
func ItemsorderUpdate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	usr := Itemsorder{ID: id}
	json.NewDecoder(req.Body).Decode(&usr)

	err = usr.UpdateItemsorder()
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 202,
		Status:     "Itemsorder Updated Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(res, "%s\n", rj)
}

// ItemsorderDelete deletes a itemsorder instance from the database
func ItemsorderDelete(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	u := Itemsorder{ID: id}

	err = u.DeleteItemsorder()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "Itemsorder succesfully DELETED",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - ITEMSORDERS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}
