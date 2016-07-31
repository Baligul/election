package main

import (
    "os"
    "strconv"
    
	_ "github.com/Baligul/election/docs"
	_ "github.com/Baligul/election/routers"

	"github.com/astaxie/beego"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
        
    beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(port)
        
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
    
	beego.Run()
}