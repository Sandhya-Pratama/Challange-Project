package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/models"
	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/services"
	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/views"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	data := services.GetAllData()
	views.JSONResponse(w, data)
}

func GetDataByCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	data, err := services.GetDataByCode(code)
	if err != nil {
		views.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}
	views.JSONResponse(w, data)
}

func FilterData(w http.ResponseWriter, r *http.Request) {
	model := r.URL.Query().Get("model")
	tech := r.URL.Query().Get("tech")
	data := services.FilterData(model, tech)
	views.JSONResponse(w, data)
}

func AddData(w http.ResponseWriter, r *http.Request) {
	var newData models.Data
	err := json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		views.ErrorResponse(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	err = services.AddData(newData)
	if err != nil {
		views.ErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	views.JSONResponse(w, map[string]string{"message": "Data added successfully"})
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	var updatedData models.Data

	err := json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		views.ErrorResponse(w, "Invalid input format", http.StatusBadRequest)
		return
	}

	err = services.UpdateData(code, updatedData)
	if err != nil {
		views.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	views.JSONResponse(w, map[string]string{"message": "Data updated successfully"})
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	err := services.DeleteData(code)
	if err != nil {
		views.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	views.JSONResponse(w, map[string]string{"message": "Data deleted successfully"})
}
