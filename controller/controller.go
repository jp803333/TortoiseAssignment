package controller

import (
	"TortoiseAssignment/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBHandler struct {
	db *gorm.DB
}

func NewDBHandler(db *gorm.DB) *DBHandler {
	return &DBHandler{
		db: db,
	}
}

func printurl(r *http.Request) {
	fmt.Println(r.Method + " " + r.URL.Path)
}

func (dbh DBHandler) CreateAProfile(w http.ResponseWriter, r *http.Request) {
	printurl(r)

	w.Header().Set("Content-Type", "application/json")

	b, _ := ioutil.ReadAll(r.Body)

	var temp map[string]interface{}
	_ = json.Unmarshal([]byte(string(b)), &temp)

	dt, err := time.Parse("2006-01-02T15:04:05.000Z", temp["dateofbirth"].(string))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Date should like this 2006-01-02T15:04:05.000Z", http.StatusBadRequest)
		// json.NewEncoder(w).Encode("Date should like this 2006-01-02T15:04:05.000Z")
		return
	}
	status := temp["status"].(string)
	if status != "ACTIVE" && status != "PAUSED" {
		http.Error(w, "Status can either be \"ACTIVE\" or \"PAUSED\"", http.StatusBadRequest)
		// json.NewEncoder(w).Encode("Status can either be \"ACTIVE\" or \"PAUSED\"")
		return
	}
	var profile model.Profile = model.Profile{
		Name:        temp["name"].(string),
		Dateofbirth: dt,
		Status:      temp["status"].(string),
	}

	result := dbh.db.Table("profile").Create(&profile)

	if result.Error != nil {
		fmt.Println(result.Error.Error())
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(result.Error.Error())
		return
	}

	json.NewEncoder(w).Encode(profile)
}
func (dbh DBHandler) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	printurl(r)

	var profiles []model.Profile

	result := dbh.db.Table("profile").Find(&profiles)

	fmt.Println(profiles)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(result.Error.Error())
		return
	}

	json.NewEncoder(w).Encode(profiles)
}
func (dbh DBHandler) ToggleStatusOfAProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	printurl(r)

	params := mux.Vars(r)

	var profile model.Profile
	result := dbh.db.Raw("SELECT * FROM profile WHERE profileid = ?", params["profileid"]).Scan(&profile)
	if result.RowsAffected < 1 {
		http.Error(w, "No profile with given id", http.StatusBadRequest)
		// json.NewEncoder(w).Encode("No profile with given id")
		return
	}
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(result.Error.Error())
		return
	}
	if profile.Status == "ACTIVE" {
		profile.Status = "PAUSED"
	} else {
		profile.Status = "ACTIVE"
	}

	newresult := dbh.db.Table("profile").Save(&profile)

	if newresult.Error != nil {
		fmt.Println(newresult.Error.Error())
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(result.Error.Error())
		return
	}

	json.NewEncoder(w).Encode(profile)
}
func (dbh DBHandler) GetAllPausedProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	printurl(r)

	var profiles []model.Profile

	result := dbh.db.Table("profile").Where("status = ?", "PAUSED").Find(&profiles)

	fmt.Println(profiles)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(result.Error.Error())
		return
	}

	json.NewEncoder(w).Encode(profiles)
}
func (dbh DBHandler) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	printurl(r)

	params := mux.Vars(r)

	result := dbh.db.Table("profile").Delete(&model.Profile{}, params["profileid"])
	if result.RowsAffected < 1 {
		http.Error(w, "No profile with given id", http.StatusBadRequest)
		// json.NewEncoder(w).Encode("No profile with given id")
		return
	}
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		http.Error(w, "result.Error.Error()", http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(result.Error.Error())
		return
	}

	json.NewEncoder(w).Encode("succesfully deleted")
}
