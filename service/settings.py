"""
MAIOS Mobile Service - Maimai Edition
Global Python Settings & Environment Configuration
"""

import os

# --- Informasi Dasar Proyek ---
PROJECT_NAME = "maios-mobile-service"
EDITION = "Maimai Edition"
VERSION = "2.0.0"
AUTHOR = "Ghazi"

# --- Konfigurasi Server & Runtime ---
DEBUG_MODE = True
SERVER_HOST = "0.0.0.0"
SERVER_PORT = 8080
RESPONSE_RATE = "Ultra-Fast"

# --- Direktori & Path Sistem ---
BASE_DIR = os.path.dirname(os.path.abspath(__file__))
SERVICE_CORE_PATH = os.path.join(BASE_DIR, "service_core.maios")

# --- Parameter Fitur Maimai Edition ---
MAIMAI_CONFIG = {
    "bridge_sync": True,
    "crash_handler_active": True,
    "low_level_runtime": "maios-runtime.c",
    "theme_engine": "Maimai_Vibrant_Style"
}

def display_settings_info():
    print(f"[{PROJECT_NAME}] Loaded: {EDITION} (v{VERSION})")
    print(f"Target Port: {SERVER_PORT} | Mode: {RESPONSE_RATE}")

if __name__ == "__main__":
    display_settings_info()
  
