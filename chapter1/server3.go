package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//http.HandleFunc("/", handler3)

	handler4 := func(w http.ResponseWriter, r *http.Request){
		if keys, ok := r.URL.Query()["cycles"]; ok && len(keys) > 0 {
			cycles, err := strconv.Atoi(keys[0])
			if err != nil {
				log.Print("er r not nil")
				Lissajous(w, 5)
			} else {
				log.Print("made it to " + keys[0])
				Lissajous(w, float64(cycles))
			}
		} else {
			Lissajous(w, 5)
		}
	}
	http.HandleFunc("/", handler4)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "For[%q] = %q\n", k, v)
	}
}
