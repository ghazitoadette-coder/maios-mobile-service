package main

import (
	"fmt"
	"net/http"
)

// SetupRoutes mendaftarkan seluruh handler endpoint untuk service MAIOS
func SetupRoutes() {
	fmt.Println("[service/router] Mengonfigurasi rute-rute utama...")

	// Endpoint 1: Cek status service
	http.HandleFunc("/api/v1/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"status": "online", "system": "MAIOS Mobile Service"}`)
	})

	// Endpoint 2: Informasi Akun MAIOS
	http.HandleFunc("/api/v1/account", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"user": "Ghazi", "email": "ghazi@maios.com", "role": "Owner"}`)
	})

	// Endpoint 3: Modul Maimai Service
	http.HandleFunc("/api/v1/maimai", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"module": "Maimai Edition", "active_asset": "hatsunemiku.mai"}`)
	})

	fmt.Println("[service/router] Rute berhasil didaftarkan:")
	fmt.Println("  [GET] /api/v1/status")
	fmt.Println("  [GET] /api/v1/account")
	fmt.Println("  [GET] /api/v1/maimai")
}
