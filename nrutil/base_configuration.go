package nrutil

// BaseConfiguration holds base conifiguration items for application
type BaseConfiguration struct {
	NrBaseSyntheticsAPIURL string
	DefaultConfigFileName  string
	DefaultConfigYaml      string
}

// GetBaseConfiguration returns struct with values
func GetBaseConfiguration() *BaseConfiguration {
	return &BaseConfiguration{
		NrBaseSyntheticsAPIURL: "https://synthetics.newrelic.com/synthetics/api/v3/",
		DefaultConfigFileName:  ".nrutil.yml",
		DefaultConfigYaml:      "LS0tCm5yYWRtaW5rZXk6IDx5b3VyLWFkbWluLWtleT4KYmFzZXBhdGg6IH4vbnJzeW50aGV0aWNzCnN5bnRoZXRpY21vbml0b3JzOgogIC0gZ3VpZC1vZi1tb25pdG9yLTEtMjM0NTYKICAtIGd1aWQtb2YtbW9uaXRvci0yLTM0NTY3CiAgLSBndWlkLW9mLW1vbml0b3Itbi1vcHFycyAK",
	}
}
