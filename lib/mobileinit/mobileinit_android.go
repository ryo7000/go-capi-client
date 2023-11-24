package mobileinit

/*
#cgo LDFLAGS: -landroid -llog

#include <android/log.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"log"
	"unsafe"
)

var ctag = C.CString("GoLog")

type infoWriter struct{}

func (infoWriter) Write(p []byte) (n int, err error) {
	cstr := C.CString(string(p))
	C.__android_log_write(C.ANDROID_LOG_INFO, ctag, cstr)
	C.free(unsafe.Pointer(cstr))
	return len(p), nil
}

func init() {
	log.SetOutput(infoWriter{})
	// android logcat includes all of log.LstdFlags
	log.SetFlags(log.Flags() &^ log.LstdFlags)
}
