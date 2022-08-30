package utils

import (
	"encoding/csv"
	"os"
)

func ExportToCsv(data [][]string, csvName string) error {
	file, err := os.Create(csvName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		if err := writer.Write(value); err != nil {
			return err // let's return errors if necessary, rather than having a one-size-fits-all error handler
		}
	}
	return nil
}