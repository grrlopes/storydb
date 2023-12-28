package helper

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func Echobuffer(cmd string, env string) {
	fd, err := syscall.Open(env, syscall.O_RDWR, 0)
	if err != nil {
		fmt.Println("Error opening Device:", err)
		os.Exit(1)
	}

	defer syscall.Close(fd)

	for i := 0; i < len(cmd); i++ {
		char := cmd[i]
		b := []byte{char}
		syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), syscall.TIOCSTI, uintptr(unsafe.Pointer(&b[0])))
	}

	fmt.Printf("%+v\n\n", " ---")
}

