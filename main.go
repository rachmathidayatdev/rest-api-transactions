package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rest-api-transaction/config"
	"github.com/rest-api-transaction/router"
)

//App struct
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

//main function
func main() {
	error := godotenv.Load()

	if error != nil {
		log.Fatalf("Error loading .env file: %v", error)
	}

	app := App{}
	app.Initialize()
}

//Initialize func
func (app *App) Initialize() {
	app.Router = mux.NewRouter()

	connection := config.GetConnection()

	router.Route(app.Router, connection)

	apiServicePort := os.Getenv("API_PORT")

	if apiServicePort == "" {
		apiServicePort = "8080"
	}

	http.Handle("/", app.Router)

	log.Printf("API Service listening on port %v", apiServicePort)

	apiServer := &http.Server{
		Addr:    ":" + apiServicePort,
		Handler: config.EnableCorsOptions(app.Router),
	}

	error := apiServer.ListenAndServe()

	if error != nil && error != http.ErrServerClosed {
		log.Fatal(error)
	}

	defer connection.Close()
}
