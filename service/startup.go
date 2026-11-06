package main

import (
	"fmt"
	"os"
)

// InitSystemStartup bertugas membaca folder init/ dan menyiapkan service
func InitSystemStartup() {
	fmt.Println("==================================================")
	fmt.Println("  MAIOS MOBILE SERVICE - STARTUP SYSTEM")
	fmt.Println("  Owner: Ghazi | Account: ghazi@maios.com")
	fmt.Println("==================================================")

	// 1. Cek keberadaan folder service/init/
	initFolderPath := "init"
	if _, err := os.Stat(initFolderPath); os.IsNotExist(err) {
		fmt.Println("[startup/warning] Folder 'init/' belum ditemukan, membuat folder baru...")
		os.Mkdir(initFolderPath, 0755)
	} else {
		fmt.Println("[startup/status] Folder 'init/' terdeteksi: OK")
	}

	// 2. Memeriksa file script .rc di dalam folder init/
	rcFilePath := "init/init.rc"
	if _, err := os.Stat(rcFilePath); err == nil {
		fmt.Printf("[startup/init] Berhasil memuat file configuration: %s\n", rcFilePath)
	} else {
		fmt.Printf("[startup/init] Menunggu file %s dimasukkan ke folder init/\n", rcFilePath)
	}

	// 3. Status Siap
	fmt.Println("[startup/ready] Modul startup berhasil dijalankan!")
}
