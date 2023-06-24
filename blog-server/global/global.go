package global

import (
	"Blog/config"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config   *config.Config
	DB       *gorm.DB
	Logger   *logrus.Logger
	MysqlLog logger.Interface
	Redis    *redis.Client
)
