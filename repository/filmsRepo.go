package repository

import (
	"encoding/csv"
	"log"
	"os"
)

type Repository interface {
	UpdateCSVFile(filename string, record []string) error
	ReadCSVFile(filename string) ([][]string, error)
}

type FilmsRepository struct {
}

func NewFilmsRepo() FilmsRepository {
	return FilmsRepository{}
}

// ReadCSVFile returns a slice of model.Item objects that were
// read from the .CSV file in the specified file path.
func (fr FilmsRepository) ReadCSVFile(filename string) ([][]string, error) {
	return fr.readCSVFileRecords(filename)
}

func (fr FilmsRepository) UpdateCSVFile(filename string, record []string) error {
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

func (fr FilmsRepository) readCSVFileRecords(filename string) ([][]string, error) {
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
