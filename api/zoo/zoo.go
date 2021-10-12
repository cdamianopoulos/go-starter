package zoo

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"go-starter/separateRepos/cnst/hdr"
	"go-starter/separateRepos/cnst/mime"
)

// URLPattern is used in router.go
const URLPattern = "/zoo"

// Status returns the status and quantity of animals in the zoo.
func Status(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(hdr.ContentType, mime.JSON)

	_, err := fmt.Fprintf(w, `{"zoo": "closed", "visitors": %d, "animal_qty": 11}`, time.Now().Local().Second())
	if err != nil {
		logrus.Error("writing to stream:", err)
	}
}
