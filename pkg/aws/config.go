package aws

type Config struct {
	Region   string `envconfig:"AWS_REGION" default:"ap-southeast-2"`
	Endpoint string `envconfig:"AWS_ENDPOINT"`
}
