package common

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const SPACE = " "

var DB *gorm.DB

type dbConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
	TimeZone string
}

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(GetDns()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func GetDbConfig() dbConfig {
	var conf dbConfig
	conf.host = os.Getenv("DB_HOST")
	conf.user = os.Getenv("POSTGRES_USER")
	conf.password = os.Getenv("POSTGRES_PASSWORD")
	conf.dbname = os.Getenv("POSTGRES_DB")
	conf.port = os.Getenv("DB_PORT")
	conf.sslmode = "disable"
	conf.TimeZone = "Asia/Shanghai"
	return conf
}

func GetDns() string {
	conf := GetDbConfig()
	return "host=" + conf.host + SPACE +
		"user=" + conf.user + SPACE +
		"password=" + conf.password + SPACE +
		"dbname=" + conf.dbname + SPACE +
		"port=" + conf.port + SPACE +
		"sslmode=" + conf.sslmode + SPACE +
		"TimeZone=" + conf.TimeZone
}
