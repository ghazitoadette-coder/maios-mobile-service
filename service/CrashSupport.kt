package com.maios.service

import android.util.Log
import java.io.PrintWriter
import java.io.StringWriter

/**
 * MAIOS Mobile Service - Maimai Edition
 * Kotlin Crash Support Handler
 */
object CrashSupport {
    private const val TAG = "MAIOS_CRASH_SUPPORT"

    fun initializeGlobalCrashHandler() {
        val defaultHandler = Thread.getDefaultUncaughtExceptionHandler()
        
        Thread.setDefaultUncaughtExceptionHandler { thread, throwable ->
            val sw = StringWriter()
            throwable.printStackTrace(PrintWriter(sw))
            val stackTraceString = sw.toString()

            Log.e(TAG, "=== MAIOS CRASH TERDETEKSI (Maimai Edition) ===")
            Log.e(TAG, "Thread: ${thread.name}")
            Log.e(TAG, "Pesan Error: ${throwable.localizedMessage}")
            Log.e(TAG, "Stack Trace:\n$stackTraceString")

            // Tindakan pemulihan aman atau meneruskan ke handler bawaan sistem
            if (defaultHandler != null) {
                defaultHandler.uncaughtException(thread, throwable)
            } else {
                System.exit(1)
            }
        }
        
        Log.i(TAG, "Global Crash Support Kotlin berhasil diaktifkan!")
    }
}
