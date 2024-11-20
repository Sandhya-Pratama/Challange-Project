package data

import (
	"encoding/csv"
	"os"
	"strings"

	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/models"
)

func LoadData(filename string) ([]models.Data, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []models.Data
	for _, record := range records[1:] {
		techs := strings.Split(record[4], ",")
		data = append(data, models.Data{
			Code:        record[0],
			Name:        record[1],
			Description: record[2],
			Model:       record[3],
			Tech:        techs,
			Status:      record[5],
		})
	}
	return data, nil
}
