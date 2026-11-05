package com.maios.service

class MaiMaiService {
    private var isServiceRunning: Boolean = false

    fun startMaiMaiService() {
        isServiceRunning = true
        println("MAIOS Mobile Service (Maimai Edition) berhasil dijalankan!")
    }

    fun stopMaiMaiService() {
        isServiceRunning = false
        println("MAIOS Mobile Service dihentikan.")
    }
}
