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

func (handler *BlogHandler) Get(writer http.ResponseWriter, req *http.Request) {
	blogId := mux.Vars(req)["id"]
	log.Printf("Blog with id: %s", blogId)

	blog, err := handler.BlogService.FindBlog(blogId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	/*comments, _ := handler.BlogCommentService.GetByBlogId(blogId)

	var commentPointers []*model.BlogComment
	for _, comment := range comments {
		commentPointers = append(commentPointers, &comment)
	}

	blog.BlogComments = commentPointers*/

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
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
		/*
			// Konvertujte TransportType u string u modifikovanom tura objektu
			characteristics := make([]map[string]interface{}, len(tour.TourCharacteristics))
			for j, characteristic := range tour.TourCharacteristics {
				characteristics[j] = map[string]interface{}{
					"distance":      characteristic.Distance,
					"duration":      characteristic.Duration,
					"transportType": convertTransportTypeToString(characteristic.TransportType),
				}
			}
			modifiedTour["tourCharacteristics"] = characteristics*/

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
