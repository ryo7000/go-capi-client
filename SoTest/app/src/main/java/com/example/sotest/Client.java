package com.example.sotest;

import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.ptr.ByteByReference;

public interface Client extends Library {
    Client INSTANCE = Native.load("client", Client.class);

    int geterrors();
    long genrand();
    ByteByReference get(String latitude, String longitude);
    void gofree(ByteByReference res);
}
