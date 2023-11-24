package main

/*
#include <stdlib.h>
#include "codes.h"
*/
import "C"

import (
	"crypto/rand"
	"io"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"unsafe"

	_ "github.com/ryo7000/go-capi-client/mobileinit"
)

var messages = map[C.int][]byte{
	C.CODE_ERROR_A: []byte("error a\x00"),
	C.CODE_ERROR_B: []byte("error b\x00"),
}

//export get
func get(latitude *C.char, longitude *C.char) *C.char {
	log.Println("[start] get")
	endpoint := "https://api.open-meteo.com/v1/forecast"

	u, err := url.Parse(endpoint)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set("latitude", C.GoString(latitude))
	q.Set("longitude", C.GoString(longitude))
	q.Set("hourly", "temperature_2m")
	q.Set("timezone", "Asia/Tokyo")
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var body []byte
	switch res.StatusCode {
	case http.StatusOK:
		body, err = io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
	default:
		panic("http error")
	}

	log.Println("[end] get")
	// C.CStringは内部でmallocしている
	return C.CString(string(body))
}

//export gofree
func gofree(cs *C.char) {
	C.free(unsafe.Pointer(cs))
}

//export geterrors
func geterrors() C.Codes {
	return C.CODE_ERROR_B
}

//export errtostr
func errtostr(code C.int) uintptr {
	msg := messages[code]
	return uintptr(unsafe.Pointer(&msg[0]))
}

//export genrand
func genrand() C.longlong {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		panic(err)
	}
	return C.longlong(n.Int64())
}

//export panic_test
func panic_test() {
	panic("panic test")
}

func main() {}
