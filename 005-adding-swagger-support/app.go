package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/ferventdev/pieces-to-go/005-adding-swagger-support/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

const TWO_SPACES_INDENT = "  "

var startedAt time.Time

type appInfo struct {
	Running  bool     `json:"running"`
	TimeInfo timeInfo `json:"time"`
}

type timeInfo struct {
	StartedAt     string  `json:"startedAt"`
	Now           string  `json:"now"`
	UptimeInSec   int64   `json:"uptimeInSec"`
	UptimeInHours float32 `json:"uptimeInHours"`
}

// @title Server with healthcheck
// @version 0.0.1
// @description This is a sample server to show the added swagger support.
// @host localhost:8080
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", check)
	mux.HandleFunc("/info", info)
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	s := &http.Server{
		Addr:        ":8080",
		ReadTimeout: 1 * time.Second,
		Handler:     mux,
	}
	startedAt = time.Now()
	log.Fatalln(s.ListenAndServe())
}

// @Summary Healthcheck endpoint
// @Description Returns 200 OK up, if all's good.
// @Produce plain
// @Success 200 {object} string
// @Router /healthcheck [get]
func check(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "up")
}

// @Summary Information endpoint
// @Description This is a current app's state endpoint, may be used for a diagnostic purpose.
// @Produce json
// @Success 200 {object} main.appInfo
// @Failure 500 {object} string
// @Router /info [get]
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
