package handlers

import (
	"encoding/json"
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
	w.WriteHeader(http.StatusOK)
	respJson, _ := json.Marshal(response)
	w.Write(respJson)
	return
}
