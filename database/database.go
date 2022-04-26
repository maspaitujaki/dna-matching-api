package database

import "fmt"

type Config struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Shanghai", config.Host, config.User, config.Password, config.Name, config.Port)
	return connectionString
}
