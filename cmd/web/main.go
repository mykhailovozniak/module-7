package main

import (
	"context"
	"flag"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"module-7/pkg/models/mongodb"
	"net/http"
	"os"

	"module-7/pkg/models"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	materials interface{
		FindAll(ctx context.Context) ([]*models.Material, error)
	}
	cache map[string] string
}

func main()  {
	loadErr := godotenv.Load()

	if loadErr != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	addr := flag.String("addr", ":" + port, "HTTP network address")
	flag.Parse()

	mongoUri := os.Getenv("MONGO_URI")

	clientOptions := options.
		Client().
		ApplyURI(mongoUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	database := client.Database("test")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	infoLog.Println("Connected to MongoDB!")

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		materials: &mongodb.MaterialsModel{Collection: database.Collection("materials")},
		cache: make(map[string]string),
	}

	srv := &http.Server{
		Addr: *addr,
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()

	if err != nil {
		app.errorLog.Fatal("Error during start app")
	}
}
