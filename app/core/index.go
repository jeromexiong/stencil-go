package core

import (
	. "stencil-go/app/core/config"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *Config
	DB     *gorm.DB
	Log    *logrus.Logger
	Redis  *redis.Client
)
