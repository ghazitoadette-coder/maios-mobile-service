package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// AppConfig merepresentasikan konfigurasi global MAIOS Mobile Service (Maimai Edition)
type AppConfig struct {
	AppName      string `json:"app_name"`
	Version      string `json:"version"`
	Environment  string `json:"environment"`
	Port         string `json:"port"`
	ResponseRate string `json:"response_rate"`
}

// LoadGlobalConfig memuat konfigurasi utama atau menggunakan nilai default (fallback)
func LoadGlobalConfig() (*AppConfig, error) {
	fmt.Println("[MAIOS Config] Memuat konfigurasi global Maimai Edition...")

	// Default configuration khusus Maimai Edition
	config := &AppConfig{
		AppName:      "MAIOS Mobile Service",
		Version:      "2.0.0-maimai",
		Environment:  "production",
		Port:         ":8080",
		ResponseRate: "Ultra-Fast",
	}

	return config, nil
}

// SaveConfigAsJSON menyimpan konfigurasi ke dalam bentuk format JSON jika diperlukan
func SaveConfigAsJSON(config *AppConfig, filename string) error {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("[MAIOS Config] Berhasil menyimpan konfigurasi ke %s\n", filename)
	return nil
}
