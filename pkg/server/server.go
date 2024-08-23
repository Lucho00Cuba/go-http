package server

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Map of valid HTTP status codes
var validStatusCodes = map[int]bool{
	100: true, 101: true, 102: true, // 1xx Informational
	200: true, 201: true, 202: true, 203: true, 204: true, 205: true, 206: true, 207: true, 208: true, 226: true, // 2xx Success
	300: true, 301: true, 302: true, 303: true, 304: true, 305: true, 306: true, 307: true, 308: true, // 3xx Redirection
	400: true, 401: true, 402: true, 403: true, 404: true, 405: true, 406: true, 407: true, 408: true, 409: true, 410: true, 411: true, 412: true, 413: true, 414: true, 415: true, 416: true, 417: true, 418: true, 421: true, 422: true, 423: true, 424: true, 426: true, 428: true, 429: true, 431: true, 451: true, // 4xx Client Error
	500: true, 501: true, 502: true, 503: true, 504: true, 505: true, 506: true, 507: true, 508: true, 510: true, // 5xx Server Error
}

type Server struct {
	Port string
}

// Create a new server instance
func NewServer(port string) *Server {
	return &Server{Port: port}
}

// Run the server
func (s *Server) Run() {
	http.HandleFunc("/", s.logRequest(s.handler))
	log.Printf("Server running on port %s\n", s.Port)
	if err := http.ListenAndServe(":"+s.Port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Middleware to log each request
func (s *Server) logRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log request start details
		log.Printf("Started %s %s for %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Custom ResponseWriter to capture the status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Call the original handler
		handler(lrw, r)

		// Log request completion details
		duration := time.Since(start).Milliseconds()
		log.Printf("Completed %s %s with %d in %dms", r.Method, r.URL.Path, lrw.statusCode, duration)
	}
}

// Handle incoming requests and respond with the specified status code
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")

	if len(parts) > 0 {
		statusCodeStr := parts[0]
		statusCode, err := strconv.Atoi(statusCodeStr)
		if err == nil && isValidStatusCode(statusCode) {
			w.WriteHeader(statusCode)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

// Check if the status code is valid
func isValidStatusCode(code int) bool {
	return validStatusCodes[code]
}

// Custom ResponseWriter to capture the status code
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Capture the status code and forward it to the original ResponseWriter
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
