package config

type ConfigurationError struct {
	Message string
}

func (config ConfigurationError) Error() string {
	return config.Message
}

type NoConfigurationFoundError struct {
	Message string
}

func (config NoConfigurationFoundError) Error() string {
	return config.Message
}
