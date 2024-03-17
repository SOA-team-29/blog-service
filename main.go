package main

import (
	"blog/handler"
	"blog/model"
	"blog/repo"
	"blog/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/russross/blackfriday/v2"
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
	database.AutoMigrate(&model.Blog{}, &model.BlogComment{}, &model.BlogRating{})
	return database
}

func startServer(blogHandler *handler.BlogHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/blogs", blogHandler.CreateBlog).Methods("POST")
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func ShowBlogDetails(blog *model.Blog) {
	// Pretvaranje Markdown-a u HTML
	html := blackfriday.Run([]byte(blog.DescriptionMarkdown))

	// html sadrži HTML verziju opisa bloga koju možete prikazati na frontend-u
	fmt.Println("HTML verzija opisa bloga:")
	fmt.Println(string(html))
}
func main() {
	database := initDB()

	blog := &model.Blog{
		// Postavite vrednosti polja Blog strukture
		AuthorID:            0,                                  // Postavljanje novog UUID za AuthorID
		TourID:              1,                                  // Postavljanje novog UUID za TourID
		Title:               "My Blog Title",                    // Postavljanje naslova bloga
		Description:         "My Blog Description",              // Postavljanje opisa bloga
		DescriptionMarkdown: "*This is a markdown description*", // Postavljanje Markdown opisa bloga
		PublishedDateTime:   nil,                                // Postavljanje vremena objave na nil (nije još objavljen)
		ImageID:             pq.StringArray{"image1", "image2"}, // Postavljanje niza slika
		Status:              0,                                  // Postavljanje statusa na Draft
		BlogComments:        nil,                                // Inicijalizacija komentara na nil
		BlogRatings:         nil,                                // Inicijalizacija ocena na nil

	}

	// Prikaz detalja blog posta
	ShowBlogDetails(blog)
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	blogRepo := &repo.BlogRepository{DatabaseConnection: database}
	blogService := &service.BlogService{BlogRepo: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	startServer(blogHandler)

}
