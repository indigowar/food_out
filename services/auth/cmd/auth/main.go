package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Auth service"))
	})

	_ = http.ListenAndServe(":80", nil)
}
