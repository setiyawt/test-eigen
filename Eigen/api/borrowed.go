package api

import (
	"encoding/json"
	"myproject/model"
	"net/http"
	"strconv"
	"time"
)

func (api *API) FetchAllBorrow(w http.ResponseWriter, r *http.Request) {
	borrow, err := api.borrowService.FetchAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(borrow)
}

func (api *API) FetchBorrowById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	borrow, err := api.borrowService.FetchByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(borrow)
}

func (api *API) StoreBorrow(w http.ResponseWriter, r *http.Request) {
	var borrow model.Borrowed

	err := json.NewDecoder(r.Body).Decode(&borrow)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.borrowService.Store(&borrow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(borrow)
}

func (api *API) UpdateBorrow(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var borrow model.Borrowed
	err = json.NewDecoder(r.Body).Decode(&borrow)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	err = api.borrowService.Update(idInt, &borrow)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(borrow)
}

func (api *API) DeleteBorrow(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = api.borrowService.Delete(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.SuccessResponse{Message: "borrow successfully deleted"})
}

func (api *API) ReturnBook(w http.ResponseWriter, r *http.Request) {
	codeBook := r.URL.Query().Get("code_book")
	codeMember := r.URL.Query().Get("code_member")
	returnDate := time.Now()

	if codeBook == "" || codeMember == "" {
		http.Error(w, "Missing book or member code", http.StatusBadRequest)
		return
	}

	err := api.borrowService.ReturnBook(codeBook, codeMember, returnDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book returned successfully"))
}
