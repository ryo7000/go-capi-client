package com.example.sotest;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.widget.TextView;

import com.example.sotest.databinding.ActivityMainBinding;
import com.sun.jna.ptr.ByteByReference;

public class MainActivity extends AppCompatActivity {

    private ActivityMainBinding binding;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        binding = ActivityMainBinding.inflate(getLayoutInflater());
        setContentView(binding.getRoot());

        // Example of a call to a native method
        TextView tv = binding.sampleText;
        tv.setText(String.valueOf(Client.INSTANCE.genrand()));

        TextView tv2 = binding.api;
        ByteByReference ref = Client.INSTANCE.get("35.68141046761117", "139.76716771217266");
        tv.setText(ref.getPointer().getString(0));
        Client.INSTANCE.gofree(ref);
    }
}
