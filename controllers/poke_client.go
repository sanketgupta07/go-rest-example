package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sanketgupta07/go-rest-example/model"
)

//PokeGetClient -- to get all data
func PokeGetClient(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		log.Println(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var rspObj model.Response

	json.Unmarshal(responseData, &rspObj)

	log.Println("Name: ", rspObj.Name)
	log.Println("Count of Pokemon: ", len(rspObj.Pokemon))

	// for i := 0; i < len(rspObj.Pokemon); i++ {
	// 	log.Println(rspObj.Pokemon[i].EntryNo, " : ", rspObj.Pokemon[i].Species.Name)
	// }
	json.NewEncoder(w).Encode(rspObj)

}
