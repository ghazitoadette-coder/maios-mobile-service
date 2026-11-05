package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// CrashHandlerMiddleware menangkap panic / error tak terduga agar server tidak down
func CrashHandlerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Catat detail error dan stack trace ke log sistem
				log.Printf("[MAIOS CRASH DETECTED] Panic terjadi: %v\n", err)
				log.Printf("[STACK TRACE]:\n%s\n", debug.Stack())

				// Kirim respons error yang aman ke klien alih-alih server mati
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"status": "error", "message": "Terjadi kesalahan sistem internal, namun CrashHandler berhasil mengamankan server (Maimai Edition)."}`)
			}
		}()

		// Lanjutkan eksekusi handler berikutnya jika aman
		next(w, r)
	}
}

// Contoh fungsi tes yang memicu panic untuk menguji handler
func TestCrashHandler(w http.ResponseWriter, r *http.Request) {
	panic("Simulasi error kritis pada modul MMS!")
}
