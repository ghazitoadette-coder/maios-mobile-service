package com.maios.service;

/**
 * MAIOS Mobile Service - Maimai Edition
 * Bridge Layer for Java & Android Compatibility
 */
public class MaiMaiBridge {

    private MaiMaiService maiMaiService;
    private boolean isBridgeConnected;

    public MaiMaiBridge() {
        this.maiMaiService = new MaiMaiService();
        this.isBridgeConnected = false;
    }

    public void initializeBridge() {
        System.out.println("[MAIOS Bridge] Menginisialisasi Maimai Bridge (Java)...");
        maiMaiService.startMaiMaiService();
        this.isBridgeConnected = true;
        System.out.println("[MAIOS Bridge] Jembatan Java-Kotlin berhasil terhubung!");
    }

    public boolean isConnected() {
        return isBridgeConnected;
    }

    public void terminateBridge() {
        if (isBridgeConnected) {
            maiMaiService.stopMaiMaiService();
            this.isBridgeConnected = false;
            System.out.println("[MAIOS Bridge] Jembatan terputus.");
        }
    }
}
