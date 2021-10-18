package router

import (
	"TortoiseAssignment/controller"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Router(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	dbHandler := controller.NewDBHandler(db)

	router.HandleFunc("/profiles", dbHandler.CreateAProfile).Methods("POST")
	router.HandleFunc("/profiles", dbHandler.GetAllProfiles).Methods("GET")
	router.HandleFunc("/toggleprofile/{profileid}", dbHandler.ToggleStatusOfAProfile).Methods("PUT")
	router.HandleFunc("/pausedprofiles", dbHandler.GetAllPausedProfiles).Methods("GET")
	router.HandleFunc("/profiles/{profileid}", dbHandler.DeleteProfile).Methods("DELETE")

	return router
}
