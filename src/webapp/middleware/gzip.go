package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type GzipMiddleware struct {
	Next http.Handler
}

func (gm *GzipMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Handles if this is the last piece in the chain of middlewares
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	encodings := r.Header.Get("Accept-Encoding")
	// Findout if the requester is actually going to process gzip compression.
	// If not pass to next RequestHandler
	if !strings.Contains(encodings, "gzip") {
		gm.Next.ServeHTTP(w, r)
		return
	}
	// Process compression
	w.Header().Add("Content-Encoding", "gzip")
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()
	var rw http.ResponseWriter
	if pusher, ok := w.(http.Pusher); ok {
		rw = gzipPusherResponseWriter{
			gzipResponseWriter: gzipResponseWriter{
				ResponseWriter: w,
				Writer: gzipWriter,
			},
			Pusher: pusher,
		}
	} else {
		rw = gzipResponseWriter{
			ResponseWriter: w,
			Writer: gzipWriter,
		}
	}
	
	gm.Next.ServeHTTP(rw, r)
}

type gzipResponseWriter struct {
	http.ResponseWriter
	io.Writer
}

type gzipPusherResponseWriter struct {
	gzipResponseWriter
	http.Pusher
}

func (grw gzipResponseWriter) Write(data []byte) (int, error) {
	return grw.Writer.Write(data)
}