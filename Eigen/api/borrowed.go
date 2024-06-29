package api

import (
	"encoding/json"
	"myproject/model"
	"net/http"
	"strconv"
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

func (api *API) GetAllMembersWithBorrowedCount(w http.ResponseWriter, r *http.Request) {
	borrow, err := api.borrowService.GetAllMembersWithBorrowedCount()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

// func (api *API) BorrowBook(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Query().Get("id")
// 	if id == "" {
// 		http.Error(w, "Missing book ID", http.StatusBadRequest)
// 		return
// 	}

// 	bookID, err := strconv.Atoi(id)
// 	if err != nil {
// 		http.Error(w, "Invalid book ID", http.StatusBadRequest)
// 		return
// 	}

// 	book, err := api.bookService.FetchByID(bookID)
// 	if err != nil {
// 		http.Error(w, "Error fetching book", http.StatusInternalServerError)
// 		return
// 	}

// 	if book.Status == "Borrowed" {
// 		http.Error(w, "Book is already borrowed", http.StatusBadRequest)
// 		return
// 	}

// 	book.Status = "Borrowed"
// 	err = api.bookService.Update(bookID, book)
// 	if err != nil {
// 		http.Error(w, "Error updating book status", http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	err = json.NewEncoder(w).Encode(book)
// 	if err != nil {
// 		http.Error(w, "Error encoding response", http.StatusInternalServerError)
// 		return
// 	}
// }

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
