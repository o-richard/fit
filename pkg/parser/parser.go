package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/o-richard/fit/pkg/db"
)

type Parser interface {
	GetEntries() ([]db.HealthEntry, error)
}

// returns csv records. allowed headers consists of a valid list of header names matching the first row of the csv file. if it is an empty slice, all headers are included in the result.
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

// returns filenames in the specified directory with the provided filenames as substrings (should be unique). if it is empty, no files are returned.
// for the returned filenames, the first substring match is considered so the provided filenames should be as distinguishable for correct results.
func filePathWalk(directorypath string, filenames []string) (map[string]string, error) {
	entries, err := os.ReadDir(directorypath)
	if err != nil {
		return nil, fmt.Errorf("error while opening directory, %w", err)
	}

	files := make(map[string]string, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			for _, filename := range filenames {
				if strings.Contains(entry.Name(), filename) {
					files[filename] = entry.Name()
					break
				}
			}
		}
	}
	return files, nil
}
