package common

import "github.com/spf13/viper"

var (
	CONFIG_SERVER_URL = "setting.server"
	CONFIG_GROUP_IDS  = "setting.groupids"
	CONFIG_TIMEOUT    = "setting.timeout"
	SERVER_URL        string
	GROUP_IDS         []string
	TIMEOUT           int
)

func InitValues() {
	SERVER_URL = viper.GetString(CONFIG_SERVER_URL)
	GROUP_IDS = viper.GetStringSlice(CONFIG_GROUP_IDS)
	TIMEOUT = viper.GetInt(CONFIG_TIMEOUT)
}
