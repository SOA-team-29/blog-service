package main

import (
	"blog/handler"
	"blog/model"
	"blog/repo"
	"blog/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connection_url := "user=postgres password=super dbname=SOA port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	database.AutoMigrate(&model.Blog{})
	return database
}

func startServer(blogHandler *handler.BlogHandler, blogComHandler *handler.BlogCommentHandler) {
	router := mux.NewRouter().StrictSlash(true)
	//Blog
	router.HandleFunc("/blogs", blogHandler.CreateBlog).Methods("POST")
	router.HandleFunc("/blogs/all", blogHandler.GetAll).Methods("GET")
	router.HandleFunc("/blogs/{id}", blogHandler.Get).Methods("GET")
	//BlogComments
	router.HandleFunc("/comments/{blogId}", blogComHandler.GetByBlogId).Methods("GET")
	router.HandleFunc("/comments/add{blogId}", blogComHandler.Create).Methods("Post")
	router.HandleFunc("/comments/delete{blogid}", blogComHandler.Delete).Methods("PUT")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	database := initDB()

	// Prikaz detalja blog posta

	blogCommentRepo := &repo.BlogCommentRepository{DatabaseConnection: database}
	blogCommentService := &service.BlogCommentService{BlogCommentRepository: blogCommentRepo}
	blogCommentHandler := &handler.BlogCommentHandler{BlogCommentService: blogCommentService}
	blogRepo := &repo.BlogRepository{DatabaseConnection: database}
	blogService := &service.BlogService{BlogRepository: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	startServer(blogHandler, blogCommentHandler)

}
