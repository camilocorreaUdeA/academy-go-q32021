package common

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

// ReadCSVFile returns a slice of model.Item objects that were
// read from the .CSV file in the specified file path.
func ReadCSVFile(filePath string) ([]models.Item, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []models.Item{}, fmt.Errorf("unable to read file %s: %s", filePath, err)
	}
	lines := strings.Split(string(data), newLineChar)
	if len(lines) == 0 {
		return []models.Item{}, fmt.Errorf("file %s does not have any records", filePath)
	}
	lines = lines[1:]
	response := []models.Item{}
	for _, line := range lines {
		lineContents := strings.Split(string(line), delimiter)
		itemID, _ := strconv.Atoi(lineContents[0])
		resp := models.Item{
			ID:   itemID,
			Name: lineContents[1],
			Type: lineContents[2],
		}
		response = append(response, resp)
	}

	return response, nil
}
