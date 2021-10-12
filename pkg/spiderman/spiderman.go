package spiderman

import "time"

type ClientConfig struct {
	HostURL string        `envconfig:"SM_HOST_URL"`
	Timeout time.Duration `envconfig:"HTTP_CLIENT_TIMEOUT" default:"45s"`
}
