package healthcheck

// TODO move healthcheck into a separate Git repository
import (
	"fmt"
	"go-starter/separateRepos/cnst/hdr"
	"go-starter/separateRepos/cnst/mime"
	"net/http"

	"github.com/sirupsen/logrus"
)

// URLPattern is used in router.go
const URLPattern = "/status"

// Status returns a JSON health check endpoint.
func Status(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(hdr.ContentType, mime.JSON)
	_, err := fmt.Fprint(w, `{"status":"healthy"}`)
	if err != nil {
		logrus.Error("writing to stream:", err)
	}
}
