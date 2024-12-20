package main

import application "Calc/internal/applicantion"

func main() {
	app := application.New("config/config.json")
	app.RunServer()
}
