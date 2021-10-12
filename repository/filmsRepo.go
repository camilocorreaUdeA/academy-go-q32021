package repository

import (
	"encoding/csv"
	"log"
	"os"
)

type FilmsRepository interface {
	UpdateCSVFile(record []string) error
	ReadCSVFile() ([][]string, error)
}

type filmsRepository struct {
	file string
}

// NewFilmsRepo returns a new instance of the FilmsRepository struct
func NewFilmsRepo(filename string) filmsRepository {
	return filmsRepository{
		file: filename,
	}
}

// ReadCSVFile returns all the records found in the respository csv file
func (fr filmsRepository) ReadCSVFile() ([][]string, error) {
	return fr.readCSVFileRecords()
}

// UpdateCSVFile appends a new record to the repository csv file
func (fr filmsRepository) UpdateCSVFile(record []string) error {
	records, err := fr.readCSVFileRecords()
	if err != nil {
		log.Printf("failed to read csv file: %s", err)
		return err
	}
	records = append(records, record)
	err = fr.writeRecordsToCSVFile(records)
	if err != nil {
		log.Printf("failed to write csv file: %s", err)
		return err
	}
	return nil
}

func (fr filmsRepository) readCSVFileRecords() ([][]string, error) {
	file, err := os.Open(fr.file)
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

func (fr *filmsRepository) writeRecordsToCSVFile(records [][]string) error {
	file, err := os.OpenFile(fr.file, os.O_WRONLY, 777)
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
