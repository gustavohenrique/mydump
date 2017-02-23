package main

import (
	"fmt"
	"github.com/gustavohenrique/barkup"
	"github.com/zpatrick/go-config"
)

func main() {
	file := config.NewJSONFile("config.json")
	c := config.NewConfig([]config.Provider{file})
	if err := c.Load(); err != nil{
        	fmt.Println(err)
	}

	host, _ := c.StringOr("host", "localhost")
	port, _ := c.StringOr("port", "3306")
	database, _ := c.StringOr("database", "mysql")
	user, _ := c.StringOr("user", "root")
	password, _ := c.StringOr("password", "root")
	mysql := &barkup.MySQL{
		Host:     host,
		Port:     port,
		DB:       database,
		User:     user,
		Password: password,
	}

	err := mysql.Export().To("dumps", nil)
	if err != nil {
		fmt.Println(err)
	}
}
