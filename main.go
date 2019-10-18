package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/febrielven/go_crud_posgres/module/model"
	"github.com/febrielven/go_crud_posgres/module/repository"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Totorial Golang")
	router := mux.NewRouter()
	router.HandleFunc("/profile", getProfiles).Methods("GET")
	router.HandleFunc("/profile/{id}", detailProfile).Methods("GET")
	router.HandleFunc("/profile", saveProfile).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}

func getProfiles(w http.ResponseWriter, req *http.Request) {
	// db, err := config.GetPostgersDB()
	profileRepository := repository.NewProfileRepository()
	profiles, err := profileRepository.FindAll()

	if err != nil {
		log.Print(err)
	}

	json.NewEncoder(w).Encode(profiles)

}

func detailProfile(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}

	// db, err := config.GetPostgersDB()
	profileRepository := repository.NewProfileRepository()
	profile, err := profileRepository.FindById(id)

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(profile)

}

func saveProfile(w http.ResponseWriter, req *http.Request) {

	profile := model.NewProfile()
	_ = json.NewDecoder(req.Body).Decode(&profile)

	// db, err := config.GetPostgersDB()
	profileRepository := repository.NewProfileRepository()

	// wury := model.NewProfile()
	// wury.ID = 6
	// wury.FirstName = "febrianto2"
	// wury.LastName = "al"
	// wury.Email = "@febri@gmail.com"
	// wury.Password = "1234567"

	err := profileRepository.Save(profile)

	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(profile)
}
func updateProfile(p *model.Profile, repo repository.ProfileRepository) error {
	err := repo.Update(p.ID, p)

	if err != nil {
		return err
	}

	return nil

}

func deleteProfile(id int, repo repository.ProfileRepository) error {
	err := repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
