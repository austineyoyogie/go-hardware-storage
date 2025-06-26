package servers

import (
	"flag"
	"log"
	"github.com/austineyoyogie/go-hardware-store/database"
	"github.com/austineyoyogie/go-hardware-store/api-products/models"
	"github.com/austineyoyogie/go-hardware-store/routers"
)

var (
	port = flag.Int("p", 8000, "set port")
	resetTables = flag.Bool("rt", false, "reset tables")
)

func Run() {
	flag.Parse()
	routers.RouterHandler()
	if *port != 8000 && *resetTables {
		createSuperTestTables()
	}	
}

// Create a table for debug testing
func createSuperTestTables() {
	db := database.Connect()
	if db != nil {
		defer db.Close()
	}

	tx := db.Begin()
	err := tx.Debug().DropTableIfExists(models.Product{}, &models.Category{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().CreateTable(&models.Category{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().CreateTable(&models.Product{}).Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Debug().Model(&models.Product{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE").Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
}
