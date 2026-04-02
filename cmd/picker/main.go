package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"distropicker"
	"distropicker/internal/server"
)

func main() {
	port := flag.Int("port", 9514, "HTTP port to listen on")
	noBrowser := flag.Bool("no-browser", false, "don't open browser automatically")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      server.New(distropicker.Assets),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("distro-picker listening on http://localhost:%d", *port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	if !*noBrowser {
		// Small delay so the server is ready before the browser hits it.
		time.AfterFunc(200*time.Millisecond, func() {
			openBrowser(fmt.Sprintf("http://localhost:%d", *port))
		})
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down…")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown: %v", err)
	}
}

func openBrowser(url string) {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", url}
	case "windows":
		args = []string{"rundll32", "url.dll,FileProtocolHandler", url}
	default:
		args = []string{"xdg-open", url}
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	if err := cmd.Start(); err != nil {
		log.Printf("could not open browser: %v", err)
	}
}
