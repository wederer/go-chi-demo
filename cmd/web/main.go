package main

import (
	"fmt"
	driver "github.com/arangodb/go-driver"
	driverhttp "github.com/arangodb/go-driver/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	s := CreateNewServer()

	s.MountMiddlewares()
	s.MountHandlers()
	s.SetupDatabase()

	port := "3000"
	if value, ok := os.LookupEnv("PORT"); ok {
		port = value
	}
	log.Printf("Listening on port %v\n", port)
	port = fmt.Sprintf(":%v", port)
	err := http.ListenAndServe(port, s.Router)
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

type Server struct {
	Router *chi.Mux
	Books  driver.Collection
}

func (s *Server) MountMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.CleanPath)
}
func (s *Server) MountHandlers() {
	s.Router.Get("/", s.HelloWorld)
	s.Router.Get("/book/{id}", s.GetBook)
	s.Router.Post("/book", s.CreateBook)
}

func (s *Server) SetupDatabase() {
	dbUrl := os.Getenv("DB_URL")
	dbUser := os.Getenv("DB_USER")
	dbPw := os.Getenv("DB_PW")
	if dbUrl == "" || dbUser == "" || dbPw == "" {
		log.Fatalln("Database config not provided via environment.")
	}

	conn, err := driverhttp.NewConnection(driverhttp.ConnectionConfig{
		Endpoints: []string{dbUrl},
	})
	if err != nil {
		log.Fatalf("Failed to create HTTP connection: %v", err)
	}
	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(dbUser, dbPw),
	})

	const DB_NAME = "examples_books"
	dbExists, err := c.DatabaseExists(nil, DB_NAME)
	if err != nil {
		log.Fatalf("Failed to check for database: %v", err)
	}
	var db driver.Database
	if dbExists == false {
		db, err = c.CreateDatabase(nil, DB_NAME, nil)
	} else {
		db, err = c.Database(nil, DB_NAME)
	}
	if err != nil {
		log.Fatalf("Failed to create/get database: %v", err)
	}

	colExists, err := db.CollectionExists(nil, "books")
	var coll driver.Collection
	if colExists == true {
		coll, err = db.Collection(nil, "books")
	} else {
		coll, err = db.CreateCollection(nil, "books", nil)
	}
	if err != nil {
		log.Fatalf("Failed to create/get collection: %v", err)
	}

	s.Books = coll
}
