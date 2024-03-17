package handler

import (
	"blog/model"
	"blog/service"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

func (blogHandler *BlogHandler) CreateBlog(writer http.ResponseWriter, req *http.Request) {
	// Ispisivanje podataka o zahtevu koji dolazi
	log.Println("Received request to create blog")

	// Čitanje JSON podataka iz tela zahteva
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error reading request body:", err)
		http.Error(writer, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// Ispisivanje JSON podataka pre dekodiranja
	log.Println("Received JSON data:", string(body))

	// Modifikacija polja difficultyLevel i status pre dekodiranja
	modifiedBody := modifyJSON(body)
	var blog model.Blog
	decoder := json.NewDecoder(bytes.NewReader(modifiedBody))
	err = decoder.Decode(&blog)
	if err != nil {
		println("Error while parsing json")
		http.Error(writer, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	err = blogHandler.BlogService.CreateBlog(&blog)
	if err != nil {
		log.Println("Error while creating new blog:", err)
		http.Error(writer, "Failed to create a new tour", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
func modifyJSON(data []byte) []byte {
	var modifiedData map[string]interface{}
	if err := json.Unmarshal(data, &modifiedData); err != nil {
		log.Println("Error decoding JSON:", err)
		return data
	}

	// Konverzija statusa iz stringa u broj
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

// Funkcija za konverziju statusa u broj
func convertStatusToNumber(status string) int {
	switch status {
	case "Draft":
		return 0
	case "Published":
		return 1
	case "Archived":
		return 2
	default:
		return -1
	}
}
