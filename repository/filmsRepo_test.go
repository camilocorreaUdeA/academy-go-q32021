package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testFile = "./testdata/test.csv"

func TestUpdateCSVFile(t *testing.T) {
	asserter := assert.New(t)
	record := []string{"100", "dummy", "dummy"}
	repo := NewFilmsRepo(testFile)
	err := repo.UpdateCSVFile(record)
	asserter.Nil(err)
}

func TestReadCSVFile(t *testing.T) {
	asserter := assert.New(t)
	repo := NewFilmsRepo(testFile)
	records, err := repo.ReadCSVFile()
	asserter.Nil(err)
	asserter.NotEmpty(records)
}

func Test_readCSVFileRecords(t *testing.T) {
	asserter := assert.New(t)
	repo := NewFilmsRepo(testFile)
	records, err := repo.readCSVFileRecords()
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
	repo := NewFilmsRepo(testFile)
	err := repo.writeRecordsToCSVFile(records)
	asserter.Nil(err)
}
