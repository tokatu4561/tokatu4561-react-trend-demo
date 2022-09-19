package main

import (
	"database/sql"
	"fmt"
	"log"
	"myapp/models"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type config struct {
	port string
	host string
	// db   struct {
	// 	dsn string
	// }
	stripe struct {
		secret string
		key    string
	}
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	DB       *sql.DB
	Models   models.Models
}

func main() {
	var cfg config

	_ = godotenv.Load(".env")

	cfg.port = os.Getenv("API_PORT")
	cfg.stripe.key = os.Getenv("PUBLIC_STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("SECRET_STRIPE_KEY")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := sql.Open("postgres", "user=postgres password=password host=localhost port=5432 dbname=practice sslmode=disable")
	if err != nil {
		errorLog.Fatalln(err)
	}
	defer db.Close()

	app := &application{
		config:   cfg,
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	err = app.serve()
	if err != nil {
		log.Fatal(err)
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting Back end server %s port %s\n", app.config.host, app.config.port)

	return srv.ListenAndServe()
}
