package com.example.sotest;

import com.sun.jna.Native;

public class Client {
    static {
        Native.register("client");
    }

    public static native int geterrors();
}
