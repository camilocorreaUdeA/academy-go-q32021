package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCSVFile(t *testing.T) {
	asserter := assert.New(t)
	record := []string{"100", "dummy", "dummy"}
	repo := NewFilmsRepo()
	err := repo.UpdateCSVFile("./testdata/test.csv", record)
	asserter.Nil(err)
}

func TestReadCSVFile(t *testing.T) {
	asserter := assert.New(t)
	repo := NewFilmsRepo()
	records, err := repo.ReadCSVFile("./testdata/test.csv")
	asserter.Nil(err)
	asserter.NotEmpty(records)
}

func Test_readCSVFileRecords(t *testing.T) {
	asserter := assert.New(t)
	repo := &FilmsRepository{}
	records, err := repo.readCSVFileRecords("./testdata/test.csv")
	asserter.Nil(err)
	asserter.NotEmpty(records)
}

func Test_writeRecordsToCSVFile(t *testing.T) {
	asserter := assert.New(t)
	records := [][]string{
		{"Id", "Name", "Type"},
		{"1", "House", "Home"},
		{"2", "Shoes", "Clothes"},
		{"3", "Notebook", "Scholar"},
		{"4", "IPhone", "Electronics"},
	}
	repo := &FilmsRepository{}
	err := repo.writeRecordsToCSVFile("./testdata/test.csv", records)
	asserter.Nil(err)
}
