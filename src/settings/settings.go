package settings

import (
	"fmt"
	"github.com/spf13/viper"
	"task-solver/client/src/constants"
)

func LoadSettings() error {
	viper.SetConfigName(constants.ClientSettingsFileName)
	viper.SetConfigType(constants.ClientSettingsFileType)
	viper.AddConfigPath(constants.ClientSettingsPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("cannot read client settings: %w", err)
	}

	return nil
}
