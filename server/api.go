package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/EUDAT-GEF/DEX/def"
	"github.com/gorilla/mux"
)

const (
	// ServiceName is used for HTTP API
	ServiceName = "DEX"
	// Version defines the api version
	Version = "0.1.0"
)

const apiRootPath = "/api"

// Server is a master struct for serving HTTP API requests
type Server struct {
	Server http.Server
	tmpDir string
}

// NewServer creates a new Server
func NewServer(cfg def.ServerConfig, tmpDir string) (*Server, error) {
	tmpDir, err := def.MakeTmpDir(tmpDir)
	if err != nil {
		return nil, def.Err(err, "creating temporary directory failed")
	}

	server := &Server{
		Server: http.Server{
			Addr: cfg.Address,
			// timeouts seem to trigger even after a correct read
			// ReadTimeout: 	cfg.ReadTimeoutSecs * time.Second,
			// WriteTimeout: 	cfg.WriteTimeoutSecs * time.Second,
		},
		tmpDir: tmpDir,
	}

	routes := map[string]func(http.ResponseWriter, *http.Request){
		"GET /":     server.infoHandler,
		"GET /info": server.infoHandler,

		"POST /events": server.newEventHandler,
	}

	router := mux.NewRouter()

	apirouter := router.PathPrefix(apiRootPath).Subrouter()
	for mp, handler := range routes {
		methodPath := strings.SplitN(mp, " ", 2)
		apirouter.HandleFunc(methodPath[1], handler).Methods(methodPath[0])
	}

	server.Server.Handler = router
	return server, nil
}

// Start starts a new http listener
func (s *Server) Start() error {
	return s.Server.ListenAndServe()
}

func (s *Server) infoHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	Response{w}.Ok(jmap("service", ServiceName, "version", Version))
}

func (s *Server) newEventHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	var data map[string]interface{}
	if r.Body == nil {
		http.Error(w, "Request body needed", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Printf("%#v", data)

	Response{w}.Ok(jmap("status", 0))
}
