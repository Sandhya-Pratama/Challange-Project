package services

import (
	"errors"
	"fmt"

	"strings"

	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/models"
)

var data []models.Data

func InitData(d []models.Data) {
	data = d
}

func GetAllData() []models.Data {
	return data
}

func GetDataByCode(code string) (models.Data, error) {
	for _, d := range data {
		if d.Code == code {
			return d, nil
		}
	}
	return models.Data{}, errors.New("data not found")
}

// data bisa difilter lebih dari 1 dari tech
// validasi bisa ditambahkan untuk imput kode hanya boleh alfabet, angka, # dengan panjang karakter maksimal 15
// buat model referensi, kalau input diluar referensi maka akan ditolak
func FilterData(model, tech string) []models.Data {
	var filtered []models.Data
	techFilters := strings.Split(tech, ",")

	for _, d := range data {
		matchesModel := model == "" || d.Model == model
		matchesTech := true

		for _, t := range techFilters {
			if t != "" && !contains(d.Tech, t) {
				matchesTech = false
				break
			}
		}

		if matchesModel && matchesTech {
			filtered = append(filtered, d)
		}

		if matchesModel && !matchesTech {
			fmt.Println("modtidak tersedia, silahkan coba untuk ")
		}
	}
	return filtered
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func AddData(newData models.Data) error {
	for _, d := range data {
		if d.Code == newData.Code {
			return errors.New("code already exists")
		}
	}
	data = append(data, newData)
	return nil
}

func UpdateData(code string, updatedData models.Data) error {
	for i, d := range data {
		if d.Code == code {
			// Perbarui data
			data[i] = updatedData
			return nil
		}
	}
	return errors.New("data not found")
}

func DeleteData(code string) error {
	for i, d := range data {
		if d.Code == code {
			// Hapus elemen dari slice
			data = append(data[:i], data[i+1:]...)
			return nil
		}
	}
	return errors.New("data not found")
}
