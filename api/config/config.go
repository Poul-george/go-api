package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type config struct {
	MySQL  MySQL  `mapstructure:"mysql"`
	Server Server `mapstructure:"server"`
}

var once sync.Once
var cfg config

//go:embed default/local.yml
var defaultLocal []byte

func getConfig() config {
	once.Do(func() {
		viper.SetConfigType("yml")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		def := defaultLocal

		// 設定ファイルを読み込みます
		err := viper.ReadConfig(bytes.NewBuffer(def))
		if err != nil {
			fmt.Println("Failed to read yml file:", err)
			os.Exit(1) //プログラムを終了する関数
		}

		cfg = config{}
		if err := viper.Unmarshal(&cfg); err != nil {
			panic(err)
		}
	})
	return cfg

}

func GetConfig() config {
	return getConfig()
}
