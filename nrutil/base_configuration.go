package nrutil

// BaseConfiguration holds base conifiguration items for application
type BaseConfiguration struct {
	NrBaseSyntheticsAPIURL string
	DefaultConfigFileName  string
	DefaultConfig          Config
}

// GetBaseConfiguration returns struct with values
func GetBaseConfiguration() *BaseConfiguration {
	return &BaseConfiguration{
		NrBaseSyntheticsAPIURL: "https://synthetics.newrelic.com/synthetics/api/v3/",
		DefaultConfigFileName:  ".nrutil.yml",
		DefaultConfig: Config{
			NrAdminKey: "<your-admin-key>",
			BasePath:   "~/nrsynthetics",
			SyntheticMonitors: []string{
				"uuid-of-monitor-1-23456",
				"uuid-of-monitor-2-34567",
				"uuid-of-monitor-n-opqrs",
			},
		},
	}
}
