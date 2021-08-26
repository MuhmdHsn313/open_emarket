package config

type ConfigValues interface {
	GetValues() []interface{}
}

type ServerConfigValues struct {
	Host string
	Port string
	DSN  string
}

func (v ServerConfigValues) GetValues() []interface{} {
	var values []interface{}
	values[0] = v.Host
	values[1] = v.Port
	values[2] = v.DSN
	return values
}
