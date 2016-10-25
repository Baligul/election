package main

import (
	"os"
	"strconv"

	_ "github.com/Baligul/election/docs"
	_ "github.com/Baligul/election/routers"
	
	modelGroups "github.com/Baligul/election/models/groups"
	modelVoters "github.com/Baligul/election/models/voters"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=member dbname=election sslmode=disable")
	orm.RegisterModel(new(modelVoters.Account), 
					  new(modelVoters.Voter), 
					  new(modelVoters.Voter_19), 
					  new(modelVoters.Voter_20), 
					  new(modelVoters.Voter_21),
					  new(modelGroups.UserGroup))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(port)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
