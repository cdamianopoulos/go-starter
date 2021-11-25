package zoo

import (
	"fmt"
	"go-starter/separateRepos/cnst/hdr"
	"go-starter/separateRepos/cnst/mime"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// URLPattern is used in router.go
const URLPattern = "/zoo"

// Zoo represents the zoo structure.
type Zoo struct {
	Name      string `json:"closed"`
	Visitors  string `json:"visitors"`
	AnimalQty int    `json:"animal_qty"`
}

// Status returns the status and quantity of animals in the zoo.
func Status(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set(hdr.ContentType, mime.JSON)

	// Hardcode the response for better performance than JSON marshaling.
	_, err := fmt.Fprintf(w, `{"zoo":"closed","visitors":%d,"animal_qty":11}`, time.Now().Local().Second())
	if err != nil {
		logrus.Error("writing to stream:", err)
	}
}
