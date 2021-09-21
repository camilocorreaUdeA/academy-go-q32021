package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

const filePath = "resources/items.csv"

func ReadCSV(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[1:]
	response := []models.Item{}
	for _, line := range lines {
		lineContents := strings.Split(string(line), ",")
		itemID, _ := strconv.Atoi(lineContents[0])
		resp := models.Item{
			ID:   itemID,
			Name: lineContents[1],
			Type: lineContents[2],
		}
		response = append(response, resp)
	}
	responseItems := []models.Item{}
	requestedType := r.URL.Query().Get("type")
	fmt.Println(requestedType)
	for _, item := range response {
		if item.Type == requestedType {
			responseItems = append(responseItems, item)
		}
	}
	w.WriteHeader(http.StatusOK)
	respJson, _ := json.Marshal(responseItems)
	w.Write(respJson)
	return
}
