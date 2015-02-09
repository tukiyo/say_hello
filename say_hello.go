package main

import (
	"fmt"
	"net/http"
)

func say_hi(w http.ResponseWriter, r *http.Request) {
	if err := playsnd("hi.wav"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("hi\n"))
}
func say_bye(w http.ResponseWriter, r *http.Request) {
	if err := playsnd("bye.wav"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("bye\n"))
}

func main() {
	http.HandleFunc("/hi", say_hi)
	http.HandleFunc("/bye", say_bye)
	fmt.Print("Listening 8080 port\n")
	http.ListenAndServe(":8080", nil)
}
