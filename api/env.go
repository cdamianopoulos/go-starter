package api

import (
	"fmt"
	"time"

	"github.com/creasty/defaults"
	"go-starter/pkg/aws"
	"go-starter/pkg/cuckoo"
	"go-starter/pkg/db"
	"go-starter/pkg/spiderman"
	"go-starter/separateRepos/utl"
)

// These constants could be set at compile time.
const (
	EnvPrefix = "API_NAME" // Optional
	Name      = "api-name"
)

type Env struct {
	Address       string
	Port          uint16        `default:"3000"`
	NewRelic      string        `split_words:"true" default:"xyz889"`
	HeaderTimeout time.Duration `split_words:"true" default:"37s"`
	Database      db.Config
	Cuckoo        cuckoo.ClientConfig
	SQSClient     string `envconfig:"ASYNC_REQUEST_QUEUE_URL"`
	WAPIClient    string `envconfig:"WAPI_QUEUE_URL"`
	Toggle        struct {
		PayId           bool `envconfig:"ENABLE_PAYID"`
		AsyncVACreation bool `envconfig:"ENABLE_ASYNC_VA_CREATION"`
	}
	AWS      aws.Config
	SMClient spiderman.ClientConfig
}

// Print writes the field values stored in e to standard output.
func (e *Env) Print() {
	fmt.Println("Port:", e.Port)
	fmt.Println("NewRelic:", e.NewRelic)
	fmt.Println("HeaderTimeout:", e.HeaderTimeout)

	// Values are comma separated by default, e.g: CUCKOO_SERVICES=1,2 >> []string{"1","2"}
	fmt.Println("Cuckoo.PaymentServices:", utl.Sprint(e.Cuckoo.PaymentServices))
	fmt.Println("Cuckoo.Timeout:", e.Cuckoo.Timeout)
	fmt.Println("Cuckoo.IsSandboxEnv:", e.Cuckoo.IsSandboxEnv)
	fmt.Println("Database.DBName:", e.Database.DBName)
	fmt.Println("Database.Username:", e.Database.Username)
	fmt.Println("Database.Password:", e.Database.Password)
	fmt.Println("Database.HostPort:", e.Database.HostPort)
	fmt.Println("Database.MaxConnLifetime:", e.Database.MaxConnLifetime)
	fmt.Println("AWS.Region:", e.AWS.Region)
	fmt.Println("Toggle.PayId:", e.Toggle.PayId)
	fmt.Println("Toggle.AsyncVACreation:", e.Toggle.AsyncVACreation)
	fmt.Println("SMClient.HostURL:", e.SMClient.HostURL)
	fmt.Println("SMClient.Timeout:", e.SMClient.Timeout)
	fmt.Println("SQSClient.QueueURL:", e.SQSClient)
	fmt.Println("WAPIClient.QueueURL:", e.WAPIClient)
}

// UnmarshalYAML sets default values defined in the Env struct tags before unmarshalling the YAML file.
// This method is needed because "gopkg.in/yaml.v2" doesn't support the `default` struct tags.
func (e *Env) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	// Set the default values
	err = defaults.Set(e)
	if err != nil {
		return
	}

	type env Env
	return unmarshal((*env)(e))
}
