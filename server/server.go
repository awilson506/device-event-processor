package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	devices "github.com/gocariq/golang_data_challenge/pkg"
)

// Server
type Server struct {
	server  *http.Server
	mux     *http.ServeMux
	devices *devices.DeviceService
}

// New - get a new instance of the server
func New(db *sql.DB) *Server {
	s := &Server{
		mux:     http.NewServeMux(),
		devices: devices.New(db),
	}

	handlers := []struct {
		pattern string
		handler http.HandlerFunc
	}{
		{pattern: "/devices/get-latest-details", handler: s.getAllDevicesHandler},
		{pattern: "/devices/get-latest-details/", handler: s.getDeviceHandler},
	}

	for _, h := range handlers {
		s.mux.HandleFunc(h.pattern, h.handler)
	}

	s.server = &http.Server{
		Handler: s.mux,
		Addr:    ":8080",
	}
	return s
}

// Start - start the server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// ServeHTTP - field the server requests
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// getAllDevicesHandler - get all the latest device records
func (s *Server) getAllDevicesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.devices.GetAllLatestDeviceDetails())
}

// getDeviceHandler - get a specific device
func (s *Server) getDeviceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// TODO: add validation on this param
	deviceId := strings.TrimPrefix(r.URL.Path, "/devices/get-latest-details/")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.devices.GetLatestDeviceDetails(deviceId))
}
