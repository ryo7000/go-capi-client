package com.example.sotest;

import com.sun.jna.Library;
import com.sun.jna.Native;

public interface Client extends Library {
    Client INSTANCE = Native.load("client", Client.class);

    int geterrors();
    long genrand();
    String get(String latitude, String longitude);
    void gofree(String res);
}
