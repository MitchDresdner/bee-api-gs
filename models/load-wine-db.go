package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// Uses Beego ORM to create table
func CreateDb() error {

	// Database alias.
	name := "default"

	// Drop table and re-create (change to false after created).
	force := true

	// Print log.
	verbose := true

	// Beego ORM function to create the table
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func LoadDb() error {

	// Some sample data to add to our DB
	wines := []Wine{
		{
			Product:     "SOMMELIER SELECT",
			Description: "Old vine Cabernet Sauvignon",
			Price:       159.99,
		},
		{
			Product:     "MASTER VINTNER",
			Description: "Pinot Noir captures luscious aromas",
			Price:       89.99,
		},
		{
			Product:     "WINEMAKER'S RESERVE",
			Description: "Merlot featuring complex flavors of cherry",
			Price:       84.99,
		},
		{
			Product:     "ITALIAN SANGIOVESE",
			Description: "Sangiovese grape is famous for its dry, bright cherry character",
			Price:       147.99,
		},
	}

	// Insert static entries into database
	for idx := 0; idx < len(wines); idx++ {
		w := wines[idx]
		_, err := AddWine(w)
		if err != nil {
			return err
		}
	}

	return nil
}
