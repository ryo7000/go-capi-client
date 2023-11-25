package mobileinit

/*
#cgo LDFLAGS: -landroid -llog

#include <android/log.h>
#include <stdlib.h>
#include <string.h>
*/
import "C"

import (
	"log/slog"
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
	logger := slog.New(slog.NewTextHandler(infoWriter{}, &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// android logcat includes time
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			return a
		},
	}))
	slog.SetDefault(logger)
}
