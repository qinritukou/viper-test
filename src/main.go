package main

import "com.orangeman/viper-test/src/config"

func main() {
	dbReader := config.GetConfig().Database.Reader
	println(dbReader.Hostname)
	println(dbReader.Username)
	println(dbReader.Password)
	println(dbReader.DBName)
}
