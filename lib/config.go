package lib

import (
	"github.com/BaizeSec/Pusher/common"
	"github.com/BaizeSec/Pusher/pkg/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Group struct {
	GroupId int
	Desc    string
}

/**
初始化配置文件
*/
func InitConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	// setting
	viper.SetDefault(common.CONFIG_SERVER_URL, "http://127.0.0.1:5700")
	viper.SetDefault(common.CONFIG_GROUP_IDS, []string{})
	viper.SetDefault(common.CONFIG_TIMEOUT, 5)

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("Read config failed, Please restart pusher!")
	}
	viper.WriteConfigAs("config.toml")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("Config file changed!")
	})
}
