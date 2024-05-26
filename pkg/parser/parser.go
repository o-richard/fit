package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
)

// allowed headers consists of a valid list of header names matching the first row of the csv file.
// if it is an empty slice, all headers are included in the result.
func parseCSV(filepath string, allowedHeaders []string) ([]map[string]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error while opening csv file, %w", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	// Obtain csv headers
	headers, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("error while parsing csv headers, %w", err)
	}
	// Obtain indices of concern
	indices := make([]int, 0, len(headers))
	for i, header := range allowedHeaders {
		index := slices.Index(headers, header)
		if index < 0 {
			return nil, fmt.Errorf("header %v does not exist in the file %v", header, filepath)
		}
		indices = append(indices, i)
	}
	if len(indices) == 0 {
		for i := 0; i < len(headers); i++ {
			indices = append(indices, i)
		}
	}

	var records []map[string]string
	for {
		row, err := csvReader.Read()
		if err != nil {
			if !errors.Is(err, io.EOF) {
				return nil, fmt.Errorf("error while parsing csv data, %w", err)
			}
			break
		}
		data := make(map[string]string, len(indices))
		for _, index := range indices {
			data[headers[index]] = row[index]
		}
		records = append(records, data)
	}
	return records, nil
}
