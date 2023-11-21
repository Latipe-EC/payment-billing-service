package gorm

import (
	"fmt"
	"latipe-payment-billing-service/app/config"
	"log"
)

func NewMySQLConnection(configuration *config.Config) Gorm {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&loc=Local",
		configuration.DB.Mysql.UserName,
		configuration.DB.Mysql.Password,
		configuration.DB.Mysql.Host,
		configuration.DB.Mysql.Port,
		configuration.DB.Mysql.Database,
	)
	cfg := Config{
		DSN:             dataSourceName,
		MaxOpenConns:    configuration.DB.Mysql.MaxOpenConns,
		MaxIdleConns:    configuration.DB.Mysql.MaxIdleConns,
		ConnMaxLifetime: configuration.DB.Mysql.ConnMaxLifetime,
		ConnMaxIdleTime: configuration.DB.Mysql.ConnMaxIdleTime,
		DBType:          "mysql",
		Debug:           true,
	}
	conn, err := New(cfg)
	if err != nil {
		panic(err)
	}

	log.Printf("[%s] Gorm has created database connection", "INFO")
	return conn

}
