package config

import (
	"bufio"
	"os"
	"reflect"
)

func SetupServerConfigure() (*RunServerConfig, error) {

	serverConfig := RunServerConfig{}

	err := serverConfig.ReadConfig()
	if reflect.TypeOf(err) == reflect.TypeOf(NoConfigurationFoundError{}) {
		scanner := bufio.NewScanner(os.Stdin)
		print("No Configuration found! Must create configuration file, Create one now? (y/N): ")
		scanner.Scan()
		msg := scanner.Text()

		if msg == "Y" || msg == "y" {
			serverValues := ServerConfigValues{}

			print(" - Enter Host: ")
			scanner.Scan()
			msg := scanner.Text()
			serverValues.Host = string(msg)

			print(" - Enter Port: ")
			scanner.Scan()
			msg = scanner.Text()
			serverValues.Port = string(msg)

			print(" - Enter DSN: ")
			scanner.Scan()
			msg = scanner.Text()
			serverValues.DSN = string(msg)

			err := serverConfig.InitialConfig(&serverValues)
			if err != nil {
				return nil, err
			}

			return &serverConfig, nil
		}

		return nil, ConfigurationError{Message: "Initialization Configuration refused!"}

	} else if err != nil {
		return nil, err
	}

	return &serverConfig, nil
}
