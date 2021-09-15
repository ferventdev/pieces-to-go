package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

const TWO_SPACES_INDENT = "  "

var startedAt time.Time

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/info", info)

	s := &http.Server{
		Addr:        ":8080",
		ReadTimeout: 1 * time.Second,
		Handler:     mux,
	}
	startedAt = time.Now()
	log.Fatalln(s.ListenAndServe())
}

func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func info(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	uptime := now.Sub(startedAt)

	b, err := json.MarshalIndent(
		appInfo{
			true,
			timeInfo{
				startedAt.Format(time.RFC3339),
				now.Format(time.RFC3339),
				int64(uptime.Seconds()),
				float32(uptime.Hours())},
		},
		"",
		TWO_SPACES_INDENT,
	)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(b))
}
