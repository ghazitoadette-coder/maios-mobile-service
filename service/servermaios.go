package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// MaiMaiServer mendefinisikan struktur server utama MAIOS
type MaiMaiServer struct {
	Port   string
	Status string
}

func (s *MaiMaiServer) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "success", "message": "MAIOS Mobile Service (Maimai Edition) Server is running smoothly!", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
}

func main() {
	server := &MaiMaiServer{
		Port:   ":8080",
		Status: "Active",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", server.HandleHealthCheck)

	fmt.Printf("[MAIOS Server] Memulai server Maimai Edition pada port %s...\n", server.Port)
	fmt.Println("[MAIOS Server] Sistem anti-error diaktifkan. Menunggu koneksi...")

	// Menjalankan server dengan penanganan error jika port gagal diikat
	err := http.ListenAndServe(server.Port, mux)
	if err != nil {
		log.Fatalf("[MAIOS Server Error] Server gagal berjalan: %v", err)
	}
}
