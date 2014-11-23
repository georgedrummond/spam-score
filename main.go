package main

import (
	"github.com/codegangsta/martini"
	"github.com/georgedrummond/spamscore/message"
)

func MessageViewHandler(params martini.Params) string {
	message := params["uuid"]
	return message
}

func MessageCreateHandler(params martini.Params) string {
	return "todo"
}

func main() {
	m := martini.Classic()

	m.Get("/api/v1/messages/:uuid", MessageViewHandler)
	m.Post("/api/v1/messages", MessageCreateHandler)

	m.Run()
}
