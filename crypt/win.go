package crypt

import (
	"syscall"
	"unsafe"
)

var (
	crypt32 = syscall.NewLazyDLL("crypt32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	cryptUnprotectDataFunc = crypt32.NewProc("CryptUnprotectData")
	localFreeFunc = kernel32.NewProc("LocalFree")
)

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func newBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}

	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func cryptUnprotectData(data []byte) []byte {
	var outblob DATA_BLOB

	r, _, err := cryptUnprotectDataFunc.Call(uintptr(unsafe.Pointer(newBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		panic(err)
	}

	defer localFreeFunc.Call(uintptr(unsafe.Pointer(outblob.pbData)))

	return outblob.ToByteArray()
}
