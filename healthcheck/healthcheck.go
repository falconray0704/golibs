package healthcheck

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/qdm12/golibs/network"
	"github.com/qdm12/golibs/server"
)

// Mode checks if the program is launched to run the
// internal healthcheck against another instance of the
// program on the same machine
func Mode(args []string) bool {
	return len(args) > 1 && args[1] == "healthcheck"
}

// Query sends an HTTP request to the other instance of
// the program, and to its internal healthcheck server.
func Query() error {
	request, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:9999", nil)
	if err != nil {
		return fmt.Errorf("Cannot build HTTP request: %w", err)
	}
	client := &http.Client{Timeout: 1 * time.Second}
	status, _, err := network.DoHTTPRequest(client, request)
	if err != nil {
		return fmt.Errorf("Cannot execute HTTP request: %w", err)
	}
	if status != 200 {
		return fmt.Errorf("HTTP status code is %d", status)
	}
	return nil
}

// Serve runs the server locally to serve the healthcheck
func Serve(rootURL string, isHealthy func() bool, stop <-chan struct{}) error {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if isHealthy() {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	return server.Serve("127.0.0.1:9999", router, stop)
}