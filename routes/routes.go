package routes

import (
	"net/http"

	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/controllers"
)

func RegisterRoutes() {
	http.HandleFunc("/data", controllers.GetAllData)
	http.HandleFunc("/data/detail", controllers.GetDataByCode)
	http.HandleFunc("/data/filter", controllers.FilterData)
	http.HandleFunc("/data/add", controllers.AddData)
	http.HandleFunc("/data/update", controllers.UpdateData)
	http.HandleFunc("/data/delete", controllers.DeleteData)
}
