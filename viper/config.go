/**
 * @Author: realpeanut
 * @Date: 2022/6/7 23:58
 */
package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file")
	help = pflag.BoolP("help", "h", false, "Show this help message")
)

func main() {
	demo2()
}

func demo1() {
	//设置默认值
	viper.SetDefault("ContentDir", "content")
}

func demo2() {
	//Viper 可以读取配置文件来解析配置，支持 JSON、TOML、YAML、YML、Properties、Props、Prop、HCL、Dotenv、Env 格式的配置文件。
	//Viper 支持搜索多个路径，并且默认不配置任何搜索路径，将默认决策留给应用程序
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigFile("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/go/src/repeat/viper/")
		viper.SetConfigName("conf") // 配置文件名称（没有文件扩展名）
		viper.SetConfigType("yaml") // 配置类型
	}
	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
}
