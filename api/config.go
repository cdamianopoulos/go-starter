package api

import (
	"fmt"
	"go-starter/pkg/aws"
	"go-starter/pkg/cuckoo"
	"go-starter/pkg/db"
	"go-starter/pkg/spiderman"
	"go-starter/separateRepos/utl"
	"time"

	"github.com/creasty/defaults"
)

// These constants could be set at compile time.
const (
	EnvPrefix = "API_NAME" // Optional
	Name      = "api-name"
)

type Config struct {
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
func (c *Config) Print() {
	fmt.Println("Port:", c.Port)
	fmt.Println("NewRelic:", c.NewRelic)
	fmt.Println("HeaderTimeout:", c.HeaderTimeout)

	// Values are comma separated by default, e.g: CUCKOO_SERVICES=1,2 >> []string{"1","2"}
	fmt.Println("Cuckoo.PaymentServices:", utl.Sprint(c.Cuckoo.PaymentServices))
	fmt.Println("Cuckoo.Timeout:", c.Cuckoo.Timeout)
	fmt.Println("Cuckoo.IsSandboxEnv:", c.Cuckoo.IsSandboxEnv)
	fmt.Println("Database.DBName:", c.Database.DBName)
	fmt.Println("Database.Username:", c.Database.Username)
	fmt.Println("Database.Password:", c.Database.Password)
	fmt.Println("Database.HostPort:", c.Database.HostPort)
	fmt.Println("Database.MaxConnLifetime:", c.Database.MaxConnLifetime)
	fmt.Println("AWS.Region:", c.AWS.Region)
	fmt.Println("Toggle.PayId:", c.Toggle.PayId)
	fmt.Println("Toggle.AsyncVACreation:", c.Toggle.AsyncVACreation)
	fmt.Println("SMClient.HostURL:", c.SMClient.HostURL)
	fmt.Println("SMClient.Timeout:", c.SMClient.Timeout)
	fmt.Println("SQSClient.QueueURL:", c.SQSClient)
	fmt.Println("WAPIClient.QueueURL:", c.WAPIClient)
}

// UnmarshalYAML sets default values defined in the Config struct tags before unmarshalling the YAML file.
// This method is needed when:
// - Loading YAML configuration WITHOUT envconfig.MustProcess or envconfig.Process (which has implemented the default struct tags).
// - And "gopkg.in/yaml.v3" doesn't support default struct tags (`default:"my_value"`) to load default values,
//
// This method is not required when using ONLY config.MustLoadEnvYaml, config.MustLoadYamlEnv, envconfig.Process or envconfig.MustProcess.
// An alternative to writing this function is to define a variables and hardcode the default values, however the implementation isn't as elegant.
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	// Set the default values
	err = defaults.Set(c)
	if err != nil {
		return
	}

	type config Config
	// Converting to a new type to prevent reflect panicking from using an unaddressable value. //stackoverflow.com/a/56080478
	return unmarshal((*config)(c))
}
