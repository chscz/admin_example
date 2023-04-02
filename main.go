package main

import (
	"example/router"
	"example/template"
)

func main() {
	e := router.SetRouter()
	e.Static("/web_content", "web_content")
	e.Renderer = template.SetTemplate()

	e.Logger.Fatal(e.Start(":1323"))
}
