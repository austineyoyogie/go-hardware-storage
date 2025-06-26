package database

import (
	"fmt"
	"log"
	"github.com/austineyoyogie/go-hardware-store/configs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var g = configs.LoadConfigs()

func Connect() *gorm.DB {	
	db, err := gorm.Open(g.DBC.Driver, g.DBC.Type+"://"+g.DBC.Username+":"+ g.DBC.Password+ 
	"@"+g.DBC.Hostname+":"+ g.DBC.Port +"/"+g.DBC.Database+"?sslmode=disable")
	if err != nil {	
		//panic(err)
		logFatal(err)
		return nil
	} else {
		fmt.Println("Database connect successfully!.")
	}
	return db
}

func logFatal(err error) {
	if err != nil {
		log.Fatal("Error database connection mode.")	
	} 
}

