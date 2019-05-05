package nrutil

// BaseConfiguration holds base conifiguration items for application
type BaseConfiguration struct {
	NrBaseSyntheticsAPIURL string
}

// GetBaseConfiguration returns struct with values
func GetBaseConfiguration() *BaseConfiguration {
	return &BaseConfiguration{
		NrBaseSyntheticsAPIURL: "https://synthetics.newrelic.com/synthetics/api/v3/",
	}
}
