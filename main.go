package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

var (
	port     = flag.Int("p", 0, "Web server port (e.g. 8080)")
	dir      = flag.String("d", ".", "Public HTML directory")
	retry    = flag.Int("r", 5, "Times to retry serving on a new port")
	provided = true
)

func init() {
	flag.Parse()
	if *port == 0 {
		*port = 8080
		provided = false
	}
}

type Handler struct {
	fs http.Handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("method=%s path=%s", r.Method, r.URL.Path)
	w.Header().Add("Cache-Control", "max-age=0")
	w.Header().Add("Last-Modified", time.Now().UTC().Format(time.RFC1123))
	h.fs.ServeHTTP(w, r)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("app=serve ")

	handler := &Handler{fs: http.FileServer(http.Dir(*dir))}
	for i := 0; i < *retry; i++ {
		path, _ := filepath.Abs(*dir)
		base := filepath.Base(path)

		log.Printf("fn=listen port=%d directory=%s", *port, base)
		http.ListenAndServe(fmt.Sprintf(":%d", *port), handler)
		msg := "port already in use"

		if provided {
			log.Fatalf("fn=listen error=%q", msg)
		}

		log.Printf("fn=listen error=%q", msg)
		*port++
	}

	log.Fatalf("fn=listen error=%q", "could not bind to a port")
}
