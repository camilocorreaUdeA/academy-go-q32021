package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/camilocorreaUdeA/academy-go-q32021/common"
	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

const filePath = "assets/items.csv"

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
	respJson, _ := json.Marshal(responseItems)
	w.Write(respJson)
	return
}
