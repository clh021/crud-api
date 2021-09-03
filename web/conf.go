package web

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigName(".crud-api")
	viper.SetConfigType("yaml")
	// 在 Home 及当前目录下面查找名为 ".crud-api" 的配置文件
	home, err := os.UserHomeDir()
	if err != nil {
		viper.AddConfigPath(home)
	}
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	testvar := viper.Get("servers")
	fmt.Println(testvar)
}
