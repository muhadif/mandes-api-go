package app

import (
	"fmt"
	"github.com/muhadif/mandes/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
)

type Dependency struct {
	Database *gorm.DB
	HTTP     *http.Client
	Cfg      config.Config
}

func InitDependency() *Dependency {
	cfg := config.LoadConfig()
	return &Dependency{
		Cfg:      cfg,
		Database: initDB(cfg),
		HTTP:     initHTTP(),
	}
}

func initDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySqlUsername, cfg.MySqlPassword, cfg.MySqlHost, cfg.MySqlPort, cfg.MySqlDatabase)

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		panic(any(err))
	}

	return db
}

func initHTTP() *http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	//t.MaxIdleConns = config.HTTPMaxIdleConn
	//t.MaxConnsPerHost = config.HTTPMaxConnPerHost
	//t.MaxIdleConnsPerHost = config.HTTPMaxIdleConnPerHost

	return &http.Client{
		//Timeout:   config.HTTPTimeout * time.Second,
		Transport: t,
	}
}

func (d *Dependency) Destroy() {
	d.HTTP.CloseIdleConnections()
}
