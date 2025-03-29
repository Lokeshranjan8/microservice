package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Lokeshranjan8/user-svc/models"
	"github.com/Lokeshranjan8/user-svc/utils"
)

func Create_user(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err !=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result := utils.DB.Create(&user)
	
	if result.Error != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, result.Error.Error()), http.StatusBadRequest)
		return
	}

	response := models.Response{
		Message:     "User Created",
		Explanation: "A new user with given information has been created",
		Data:        user,
	}
	json.NewEncoder(w).Encode(response)
}




func Get_user(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	userID := r.URL.Query().Get("user_id")
	if userID ==" "{
		http.Error(w,`{"error":"user_id param is required}`,http.StatusBadRequest)
		return
	}
	var user models.User
	if err := utils.DB.First(&user, userID).Error; err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusNotFound)
		return
	}
	response := models.Response{
		Message:     "User Found",
		Explanation: "User with the given ID is found",
		Data:        user,
	}
	json.NewEncoder(w).Encode(response)


}
func Getall_user(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var users []models.User
	if err := utils.DB.Find(&users).Error; err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	response := models.Response{
		Message: "users data",
		Explanation: "data of all users" ,
		Data:users,
	}

	json.NewEncoder(w).Encode(response)


}