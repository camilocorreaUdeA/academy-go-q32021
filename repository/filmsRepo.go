package repository

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/camilocorreaUdeA/academy-go-q32021/models"
)

type Repository interface {
	UpdateCSVFile(filename string, record []string) error
}

type FilmsRepository struct {
}

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

func (fr *FilmsRepository) UpdateCSVFile(filename string, record []string) error {
	records, err := fr.readCSVFileRecords(filename)
	if err != nil {
		log.Printf("failed to read csv file: %s", err)
		return err
	}
	records = append(records, record)
	err = fr.writeRecordsToCSVFile(filename, records)
	if err != nil {
		log.Printf("failed to write csv file: %s", err)
		return err
	}
	return nil
}

func (fr *FilmsRepository) readCSVFileRecords(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("could not open csv file (read): %s", err)
		return [][]string{}, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Printf("failed read all operation: %s", err)
		return [][]string{}, err
	}
	return records, nil
}

func (fr *FilmsRepository) writeRecordsToCSVFile(filename string, records [][]string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 777)
	if err != nil {
		log.Printf("could not open csv file (write): %s", err)
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		log.Printf("failed write all operation: %s", err)
		return err
	}
	return nil
}
