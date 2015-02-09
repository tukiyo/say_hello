package main

import (
	"syscall"
	"unsafe"
)

const (
	SND_SYNC      = 0x0000
	SND_NODEFAULT = 0x0002
	SND_FILENAME  = 0x00020000
)

var (
	libwinmm      = syscall.NewLazyDLL("winmm")
	procPlaySound = libwinmm.NewProc("PlaySoundW")
)

func playsnd(filename string) error {
	_, r1, err := procPlaySound.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(filename))),
		SND_SYNC|SND_NODEFAULT|SND_FILENAME)
	if r1 == 0 {
		return err
	}
	return nil
}
