package main

import (
	"github.com/d0lim/turnstile/internal/di"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Error loading .env file")
	}

	app, err := di.InitializeApp()
	if err != nil {
		panic(err)
	}

	logrus.Fatal(app.Listen(":8080"))
}
