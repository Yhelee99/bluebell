package setting

import (
	"bluebell/mod"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Init() {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		zap.L().Fatal("读取配置文件失败！")
		return
	}

	//转义时还是传指针
	if err = viper.Unmarshal(&mod.Conf); err != nil {
		zap.L().Fatal("转义到结构体失败！")
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		zap.L().Info("配置文件被修改！")
		viper.Unmarshal(mod.Conf)
	})
}
