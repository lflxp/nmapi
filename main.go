package main

import (
	"encoding/json"
	"fmt"
	// _ "github.com/lflxp/nmapi/routers"
	"github.com/lflxp/nmapi/pkg"
	// "github.com/astaxie/beego"
)

// func main() {
// 	if beego.BConfig.RunMode == "dev" {
// 		beego.BConfig.WebConfig.DirectoryIndex = true
// 		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
// 	}
// 	beego.Run()
// }

func main() {
	s := pkg.NewScanner()
	s.SetArgs("-A").SetArgs("192.168.40.228")
	ss, err := s.Parse()
	if err != nil {
		panic(err)
	}
	a123, err := json.Marshal(ss)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(a123))
}
