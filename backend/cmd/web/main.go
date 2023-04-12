package main

import (
	"encoding/json"
	driver "github.com/arangodb/go-driver"
	driverhttp "github.com/arangodb/go-driver/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/wederer/go-chi-demo/internal/models"
	"log"
	"net/http"
	"os"
)

func main() {
	s := CreateNewServer()

	s.MountMiddlewares()
	s.MountHandlers()
	s.SetupDatabase()

	s.Start()
}

type Server struct {
	Router *chi.Mux
	DB     driver.Database
	Books  models.ModelInterface[models.Book]
}

func CreateNewServer() *Server {
	s := &Server{}
	s.Router = chi.NewRouter()
	return s
}

func (s *Server) MountMiddlewares() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.CleanPath)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
}
func (s *Server) MountHandlers() {
	// Public Routes
	s.Router.Group(func(r chi.Router) {
		s.Router.Get("/", s.HelloWorld)
		s.Router.Get("/books", s.GetBooks)
		s.Router.Get("/books/{id}", s.GetBook)
		s.Router.Post("/books", s.CreateBook)
		s.Router.Delete("/books/{id}", s.DeleteBook)
	})

	// Protected Routes
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	s.Router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Get("/protected", s.GetProtectedData)

		// Admin Routes that are only accessible if admin: true is set
		r.Route("/admin", func(r chi.Router) {
			r.Use(ClaimBoolAuthenticator("admin", true))
			r.Get("/info", s.GetAdminInfo)
		})
	})

}

// ClaimBoolAuthenticator returns an Authenticator that checks if a claim is present and equal to the given value
func ClaimBoolAuthenticator(claim string, value any) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, _ := jwtauth.FromContext(r.Context())

			if claimValue, ok := token.Get(claim); ok == false || claimValue != value {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
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

	s.DB = db
	s.Books = models.New(db, "books")
}

func (s *Server) Start() {
	addr := ":3000"
	if value, ok := os.LookupEnv("SRV_ADDR"); ok {
		addr = value
	}
	log.Printf("Listening on address %v\n", addr)
	err := http.ListenAndServe(addr, s.Router)
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}

func Marshall[T any](in T, w http.ResponseWriter) {
	mJson, err := json.Marshal(in)

	if err != nil {
		log.Printf("Failed to marshall document: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(mJson)
}
