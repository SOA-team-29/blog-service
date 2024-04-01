package handler

import (
	"blog/model"
	"blog/service"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (h *BlogHandler) CreateBlog(w http.ResponseWriter, r *http.Request) {
	// Ispisivanje podataka o zahtevu koji dolazi
	log.Println("Received request to create blog")

	// Čitanje JSON podataka iz tela zahteva
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Ispisivanje JSON podataka pre dekodiranja
	log.Println("Received JSON data:", string(body))

	// Modifikacija polja difficultyLevel i status pre dekodiranja
	modifiedBody := modifyJSON(body)

	// Dekodiranje JSON podataka iz tela zahteva u tour objekat
	var blog model.Blog
	decoder := json.NewDecoder(bytes.NewReader(modifiedBody))

	err = decoder.Decode(&blog)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Prosleđivanje tour objekta servisu za kreiranje ture
	err = h.BlogService.CreateBlog(&blog)
	if err != nil {
		log.Println("Error while creating a new tour:", err)
		http.Error(w, "Failed to create a new tour", http.StatusInternalServerError)
		return
	}

	// Slanje odgovora da je tura uspešno kreirana
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}

// Funkcija za modifikaciju JSON podataka
func modifyJSON(data []byte) []byte {
	var modifiedData map[string]interface{}
	if err := json.Unmarshal(data, &modifiedData); err != nil {
		log.Println("Error decoding JSON:", err)
		return data
	}

	// Konverzija difficultyLevel iz stringa u broj
	if status, ok := modifiedData["status"].(string); ok {
		modifiedData["status"] = convertStatusToNumber(status)
	}

	// Konverzija nazad u JSON
	modifiedBody, err := json.Marshal(modifiedData)
	if err != nil {
		log.Println("Error encoding modified JSON:", err)
		return data
	}

	return modifiedBody
}

// Funkcija za konverziju difficultyLevel u broj
func convertStatusToNumber(status string) int {
	switch status {
	case "DRAFT":
		return 0
	case "PUBLISHED":
		return 1
	case "CLOSED":
		return 2
	case "ACTIVE":
		return 3
	case "FAMOUS":
		return 4
	default:
		return -1
	}
}

func (h *BlogHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to get all blogs")
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}

	blogs, err := h.BlogService.GetAll(page, pageSize)
	if err != nil {
		log.Println("Error getting tours :", err)
		http.Error(w, "Failed to get tours", http.StatusInternalServerError)
		return
	}

	log.Println("Blogs:", blogs)
	modifiedBlogs := make([]map[string]interface{}, len(*blogs))
	for i, blog := range *blogs {
		modifiedBlog := map[string]interface{}{
			"id":           blog.ID,
			"authorId":     blog.AuthorID,
			"tourId":       blog.TourID,
			"title":        blog.Title,
			"description":  blog.Description,
			"creationDate": blog.CreationDate,
			"imageURLs":    blog.ImageURLs,
			"comments":     blog.Comments,
			"ratings":      blog.Ratings,
			"status":       convertStatusToString(int(blog.Status)),
		}

		// Konvertujte TransportType u string u modifikovanom tura objektu
		comments := make([]map[string]interface{}, len(blog.Comments))
		for j, comment := range blog.Comments {
			comments[j] = map[string]interface{}{
				"text":            comment.Text,
				"userId":          comment.UserID,
				"creationTime":    comment.CreationTime,
				"lastUpdatedTime": comment.LastUpdatedTime,
			}
		}
		modifiedBlog["comments"] = comments

		modifiedBlogs[i] = modifiedBlog
	}
	log.Println(modifiedBlogs)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modifiedBlogs)
}

func convertStatusToString(status int) string {
	switch status {
	case 0:
		return "DRAFT"
	case 1:
		return "PUBLISHED"
	case 2:
		return "CLOSED"
	case 3:
		return "ACTIVE"
	case 4:
		return "FAMOUS"
	default:
		return ""
	}
}
func (h *BlogHandler) SetBlogComments(w http.ResponseWriter, r *http.Request) {

	log.Println("Received request to set blog comments")

	params := mux.Vars(r)
	blogIDStr, ok := params["blogId"]
	if !ok {
		log.Println("Blog ID not provided")
		http.Error(w, "Blog ID not provided", http.StatusBadRequest)
		return
	}
	blogID, err := strconv.Atoi(blogIDStr)
	if err != nil {
		log.Println("Invalid blog ID:", err)
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	// Čitanje JSON podataka iz tela zahteva
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Ispisivanje JSON podataka pre dekodiranja
	log.Println("Received JSON data:", string(body))

	// Modifikacija polja transport pre dekodiranja
	modifiedBody := modifyJSONForComments(body)

	var blogComment []model.BlogComment
	decoder := json.NewDecoder(bytes.NewReader(modifiedBody))

	err = decoder.Decode(&blogComment)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.BlogService.SetBlogComments(blogID, blogComment)
	if err != nil {
		log.Println("Error setting blog comments:", err)
		http.Error(w, "Failed to set blog comments", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog comments successfully set"))
}

func modifyJSONForComments(data []byte) []byte {
	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		log.Println("Error decoding JSON:", err)
		return data
	}

	// Pretvaranje modifikovanog objekta u niz objekata
	blogComments := make([]model.BlogComment, 1)
	blogComments[0] = model.BlogComment{
		Text:            jsonData["text"].(string),
		UserID:          jsonData["userId"].(int),
		CreationTime:    jsonData["creationTime"].(time.Time),
		LastUpdatedTime: jsonData["lastUpdatedTime"].(time.Time),
	}

	// Konverzija nazad u JSON
	modifiedBody, err := json.Marshal(blogComments)
	if err != nil {
		log.Println("Error encoding modified JSON:", err)
		return data
	}

	return modifiedBody
}
func (h *BlogHandler) GetBlogByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	IDStr, ok := params["id"]
	if !ok {
		log.Println("ID not provided")
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}
	ID, err := strconv.Atoi(IDStr)
	if err != nil {
		log.Println("Invalid tour ID:", err)
		http.Error(w, "Invalid tour ID", http.StatusBadRequest)
		return
	}

	blog, err := h.BlogService.GetBlogByID(ID)
	if err != nil {
		log.Println("Error getting tour by ID:", err)
		http.Error(w, "Failed to get tour by ID", http.StatusInternalServerError)
		return
	}

	//Moja provera da li je nasao dobro iz baze
	log.Println("Blogs:", blog)
	modifiedBlog := map[string]interface{}{
		"id":           blog.ID,
		"authorId":     blog.AuthorID,
		"tourId":       blog.TourID,
		"title":        blog.Title,
		"description":  blog.Description,
		"creationDate": blog.CreationDate,
		"imageURLs":    blog.ImageURLs,
		"comments":     blog.Comments,
		"ratings":      blog.Ratings,
		"status":       convertStatusToString(int(blog.Status)),
	}

	// Konvertujte TransportType u string u modifikovanom tura objektu
	comments := make([]map[string]interface{}, len(blog.Comments))
	for j, comment := range blog.Comments {
		comments[j] = map[string]interface{}{
			"text":            comment.Text,
			"userId":          comment.UserID,
			"creationTime":    comment.CreationTime,
			"lastUpdatedTime": comment.LastUpdatedTime,
		}
	}
	modifiedBlog["comments"] = comments

	// Slanje odgovora sa tura podacima kao JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modifiedBlog)
}
