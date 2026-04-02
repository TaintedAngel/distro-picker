package server

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"

	"distropicker/internal/engine"
)

// New creates an http.Handler wired to the recommendation engine
// and serving embedded static assets.
func New(assets fs.FS) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/questions", handleQuestions)
	mux.HandleFunc("GET /api/profiles", handleProfiles)
	mux.HandleFunc("POST /api/recommend", handleRecommend)

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

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("json encode: %v", err)
	}
}
