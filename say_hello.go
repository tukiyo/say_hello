package main

import (
    "fmt"
    "net/http"
    "runtime"
    "os/exec"
)

func say_hi(w http.ResponseWriter, r *http.Request) {
    if runtime.GOOS == "windows" {
        exec.Command("playsnd.exe", "hi.wav").Output()
    } else {
        exec.Command("open", "hi.wav").Output()
    }
    fmt.Print("hi\n")
}
func say_bye(w http.ResponseWriter, r *http.Request) {
    if runtime.GOOS == "windows" {
        exec.Command("playsnd.exe", "bye.wav").Output()
    } else {
        exec.Command("open", "bye.wav").Output()
    }
    fmt.Print("bye\n")
}

func main() {
    http.HandleFunc("/hi", say_hi)
    http.HandleFunc("/bye", say_bye)
    fmt.Print("Listening 8080 port\n")
    http.ListenAndServe(":8080", nil)
}
