package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	millisecs := float64(time.Now().UnixNano()) / float64(1000000)
	urlPathString := r.URL.Path[1:]
	urlPathInt, err := strconv.ParseInt(urlPathString, 10, 0)

	if err == nil {
		if urlPathInt > 86400 {
			fmt.Fprint(w, "woah, do you REALLY want to wait that long?")
		} else {
			duration := time.Duration(urlPathInt) * time.Second
			time.Sleep(duration)
			fmt.Fprintf(w, "app resp time: %fs",
				(float64(time.Now().UnixNano())/float64(1000000)-millisecs)/1000)
		}
	} else {
		fmt.Fprintf(w, "you didn't specify an int as the pause value, silly..")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
