package promotions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GenericResponse returns timestamp and status to promotion in JSON
type GenericResponse struct {
	StatusCode int16
	Status     string
	TimeStamp  time.Time
}

// PromotionIndex => returns a slice of all promotions in the database
func PromotionIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usrs, err := AllPromotions()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usrs)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)

}

// PromotionCreate creates a promotion instance in the database
func PromotionCreate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	usr := Promotion{}
	json.NewDecoder(req.Body).Decode(&usr)

	err := usr.CreatePromotion()
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "Promotion Created Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}

// PromotionFind finds a promotion instance by ID in the database
func PromotionFind(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	u := Promotion{ID: id}

	usr, err := u.FindPromotion()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(usr)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", uj)
}

// PromotionUpdate Updates a promotion instance in the database
func PromotionUpdate(res http.ResponseWriter, req *http.Request) {
	if req.Method != "PUT" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	usr := Promotion{ID: id}
	json.NewDecoder(req.Body).Decode(&usr)

	err = usr.UpdatePromotion()
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 202,
		Status:     "Promotion Updated Succesfully",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(res, "%s\n", rj)
}

// PromotionDelete deletes a promotion instance from the database
func PromotionDelete(res http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
		http.Error(res, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	vars := mux.Vars(req)
	i := vars["id"]
	id, err := strconv.Atoi(i)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(404), http.StatusNotFound)
		return
	}

	u := Promotion{ID: id}

	err = u.DeletePromotion()

	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	r := GenericResponse{
		StatusCode: 201,
		Status:     "Promotion succesfully DELETED",
		TimeStamp:  time.Now(),
	}

	rj, err := json.Marshal(r)
	if err != nil {
		log.Printf("[ERROR - PROMOTIONS - CONTROLLER] => %v", err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprintf(res, "%s\n", rj)
}
