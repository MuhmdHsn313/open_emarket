package config

import (
	"io/ioutil"
	"open_emarker/settings"
	"regexp"
	"strings"
)

type Config interface {
	ReadConfig() error
	InitialConfig(values *ConfigValues) error
	ValidateConfig(data string) bool
}

type RunServerConfig struct {
	Host   string
	Port   string
	DSN    string
	values map[string]string
}

func (r *RunServerConfig) InitialConfig(values *ServerConfigValues) error {

	r.DSN = values.DSN
	r.Port = values.Port
	r.Host = values.Host

	if r.values == nil {
		r.values = make(map[string]string)
	}

	r.values["HOST"] = values.Host
	r.values["PORT"] = values.Port
	r.values["DSN"] = values.DSN

	fileData := ""

	for key, value := range r.values {
		fileData += key + "=" + value + "\n"
	}

	err := ioutil.WriteFile(settings.CONFIG_FILE_PATH, []byte(fileData), 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r *RunServerConfig) ReadConfig() error {

	file, err := ioutil.ReadFile(settings.CONFIG_FILE_PATH)
	if err != nil {
		return NoConfigurationFoundError{Message: "Can't configuration open file!\n" + err.Error()}
	}

	fileData := string(file)

	if r.ValidateConfig(fileData) {
		compile := regexp.MustCompile(settings.CONFIG_VALUE_VALIDATOR)
		hosts := compile.FindAllStringSubmatch(fileData, -1)

		if r.values == nil {
			r.values = make(map[string]string)
		}

		for _, host := range hosts {
			r.values[host[1]] = host[2]
		}

		r.Host = r.values["HOST"]
		r.Port = r.values["PORT"]
		r.DSN = r.values["DSN"]
	} else {
		return ConfigurationError{Message: "Can't read configuration values!"}
	}

	return nil
}

// ValidateConfig Using for validating the data coming from settings.CONFIG_FILE_PATH for initialize configuration
// struct.
func (r *RunServerConfig) ValidateConfig(data string) bool {
	values := strings.Split(data, "\n")

	for _, str := range values {
		match, err := regexp.MatchString(settings.CONFIG_VALUE_VALIDATOR, str)
		if (err != nil || !match) && str != "" {
			return false
		}
	}

	return true
}

// GetFullAddress return the host with port number that limited in RunServerConfig struct.
func (r *RunServerConfig) GetFullAddress() string {
	return r.Host + ":" + r.Port
}
