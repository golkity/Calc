package main

import application "github.com/golkity/Calc/internal/applicantion"

func main() {
	app := application.New("config/config.json")
	app.RunServer()
}
