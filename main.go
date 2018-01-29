// Copyright 2013 Beego Samples authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// This sample is about using long polling and WebSocket to build a web-based chat room based on beego.
package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"

	"github.com/nomin-project/nomin-project.github.io/controllers"
)

var (
	servingCertFile = os.Getenv("SERVING_CERT")
	servingKeyFile  = os.Getenv("SERVING_KEY")
)

func main() {
	beego.Info(beego.AppName, beego.AppConfig.String("version"))

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting workdirectory:", err)
		os.Exit(133)
	}

	beego.SetViewsPath(wd)

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/healthz", &controllers.HealthzController{})

	beego.Run()
}
