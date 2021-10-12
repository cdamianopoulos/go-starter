package main

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go-starter/pkg/aws"
	"go-starter/pkg/cuckoo"
	"go-starter/pkg/db"
	"go-starter/pkg/spiderman"
	"go-starter/separateRepos/config"
	"go-starter/separateRepos/utl"
)

const ApiName = "api-name" // This could be set at compile time.

var env Env

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

func loadConfig(envPrefix, configFile string) {
	// Load environment variables from a YAML config file.
	config.MustLoadYaml(configFile, &env)
	// Any environment variables loaded are over-riding any values defined in config.yml.
	// Move above MustLoadYaml for the reverse effect.
	envconfig.MustProcess(envPrefix, &env)

	fmt.Println("Port:", env.Port)
	fmt.Println("NewRelic:", env.NewRelic)
	fmt.Println("HeaderTimeout:", env.HeaderTimeout)

	// Values are comma separated by default, e.g: CUCKOO_SERVICES=1,2 >> []string{"1","2"}
	fmt.Printf("Cuckoo.PaymentServices: %s\n", utl.Sprint(env.Cuckoo.PaymentServices))
	fmt.Println("Cuckoo.Timeout:", env.Cuckoo.Timeout)
	fmt.Println("Cuckoo.IsSandboxEnv:", env.Cuckoo.IsSandboxEnv)
	fmt.Println("Database.DBName:", env.Database.DBName)
	fmt.Println("Database.Username:", env.Database.Username)
	fmt.Println("Database.Password:", env.Database.Password)
	fmt.Println("Database.HostPort:", env.Database.HostPort)
	fmt.Println("Database.MaxConnLifetime:", env.Database.MaxConnLifetime)
	fmt.Println("AWS.Region:", env.AWS.Region)
	fmt.Println("Toggle.PayId:", env.Toggle.PayId)
	fmt.Println("Toggle.AsyncVACreation:", env.Toggle.AsyncVACreation)
	fmt.Println("SMClient.HostURL:", env.SMClient.HostURL)
	fmt.Println("SMClient.Timeout:", env.SMClient.Timeout)
	fmt.Println("SQSClient.QueueURL:", env.SQSClient)
	fmt.Println("WAPIClient.QueueURL:", env.WAPIClient)
}
