package conf

import (
	"os"

	"github.com/spf13/viper"
)

type Server struct {
	ConnMaxLifetime string `yaml:"connMaxLifetime"`
	Dsn             string `yaml:"dsn"`
	MaxIdleConns    int32  `yaml:"maxIdleConns"`
	MaxOpenConns    int32  `yaml:"maxOpenConns"`
	Tag             string `yaml:"tag"`
	Type            string `yaml:"type"`
}
type Config struct {
	Port    int32
	Servers []Server
}

func defaultConfig() Config {
	return Config{
		Port: 8000,
	}
}
func loadConfig() Config {
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
	config := defaultConfig()
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	// fmt.Println(viper.Get("servers"))
	return config
}

var conf = loadConfig()

func Get() *Config {
	return &conf
}

func GetServerByTag(tag string) *Server {
	for _, s := range conf.Servers {
		if s.Tag == tag {
			return &s
		}
	}
	return nil
}

func Refresh() *Config {
	conf = loadConfig()
	return &conf
}
