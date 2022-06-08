/**
 * @Author: realpeanut
 * @Date: 2022/6/7 23:58
 */
package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file")
	help = pflag.BoolP("help", "h", false, "Show this help message")
)

func main() {
	demo9()
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

func demo3() {
	// 监听和重新读取配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("config file changed", in.Name)
	})
}

func demo4() {
	// 显式的设置配置项
	viper.Set("user.username", "tom")
}

func demo5() {
	// 设置环境变量
	// Viper 读取环境变量是区分大小写的
	_ = os.Setenv("VIPER_USER_SECRET_ID", "abcdefg")
	_ = os.Setenv("VIPER_TEST", "abcdefg")

	viper.AutomaticEnv()           // 加载环境变量
	viper.SetEnvPrefix("VIPER")    // 设置前缀
	fmt.Println(viper.Get("test")) // 读取test 实际读取的是 VIPER_TEST
}

func demo6() {
	//读取配置
	viper.AddConfigPath("$HOME/go/src/repeat/viper/")
	viper.SetConfigName("config")                // 配置文件名称（没有文件扩展名）
	viper.SetConfigType("json")                  // 配置类型
	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(viper.GetString("datastore.metric.host"))
	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
}

func demo7() {
	//反序列化
	type config struct {
		Port    int
		Name    string
		PathMap string `mapstructure:"path_map"`
	}
	var C config
	err := viper.Unmarshal(&C)
	if err != nil {
		return
	}
	fmt.Println(C)
}

func demo8() {
	// 如果想要解析那些键本身就包含.(默认的键分隔符）的配置，则需要修改分隔符：
	v := viper.NewWithOptions(viper.KeyDelimiter("::"))

	v.SetDefault("chart::values", map[string]interface{}{
		"ingress": map[string]interface{}{
			"annotations": map[string]interface{}{
				"traefik.frontend.rule.type":                 "PathPrefix",
				"traefik.ingress.kubernetes.io/ssl-redirect": "true",
			},
		},
	})

	type config struct {
		Chart struct {
			Values map[string]interface{}
		}
	}

	var C config

	err := v.Unmarshal(&C)
	if err != nil {
		return
	}
}

func demo9() {
	// 序列化成字符串
	viper.AddConfigPath("$HOME/go/src/repeat/viper/")
	viper.SetConfigName("conf") // 配置文件名称（没有文件扩展名）
	viper.SetConfigType("yaml") // 配置类型
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	fmt.Println(string(bs))
}
