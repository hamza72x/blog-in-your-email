package data

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/hamza72x/blog-in-your-email/helper"
	"github.com/hamza72x/blog-in-your-email/model"
)

func GetBlogDataFromCSV() []model.Blog {

	lines, err := readCsv(helper.GetDataFile())

	if err != nil {
		panic(err)
	}

	var data []model.Blog

	// skipping FIRST line which is the header/titles of the columns
	for _, line := range lines[1:] {

		title := strings.TrimSpace(line[0])
		link := strings.TrimSpace(line[1])

		if len(title) == 0 || len(link) == 0 {
			log.Println("Skipping line:", line)
			continue
		}

		data = append(data, model.Blog{Title: title, Link: link})
	}

	return data
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func readCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
