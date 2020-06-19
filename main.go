package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/mjd/bee-api-gs/models"
	_ "github.com/mjd/bee-api-gs/routers"
	"github.com/mjd/bee-api-gs/util"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

// init behaves like an object constructor
func init() {

	// 1 - Enable debug
	orm.Debug = false

	// 2 - Register object with Beego ORM
	orm.RegisterModel(new(models.Wine))

	// 3 - Fetch database properties stored as YAML, decode secrets
	connStr, err := util.FetchYAML()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// 4 - Register posgres driver and db
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", connStr)

	// 5 - Create and load DB with sample data, can comment after initial load
	models.CreateDb()
	models.LoadDb()
}

func main() {

	// Load Beego framework
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
