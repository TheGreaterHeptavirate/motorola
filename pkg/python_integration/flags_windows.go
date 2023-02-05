//go:build windows
// +build windows

package python

//#cgo CFLAGS: -I/usr/x86_64-w64-mingw32/sys-root/mingw/include/python3.10 -I/usr/x86_64-w64-mingw32/sys-root/mingw/include/python3.10  -Wno-unused-result -Wsign-compare  -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions --param=ssp-buffer-size=4 -DNDEBUG -g -O3 -Wall
//#cgo LDFLAGS: -static /usr/x86_64-w64-mingw32/sys-root/mingw/lib/libpython3.10.dll.a -L/usr/x86_64-w64-mingw32/sys-root/mingw/lib -lpthread -lm -lversion -lshlwapi -lm
//#cgo CPPFLAGS:
import "C"
