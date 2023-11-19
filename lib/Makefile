client: main.cpp libclient.so
	@g++ -I. -L. -Wl,-rpath=./ -o client main.cpp -lclient

libclient.so: main.go
	@go build -buildmode=c-shared -o libclient.so main.go

clean:
	rm -rf client libclient.so libclient.h

arm:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -buildmode=c-shared -o libclient.so main.go
