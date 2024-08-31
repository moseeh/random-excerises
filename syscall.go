package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	// Get the directory to list from command line argument or use current directory
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	// Open the directory
	fd, err := syscall.Open(dir, syscall.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("Error opening directory: %v\n", err)
		return
	}
	defer syscall.Close(fd)

	// Read directory contents
	files, err := syscall.ReadDirent(fd, 0)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Parse and print file information
	offset := 0
	for offset < len(files) {
		dirent := (*syscall.Dirent)(unsafe.Pointer(&files[offset]))
		offset += int(dirent.Reclen)

		if dirent.Ino == 0 {
			continue
		}

		name := make([]byte, dirent.Namlen)
		copy(name, (*[256]byte)(unsafe.Pointer(&dirent.Name))[:dirent.Namlen])

		filename := string(name)
		info, err := os.Stat(dir + "/" + filename)
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			continue
		}

		fileMode := info.Mode()
		modTime := info.ModTime().Format(time.RFC822)
		size := info.Size()

		fmt.Printf("%s %8d %s %s\n", fileMode, size, modTime, filename)
	}
}
