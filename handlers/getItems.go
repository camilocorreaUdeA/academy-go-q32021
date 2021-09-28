package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

// GetItems handles the request and returns back requested items
// in json encoded response.
func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := common.ReadCSVFile(filePath)
	if err != nil {
		log.Printf("Failed to read the file: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	responseItems := []models.Item{}
	requestedType := r.URL.Query().Get("type")
	for _, item := range items {
		if item.Type == requestedType {
			responseItems = append(responseItems, item)
		}
	}
	w.WriteHeader(http.StatusOK)
	respJson, err := json.Marshal(responseItems)
	if err != nil {
		log.Printf("Failed to marshal handler response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(respJson)
	return
}
