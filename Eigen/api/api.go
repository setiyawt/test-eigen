package api

import (
	"fmt"
	"myproject/service"
	"net/http"
)

type API struct {
	userService    service.UserService
	sessionService service.SessionService
	bookService    service.BookService
	borrowService  service.BorrowService
	mux            *http.ServeMux
}

func NewAPI(userService service.UserService, sessionService service.SessionService, bookService service.BookService, borrowService service.BorrowService) API {
	mux := http.NewServeMux()
	api := API{
		userService,
		sessionService,
		bookService,
		borrowService,
		mux,
	}

	//USERS
	mux.Handle("/user/register", api.Post(http.HandlerFunc(api.Register)))             // register
	mux.Handle("/user/login", api.Post(http.HandlerFunc(api.Login)))                   // login
	mux.Handle("/user/logout", api.Get(api.Auth(http.HandlerFunc(api.Logout))))        // logout
	mux.Handle("/user/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllUser)))) // mengambil semua user yang ada
	mux.Handle("/user/get-borrow", api.Get(api.Auth(http.HandlerFunc(api.GetAllMembersWithBorrowedCount))))

	//BUKU
	mux.Handle("/book/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllBook)))) // mengambil semua buku yang ada
	mux.Handle("/book/get", api.Get(api.Auth(http.HandlerFunc(api.FetchBookById))))    // mengambil semua buku berdasarkan ID
	mux.Handle("/book/add", api.Post(api.Auth(http.HandlerFunc(api.StoreBook))))       //menambahkan buku
	mux.Handle("/book/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateBook))))    //mengupdate buku
	mux.Handle("/book/delete", api.Delete(http.HandlerFunc(api.DeleteBook)))           //menghapus buku

	//BORROWED
	mux.Handle("/borrow/get-all", api.Get(api.Auth(http.HandlerFunc(api.FetchAllBorrow)))) // mengambil buku yang dipinjam
	mux.Handle("/borrow/get", api.Get(api.Auth(http.HandlerFunc(api.FetchBorrowById))))    // mengambil semua buku berdasarkan ID
	mux.Handle("/borrow/add", api.Post(api.Auth(http.HandlerFunc(api.StoreBorrow))))       //menambahkan buku
	mux.Handle("/borrow/update", api.Put(api.Auth(http.HandlerFunc(api.UpdateBorrow))))    //mengupdate buku
	mux.Handle("/borrow/delete", api.Delete(http.HandlerFunc(api.DeleteBorrow)))           //menghapus buku

	// mux.Handle("/book/return", api.Post(api.Auth(http.HandlerFunc(api.ReturnBook))))   //mengambil buku yang dikembalikan
	// mux.Handle("/book/check", api.Get(api.Auth(http.HandlerFunc(api.CheckBook))))      //mengambil/mengecek buku yang tersedia
	// mux.Handle("/user/check", api.Get(api.Auth(http.HandlerFunc(api.CheckUser))))      //mengambil/mengecek user yang meminjam buku
	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}
