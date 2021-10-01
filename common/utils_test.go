package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateCSVFile(t *testing.T) {
	asserter := assert.New(t)
	record := []string{"100", "dummy", "dummy"}
	err := UpdateCSVFile("./testdata/test.csv", record)
	asserter.Nil(err)
}

func Test_readCSVFileRecords(t *testing.T) {
	asserter := assert.New(t)
	records, err := readCSVFileRecords("./testdata/test.csv")
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
	err := writeRecordsToCSVFile("./testdata/test.csv", records)
	asserter.Nil(err)
}
