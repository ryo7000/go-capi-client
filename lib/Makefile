NDK_ROOT = $(HOME)/Android/Sdk/ndk/26.1.10909125

client: main.cpp libclient.so
	@g++ -I. -L. -Wl,-rpath=./ -o client main.cpp -lclient

libclient.so: main.go
	@go build -buildmode=c-shared -o libclient.so main.go

clean:
	rm -rf client libclient.so libclient.h

android:
	CGO_ENABLED=1 GOOS=android GOARCH=arm64 CC=$(NDK_ROOT)/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android21-clang go build -buildmode=c-shared -o ../SoTest/app/src/main/jniLibs/arm64-v8a/libclient.so main.go

android_amd64:
	CGO_ENABLED=1 GOOS=android GOARCH=amd64 CC=$(NDK_ROOT)/toolchains/llvm/prebuilt/linux-x86_64/bin/x86_64-linux-android21-clang go build -buildmode=c-shared -o ../SoTest/app/src/main/jniLibs/x86_64/libclient.so main.go
