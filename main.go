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
	connectionURL := "user=postgres password=super dbname=SOA host=database1 port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionURL), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	database.AutoMigrate(&model.Blog{})
	return database
}

func startServer(blogHandler *handler.BlogHandler) {
	router := mux.NewRouter().StrictSlash(true)
	//Blog
	router.HandleFunc("/blogs", blogHandler.CreateBlog).Methods("POST")
	router.HandleFunc("/blogs/all", blogHandler.GetAll).Methods("GET")
	router.HandleFunc("/blogs/{id}", blogHandler.GetBlogByID).Methods("GET")
	router.HandleFunc("/blogs/comments/{blogId}", blogHandler.SetBlogComments).Methods("PUT")

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	database := initDB()

	// Prikaz detalja blog posta

	blogRepo := &repo.BlogRepository{DatabaseConnection: database}
	blogService := &service.BlogService{BlogRepository: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	startServer(blogHandler)

}
