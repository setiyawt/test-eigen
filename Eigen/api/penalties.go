package api

import (
	"encoding/json"
	"myproject/model"
	"net/http"
	"strconv"
)

func (api *API) FetchAllPenalties(w http.ResponseWriter, r *http.Request) {
	penalties, err := api.penaltiesService.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(penalties)
}

func (api *API) FetchPenaltiesById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	penalties, err := api.penaltiesService.FetchByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(penalties)
}

func (api *API) BorrowPenalties(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing penalties ID", http.StatusBadRequest)
		return
	}

	penaltiesID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid penalties ID", http.StatusBadRequest)
		return
	}

	penalties, err := api.penaltiesService.FetchByID(penaltiesID)
	if err != nil {
		http.Error(w, "Error fetching penalties", http.StatusInternalServerError)
		return
	}

	if penalties.PenaltyActive == false {
		http.Error(w, "Member does not have penalties", http.StatusBadRequest)
		return
	}

	penalties.PenaltyActive = false
	err = api.penaltiesService.Update(penaltiesID, penalties)
	if err != nil {
		http.Error(w, "Error updating penalties status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(penalties)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func (api *API) StorePenalties(w http.ResponseWriter, r *http.Request) {
	var penalties model.Penalties

	err := json.NewDecoder(r.Body).Decode(&penalties)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.penaltiesService.Store(&penalties)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(penalties)
}

func (api *API) UpdatePenalties(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var penalties model.Penalties
	err = json.NewDecoder(r.Body).Decode(&penalties)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.penaltiesService.Update(idInt, &penalties)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(penalties)
}

func (api *API) DeletePenalties(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.penaltiesService.Delete(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "Book successfully deleted"})
}
