package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB
var err error

var dsn = "host=localhost user=postgres password=3818200 dbname=gorm port=5432 sslmode=disable"

type User struct {
	gorm.Model
	FirstName string `json:"firstnme"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")
	var user []User
	params := mux.Vars(r)
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")
	var user User
	params := mux.Vars(r)
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)

}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Conent-Type", "application/json")
	var user []User
	params := mux.Vars(r)
	DB.Delete(&user, params["id"])

	json.NewEncoder(w).Encode(`user is deleted ` + params["id"])
}
