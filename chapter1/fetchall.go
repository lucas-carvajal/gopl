package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	counter := 1
	for range os.Args[1:] {
		fmt.Println(strconv.Itoa(counter) + ". " +<-ch)
		counter += 1
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string){
	start := time.Now()
	if !strings.HasPrefix("http://", url) {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	//destination, err := os.Create(url[7:] + ".txt")
	//if err != nil {
	//	ch <- fmt.Sprint(err)
	//	return
	//}
	//defer destination.Close()
	//nbytes, err := io.Copy(destination, resp.Body)

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
