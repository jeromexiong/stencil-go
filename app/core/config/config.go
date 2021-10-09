package config

import (
	"bytes"
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

// global var
var (
	GConfig Config
)

type (
	Config struct {
		iris.Configuration `yaml:"Configuration"`
		Keys               string `yaml:"keys"`
		Own                `yaml:"own"`
		Redis              `yaml:"redis"`
		MySQL              `yaml:"mysql"`
		Dialect            string `yaml:"dialect"`
		Production         bool   // 是否是生产环境
	}
	// Own config
	Own struct {
		Port          int      `yaml:"port"`
		JWTTimeout    int      `yaml:"JWTTimeout"`
		WebsocketPool int      `yaml:"websocketPool"`
		Domains       []string `yaml:"domains"`
		IgnoreURLs    []string `yaml:"ignoreURLs"`
	}
	// Redis config
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
		PoolSize int    `yaml:"poolSize"`
		Required bool   `yaml:"required"`
	}
	// MySQL config
	MySQL struct {
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		Host         string `yaml:"host"`
		Port         int    `yaml:"port"`
		Database     string `yaml:"database"`
		Charset      string `yaml:"charset"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		Required     bool   `yaml:"required"`
	}
)

// Mysql数据库连接url
func (config Config) MysqlUrl() string {
	info := config.MySQL
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", info.User, info.Password, info.Host, info.Port, info.Database, info.Charset)
}

// Postgre数据库连接url
func (config Config) PostgreUrl() string {
	info := config.MySQL
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", info.Host, info.Port, info.User, info.Database, info.Password)
}

// [asset] 二进制资源，开发环境不需要打包二进制文件; [production] 是否是生产环境
func New(asset func(name string) ([]byte, error), production bool) *Config {
	c := Config{Production: production}

	if err := c.setConfig(asset); err != nil {
		log.Fatalln("初始化配置文件出错", err.Error())
	}
	if !production {
		c.watchConfig()
	}
	GConfig = c
	return &GConfig
}

func (c *Config) setConfig(asset func(name string) ([]byte, error)) error {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.SetConfigName("config.default")
	if err := viper.ReadInConfig(); err == nil {
		if c.Production {
			viper.SetConfigName("config.prod")
			if err := viper.MergeInConfig(); err != nil {
				return err
			}
		}
	}

	// 读取二进制配置
	if c.Production {
		var (
			data []byte
			err  error
		)
		if data, err = asset("config/config.default.yaml"); err == nil {
			if err = viper.ReadConfig(bytes.NewBuffer(data)); err != nil {
				return err
			}
		}
		if data, err = asset("config/config.prod.yaml"); err == nil {
			if err = viper.MergeConfig(bytes.NewBuffer(data)); err != nil {
				return err
			}
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
		return err
	}

	log.Printf("%#v", c.MySQL)
	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("配置文件修改更新: %s\n", e.Name)
	})
}
