package handler

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

// server is the base structure of the API
type server struct {
	router   *httprouter.Router
	database *sql.DB
}

// response contains all response infos at a glance
type response struct {
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
	Message    string      `json:"message"`
	Meta       struct {
		Query       interface{} `json:"query,omitempty"`
		ResultCount int         `json:"result_count,omitempty"`
	} `json:"meta"`
	Data []interface{} `json:"data"`
}

// StartWebServer is the function responsible for launching the API
func StartWebServer() {
	// get DB credential from environment variables
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// establish database connection
	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	// mount router and database to the server and launch
	s := server {
		router:   httprouter.New(),
		database: db,
	}
	s.router.PanicHandler = handlePanic
	s.routes()
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), s.router))
}

// routes function launches all application's routes
func (s *server) routes() {
	//home
	s.router.GET("/", s.handleGetHome())
	s.router.GET("/questions/:id", s.handleQuestionByID())
	s.router.GET("/api/questions/:limit", s.HandleAllQuestions())
	//s.router.POST("/new", s.handlePostQuote())
	//replique
	// random
}

// Gracefully handle panic without crashing the server
func handlePanic(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, err)
	w.WriteHeader(http.StatusInternalServerError)
}
