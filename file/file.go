package file

import (
	"io/ioutil"
	"io/fs"
	"os"
)

var (
	Roaming = os.Getenv("APPDATA")
	Local = os.Getenv("LOCALAPPDATA")
	Temp = os.Getenv("TEMP")
)

func PathExists(p string) bool {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return false
	}

	return true
}

func ListDir(p string) []fs.DirEntry {
	entries, err := os.ReadDir(p)
	if err != nil {
		panic(err)
	}

	return entries
}

func ReadData(p string) []byte {
	data, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	return data
} 

func CopyFile(src, dst string) {
	data, err := ioutil.ReadFile(src)
    if err != nil {
        panic(err)
    }

    err = ioutil.WriteFile(dst, data, 0644)
    if err != nil {
        panic(err)
    }
}

func DeleteFile(p string) {
	err := os.Remove(p)
    if err != nil {
        panic(err)
    }
}
