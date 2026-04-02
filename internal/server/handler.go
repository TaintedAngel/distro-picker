package server

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"sync"
	"time"

	"distropicker/internal/engine"
)

// Heartbeat tracks the last time the browser checked in.
var heartbeat struct {
	mu   sync.Mutex
	last time.Time
}

func touchHeartbeat() {
	heartbeat.mu.Lock()
	heartbeat.last = time.Now()
	heartbeat.mu.Unlock()
}

// LastHeartbeat returns when the browser last pinged.
func LastHeartbeat() time.Time {
	heartbeat.mu.Lock()
	defer heartbeat.mu.Unlock()
	return heartbeat.last
}

// New creates an http.Handler wired to the recommendation engine
// and serving embedded static assets.
func New(assets fs.FS) http.Handler {
	touchHeartbeat() // initialize so the watchdog doesn't fire immediately

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/questions", handleQuestions)
	mux.HandleFunc("GET /api/profiles", handleProfiles)
	mux.HandleFunc("POST /api/recommend", handleRecommend)
	mux.HandleFunc("POST /api/heartbeat", handleHeartbeat)

	webFS, err := fs.Sub(assets, "web")
	if err != nil {
		log.Fatalf("embedded asset error: %v", err)
	}
	mux.Handle("GET /", http.FileServer(http.FS(webFS)))

	return mux
}

func handleQuestions(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, engine.Questions)
}

func handleProfiles(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, engine.Profiles)
}

func handleRecommend(w http.ResponseWriter, r *http.Request) {
	var req engine.RecommendRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if len(req.Answers) == 0 && req.Profile == "" {
		http.Error(w, "answers or profile required", http.StatusBadRequest)
		return
	}
	writeJSON(w, engine.Recommend(req))
}

func handleHeartbeat(w http.ResponseWriter, _ *http.Request) {
	touchHeartbeat()
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("json encode: %v", err)
	}
}
