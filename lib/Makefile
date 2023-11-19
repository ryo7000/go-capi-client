client: main.cpp libclient.so
	@g++ -I. -L. -Wl,-rpath=./ -o client main.cpp -lclient

libclient.so: main.go
	@go build -buildmode=c-shared -o libclient.so main.go

clean:
	rm -rf client libclient.so libclient.h
