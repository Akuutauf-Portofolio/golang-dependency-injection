package main

import (
	"belajar-go-lang-restful-api/helper"
	"belajar-go-lang-restful-api/middleware"
	"net/http"
)

// membuat function new server
func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,

		// ketika auth middleware sudah ditambahkan, maka handler yang digunakan adalah-
		// router yang dibungkus dengan auth middleware
	}
}

// mendefinisikan main program
func main() {
	// membuat server dan sekaligus mendefinisikan hasil dari injector untuk dependency injectionnya
	server := InitializedServer()

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek error
	helper.PanicIfError(err)
}