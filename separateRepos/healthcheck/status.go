package healthcheck

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"go-starter/separateRepos/cnst/hdr"
	"go-starter/separateRepos/cnst/mime"
)

// URLPattern is used in router.go
const URLPattern = "/status"

// HealthCheck represents the health check structure.
type HealthCheck struct {
	Status string `json:"status"`
}

// Status returns a JSON health check endpoint.
func Status(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(hdr.ContentType, mime.JSON)

	// Hardcode the response for better performance than JSON marshaling.
	_, err := fmt.Fprint(w, `{"status":"healthy"}`)
	if err != nil {
		logrus.Error("writing to stream:", err)
	}
}
