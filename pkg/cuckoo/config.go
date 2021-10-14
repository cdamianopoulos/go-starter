package cuckoo

import (
	"time"
)

type ClientConfig struct {
	HostURL         string        `envconfig:"CUCKOO_HOST_URL"`
	PaymentServices []string      `envconfig:"CUCKOO_SERVICES"`
	Timeout         time.Duration `envconfig:"HTTP_CLIENT_TIMEOUT" default:"45s"`
	SeqGenURL       string        `envconfig:"SEQUENCE_GEN_URL"`
	SeqGenAPIKey    string        `envconfig:"SEQUENCE_GEN_API_KEY"`
	SSLCert         string        `envconfig:"CUCKOO_SSL_CERT"`
	IsSandboxEnv    isSandboxEnv  `envconfig:"FACTORY"`
}

type isSandboxEnv bool

func (ise *isSandboxEnv) Set(value string) error {
	*ise = value == "prelive" // Compare sandbox environment name.
	return nil
}
