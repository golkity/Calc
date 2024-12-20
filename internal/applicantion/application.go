package application

import (
	"bufio"
	"github.com/golkity/Calc/config"
	"github.com/golkity/Calc/internal/http/handler"
	"github.com/golkity/Calc/pkg/calc"
	"log"
	"net/http"
	"os"
	"strings"
)

type Application struct {
	config *config.ServerConfig
}

func New(configPath string) *Application {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	return &Application{config: cfg}
}

func (a *Application) Run() {
	log.Println("CLI mode started. Type 'exit' to quit.")
	reader := bufio.NewReader(os.Stdin)

	for {
		log.Print("Input expression: ")
		text, err := reader.ReadString('\n')

		if err != nil {
			log.Println("Failed to read input:", err)
			continue
		}

		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("Exiting application.")
			return
		}

		result, err := calc.Calc(text)
		if err != nil {
			log.Printf("Error: %v\n", err)
		} else {
			log.Printf("Result: %s = %f\n", text, result)
		}
	}
}

func (a *Application) RunServer() {
	http.HandleFunc("/api/v1/calculate", handler.CalcHandler)
	addr := ":" + a.config.Port
	log.Printf("Server is running on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
